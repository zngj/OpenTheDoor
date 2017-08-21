package service

import (
	"common/codec"
	"common/errcode"
	"common/model"
	"common/sgconst"
	"common/sqlx"
	"common/util"
	"fmt"
	"github.com/carsonsx/log4g"
	"etransin/dao"
	"time"
)

const (
	GATE_DIRECTION_IN         = int8(0)
	GATE_DIRECTION_OUT        = int8(1)
	EVIDENCE_STATUS_DELIVERED = int8(1)
	EVIDENCE_STATUS_USED      = int8(2)
	EVIDENCE_STATUS_EXPIRED   = int8(3)
)

func CreateEvidenceWithEncrypt(userId string, typ int8) (evidence *model.RouterEvidence, evidenceKey string, err error) {
	evidence, err = CreateEvidence(userId, typ)
	if err != nil {
		return
	}
	_evidenceKey := fmt.Sprintf("%s%d", evidence.EvidenceId, evidence.ExpiresAt.Unix())
	evidenceKey, err = codec.PrivateEncrypt(_evidenceKey)
	return
}

func CreateEvidence(userId string, typ int8) (evidence *model.RouterEvidence, err error) {
	evidence = new(model.RouterEvidence)
	evidence.EvidenceId = util.NewUuid()
	evidence.UserId = userId
	evidence.Direction = typ
	evidence.CreateTime = time.Now()
	evidence.ExpiresAt = evidence.CreateTime.Add(5 * time.Minute)
	evidence.Status = 1
	err = dao.NewEvidenceDao().Insert(evidence)
	if err == nil {
		log4g.Info("create new evidence %s for user %s", evidence.EvidenceId, evidence.UserId)
	}
	return
}

func VerifyEvidenceKey(evidenceKey, gateId string) (code int, err error) {
	evidenceId, err := codec.PublicDecrypt(evidenceKey)
	if err != nil {
		return 0, err
	}
	evidenceId = evidenceId[0 : len(evidenceId)-10]
	return VerifyEvidence(evidenceId, gateId)
}

func VerifyEvidence(evidenceId, gateId string) (code int, err error) {
	var evidence model.RouterEvidence
	err = dao.NewEvidenceDao().Get(evidenceId, &evidence)
	if err == sqlx.ErrNotFound {
		code = errcode.CODE_GATE_INVALID_GATE
		err = nil
		return
	} else if err != nil {
		return
	}
	if evidence.Status == EVIDENCE_STATUS_USED {
		code = errcode.CODE_GATE_USED_EVIDENCE
		return
	}
	if evidence.Status == EVIDENCE_STATUS_EXPIRED {
		code = errcode.CODE_GATE_EXPIRED_EVIDENCE
		return
	}

	var gate model.GateInfo
	err = dao.NewGateDao().Get(gateId, &gate)
	if err != nil {
		return
	}

	if evidence.Direction != 3 && gate.Direction != evidence.Direction {
		err = errcode.SGErrNotMatchGateDirection
		log4g.Error(err)
		return
	}

	//check user router
	var inRouter model.RouterInfo
	err = dao.NewRouterDao().GetIn(evidence.UserId, &inRouter)
	if gate.Direction == GATE_DIRECTION_IN {
		if inRouter.InStationId.String() != gate.StationId {
			err = errcode.SGErrDiffIn
			log4g.Error(err)
			return
		}
		//check group
		if getUserGroupNo(inRouter.UserId) != inRouter.GroupNo {
			err = errcode.SGErrExistIn
			log4g.Error(err)
			return
		}
	} else if gate.Direction == GATE_DIRECTION_OUT {
		if err != nil && err != sqlx.ErrNotFound {
			log4g.Error(err)
			return
		} else if err == sqlx.ErrNotFound {
			log4g.Error(err)
			code = errcode.CODE_GATE_ROUTER_NO_IN
			err = nil
			return
		}
	}

	//Check user pay status

	return errcode.CODE_COMMON_SUCCESS, nil
}

func GetRouterStatus(userId string) (status int8, err error) {
	var inRouter model.RouterInfo
	err = dao.NewRouterDao().GetIn(userId, &inRouter)
	if err != nil && err != sqlx.ErrNotFound {
		return -1, err
	}
	if err == sqlx.ErrNotFound {
		status = 0
	} else {
		status = sgconst.ROUTER_STATUS_NORMAL_IN
	}
	return
}

func SubmitEvidenceKey(evidenceKey string, scanTime int64, gateId string) error {
	evidenceId, err := codec.PublicDecrypt(evidenceKey)
	if err != nil {
		return err
	}
	evidenceId = evidenceId[0 : len(evidenceId)-10]
	return SubmitEvidence(evidenceId, scanTime, gateId)
}

func SubmitEvidence(evidenceId string, scanTime int64, gateId string) error {
	log4g.Debug("submitting evidence %s", evidenceId)
	var gate model.GateInfo
	err := dao.NewGateDao().Get(gateId, &gate)
	if err != nil {
		return err
	}
	log4g.Debug("gate: " + log4g.JsonString(gate))
	if err = _submitEvidence(evidenceId, scanTime, &gate); err != nil {
		//TODO submit evidence later
		return err
	}
	return nil
}

func _submitEvidence(evidenceId string, scanUnixTime int64, gateInfo *model.GateInfo) error {
	var evidence model.RouterEvidence
	err := dao.NewEvidenceDao().Get(evidenceId, &evidence)
	if err != nil {
		return err
	}
	if evidence.UserId == "" {
		log4g.Error(log4g.JsonString(evidence))
		panic("required user id")
	}
	log4g.Debug("get user %s from evidence %s", evidence.UserId, evidence.EvidenceId)
	scanTime := time.Unix(scanUnixTime, 0)
	ongoingRouter := new(model.RouterInfo)
	err = dao.NewRouterDao().GetOngoing(evidence.UserId, ongoingRouter)
	if err != nil && err != sqlx.ErrNotFound {
		log4g.Error(err)
		return err
	}
	status := ongoingRouter.Status
	if evidence.Direction == GATE_DIRECTION_IN {
		if err == sqlx.ErrNotFound { // 没有未完成的行程，正常入站
			err = in(evidence.UserId, evidenceId, &scanTime, gateInfo)
			log4g.Error(err)
			return err
		} else { //有未完成的行程
			if status.Int8() == sgconst.ROUTER_STATUS_NORMAL_IN {
				err = in(evidence.UserId, evidenceId, &scanTime, gateInfo)
				if err == nil {
					err = errcode.SGErrMoreIn
					log4g.Error(err)
				}
				return err
			} else if status.Int8() == sgconst.ROUTER_STATUS_EARLY_OUT {
				err = lateIn(ongoingRouter, evidence.UserId, evidenceId, &scanTime, gateInfo)
				if err == nil {
					err = errcode.SGErrLateIn
					log4g.Error(err)
				}
				return err
			} else {
				err = in(evidence.UserId, evidenceId, &scanTime, gateInfo)
				log4g.Error(err)
				return err
			}
		}
	} else if evidence.Direction == GATE_DIRECTION_OUT {
		if err == nil { //存在行程
			if status.Int8() == sgconst.ROUTER_STATUS_NORMAL_IN { //行程为正常入站，正常出站
				err = out(ongoingRouter, evidenceId, &scanTime, gateInfo)
				log4g.Error(err)
				return err
			} else if status.Int8() == sgconst.ROUTER_STATUS_EARLY_OUT { //已经存在提前出站的行程
				err = earlyOut(evidence.UserId, evidenceId, &scanTime, gateInfo)
				if err == nil {
					err = errcode.SGErrMoreOut
					log4g.Error(err)
				}
				return err
			} else {
				err = earlyOut(evidence.UserId, evidenceId, &scanTime, gateInfo)
				if err == nil {
					err = errcode.SGErrEarlyOut
					log4g.Error(err)
				}
				return err
			}
		} else if err == sqlx.ErrNotFound { //不存在，异常出站
			err = earlyOut(evidence.UserId, evidenceId, &scanTime, gateInfo)
			if err == nil {
				err = errcode.SGErrEarlyOut
				log4g.Error(err)
			}
			return err
		}
	} else {
		panic("wrong evidence type")
	}

	return nil
}

var userGroups = make(map[string]int)

func getUserGroupNo(userId string) int16 {
	group := 0
	if g, ok := userGroups[userId]; ok {
		group = g
	}
	now := time.Now()
	maybeGroup := now.Hour()*60 + now.Minute()
	if maybeGroup-2 > group {
		group = maybeGroup
		userGroups[userId] = group
	}
	return int16(group)
}

func in(userId, evidenceId string, inTime *time.Time, gate *model.GateInfo) (err error) {
	router := new(model.RouterInfo)
	router.UserId = userId
	router.GroupNo = getUserGroupNo(userId)
	router.InStationId.String(gate.StationId)
	router.InStationName.String(gate.StationName)
	router.InGateId.String(gate.Id)
	router.InEvidence.String(evidenceId)
	router.InTime = inTime
	router.Status.Int8(sgconst.ROUTER_STATUS_NORMAL_IN)
	err = dao.NewRouterDao().InsertIn(router)
	if err != nil {
		return err
	}
	err = CreateNotification(userId, sgconst.NOTIFICATION_TYPE_INBOUND)
	if err != nil {
		return err
	}
	return dao.NewEvidenceDao().Consume(evidenceId)
}

func lateIn(router *model.RouterInfo, userId, evidenceId string, inTime *time.Time, gateInfo *model.GateInfo) (err error) {
	router.UserId = userId
	router.InStationId.String(gateInfo.StationId)
	router.InStationName.String(gateInfo.StationName)
	router.InGateId.String(gateInfo.Id)
	router.InEvidence.String(evidenceId)
	router.InTime = inTime
	router.Status.Int8(sgconst.ROUTER_STATUS_NORMAL_OUT)
	router.Money.Float64(2)
	err = dao.NewRouterDao().UpdateIn(router)
	if err != nil {
		return err
	}
	err = ConsumeWallet(router.UserId, router.Money.Float32())
	if err != nil {
		return err
	}
	err = CreateNotification(router.UserId, sgconst.NOTIFICATION_TYPE_INBOUND)
	if err != nil {
		return err
	}
	return dao.NewEvidenceDao().Consume(evidenceId)
}

func out(router *model.RouterInfo, evidenceId string, outTime *time.Time, gateInfo *model.GateInfo) (err error) {
	router.OutStationId.String(gateInfo.StationId)
	router.OutStationName.String(gateInfo.StationName)
	router.OutGateId.String(gateInfo.Id)
	router.OutEvidence.String(evidenceId)
	router.OutTime = outTime
	router.Status.Int8(sgconst.ROUTER_STATUS_NORMAL_OUT)
	//TODO 计算费用
	router.Money.Float64(2)
	err = ConsumeWallet(router.UserId, router.Money.Float32())
	if err != nil {
		return err
	}
	router.Paid = true
	err = dao.NewRouterDao().UpdateOut(router)
	if err != nil {
		return err
	}

	err = CreateNotification(router.UserId, sgconst.NOTIFICATION_TYPE_INBOUND)
	if err != nil {
		return err
	}
	return dao.NewEvidenceDao().Consume(evidenceId)
}

func earlyOut(userId, evidenceId string, outTime *time.Time, gateInfo *model.GateInfo) (err error) {
	router := new(model.RouterInfo)
	router.UserId = userId
	router.GroupNo = getUserGroupNo(userId)
	router.OutStationId.String(gateInfo.StationId)
	router.OutStationName.String(gateInfo.StationName)
	router.OutGateId.String(gateInfo.Id)
	router.OutEvidence.String(evidenceId)
	router.OutTime = outTime
	router.Status.Int8(sgconst.ROUTER_STATUS_EARLY_OUT)
	err = dao.NewRouterDao().InsertOut(router)
	if err != nil {
		return err
	}
	err = CreateNotification(userId, sgconst.NOTIFICATION_TYPE_OUTBOUND)
	if err != nil {
		return err
	}
	return dao.NewEvidenceDao().Consume(evidenceId)
}
