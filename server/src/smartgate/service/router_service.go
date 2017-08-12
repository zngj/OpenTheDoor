package service

import (
	"smartgate/dao"
	"common/util"
	"github.com/carsonsx/log4g"
	"common/model"
	"time"
	"github.com/go-redis/redis"
	"common/dbx"
	"encoding/json"
	"fmt"
	"common/errcode"
	"common/vo"
	"common/redisx"
	"common/codec"
)


const (
	GATE_DIRECTION_IN   = int8(0)
	GATE_DIRECTION_OUT  = int8(1)
	NOTIFICATION_ROUTER = "router"
)

func CreateEvidenceWithEncrypt(userId string, typ int8) (evidence *model.RouterEvidence, evidenceKey string, err error) {
	evidence, err = CreateEvidence(userId, typ)
	if err != nil {
		return
	}
	_evidenceKey := fmt.Sprintf("%s%d",evidence.EvidenceId,evidence.ExpiresAt.Unix())
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
	err = dao.NewRouterDao().InsertEvidence(evidence)
	if err == nil {
		log4g.Info("create new evidence %s for user %s", evidence.EvidenceId, evidence.UserId)
	}
	return
}


func VerifyEvidence(evidenceId, gateId string) (code int, err error) {

	var evidence *model.RouterEvidence
	evidence, err = dao.GetRouterEvidence(evidenceId)
	if err == dbx.ErrNotFound {
		code = errcode.CODE_GATE_INVALID_GATE
		err = nil
		return
	} else if err != nil {
		return
	}

	var gate *model.GateInfo
	gate, err = dao.GetGateInfo(gateId)
	if err != nil {
		return
	}

	if evidence.Direction != 3 && gate.Direction != evidence.Direction {
		code = errcode.CODE_GATE_NOT_MATCH_EVIDENCE
		log4g.Error("evidence direction is %d, but gate direction is %d", evidence.Direction, gate.Direction)
		return
	}

	//check user router
	_, err = dao.GetExceptionRouterInfo(evidence.UserId)
	if err == nil {
		code = errcode.CODE_GATE_ROUTER_EXCEPTION
		return
	} else if err != dbx.ErrNotFound {
		return
	}

	//Check user pay status

	return code, nil
}

const (
	ROUTER_STATUS_NORMAL_IN          = int8(1)
	ROUTER_STATUS_NORMAL_OUT         = int8(2)
	ROUTER_STATUS_EARLY_OUT          = int8(3)
	ROUTER_STATUS_EXCEPTION_ONLY_IN  = int8(4)
	ROUTER_STATUS_EXCEPTION_ONLY_OUT = int8(5)
)

func GetRouterStatus(userId string) (status int8, err error) {
	_, err = dao.GetOngoingRouterInfo(userId)
	if err != nil && err != dbx.ErrNotFound {
		return -1, err
	}
	if err == dbx.ErrNotFound {
		status  = 0
	} else {
		status = ROUTER_STATUS_NORMAL_IN
	}
	return
}

func SubmitEvidenceKey(evidenceKey string, scanTime int64, gateId string) error {
	evidenceId, err := codec.PublicDecrypt(evidenceKey)
	if err != nil {
		return err
	}
	evidenceId = evidenceId[0:len(evidenceId)-10]
	return SubmitEvidence(evidenceId, scanTime, gateId)
}

func SubmitEvidence(evidenceId string, scanTime int64, gateId string) error {
	log4g.Debug("submitting evidence %s", evidenceId)
	gateInfo, err := dao.GetGateInfo(gateId)
	if err != nil {
		//TODO submit evidence later
		return err
	}
	if err := _submitEvidence(evidenceId, scanTime, gateInfo); err != nil {
		//TODO submit evidence later
		return err
	}
	return nil
}

func _submitEvidence(evidenceId string , scanUnixTime int64, gateInfo *model.GateInfo) error {

	evidence, err := dao.GetRouterEvidence(evidenceId)
	if err != nil {
		return err
	}
	scanTime := time.Unix(scanUnixTime, 0)
	var ongoingRouter *model.RouterInfo
	ongoingRouter, err = dao.GetOngoingRouterInfo(evidence.UserId)
	if err != nil && err != dbx.ErrNotFound {
		log4g.Error(err)
		return err
	}
	status := ongoingRouter.Status
	if evidence.Direction == GATE_DIRECTION_IN {
		if err == dbx.ErrNotFound { // 没有未完成的行程，正常入站
			err = in(evidence.UserId, evidenceId, &scanTime, gateInfo)
			log4g.Error(err)
			return err
		} else { //有未完成的行程
			if status == ROUTER_STATUS_NORMAL_IN {
				err = in(evidence.UserId, evidenceId, &scanTime, gateInfo)
				if err == nil {
					err = errcode.SGErrMoreIn
					log4g.Error(err)
				}
				return err
			} else if status == ROUTER_STATUS_EARLY_OUT {
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
			if status == ROUTER_STATUS_NORMAL_IN { //行程为正常入站，正常出站
				err = out(ongoingRouter, evidence.UserId, evidenceId, &scanTime, gateInfo)
				log4g.Error(err)
				return err
			} else if status == ROUTER_STATUS_EARLY_OUT { //已经存在提前出站的行程
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
		} else if err == dbx.ErrNotFound { //不存在，异常出站
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

func in(userId, evidenceId string , inTime *time.Time, gateInfo *model.GateInfo) (err error) {
	router := new(model.RouterInfo)
	router.UserId = userId
	router.InStationId = gateInfo.StationId
	router.InStationName = gateInfo.StationName
	router.InGateId = gateInfo.Id
	router.InEvidence = evidenceId
	router.InTime = *inTime
	log4g.Debug(router.InTime)
	router.Status = ROUTER_STATUS_NORMAL_IN
	err = dao.InsertRouteInfoOfIn(router)
	if err != nil {
		return err
	}
	err = CreateGateInNotification(router)
	return
}

func lateIn(inRouter *model.RouterInfo, userId, evidenceId string , inTime *time.Time, gateInfo *model.GateInfo) (err error) {
	inRouter.UserId = userId
	inRouter.InStationId = gateInfo.StationId
	inRouter.InStationName = gateInfo.StationName
	inRouter.InGateId = gateInfo.Id
	inRouter.InEvidence = evidenceId
	inRouter.InTime = *inTime
	inRouter.Status = ROUTER_STATUS_NORMAL_OUT
	inRouter.Money = 2
	err = dao.UpdateRouteInfoOfIn(inRouter)
	if err != nil {
		return err
	}
	err = ConsumeWallet(inRouter.UserId, inRouter.Money)
	if err != nil {
		return err
	}
	return CreateGateOutNotification(inRouter)
}


func out(outRouter *model.RouterInfo, userId, evidenceId string , outTime *time.Time, gateInfo *model.GateInfo) (err error) {
	outRouter.OutStationId = gateInfo.StationId
	outRouter.OutStationName = gateInfo.StationName
	outRouter.OutGateId = gateInfo.Id
	outRouter.OutEvidence = evidenceId
	outRouter.OutTime = *outTime
	outRouter.Status = ROUTER_STATUS_NORMAL_OUT
	//TODO 计算费用
	outRouter.Money = 2
	err = dao.UpdateRouteInfoOfOut(outRouter)
	if err != nil {
		return err
	}
	err = ConsumeWallet(outRouter.UserId, outRouter.Money)
	if err != nil {
		return err
	}
	return CreateGateOutNotification(outRouter)
}

func earlyOut(userId, evidenceId string , outTime *time.Time, gateInfo *model.GateInfo) (err error) {
	router := new(model.RouterInfo)
	router.UserId = userId
	router.OutStationId = gateInfo.StationId
	router.OutStationName = gateInfo.StationName
	router.OutGateId = gateInfo.Id
	router.OutEvidence = evidenceId
	router.OutTime = *outTime
	router.Status = ROUTER_STATUS_EARLY_OUT
	err = dao.InsertRouteInfoOfOut(router)
	if err != nil {
		return err
	}
	return CreateGateOutNotification(router)
}


func CreateGateInNotification(router *model.RouterInfo) error {
	nvo := new(vo.RouterNotificationVo)
	nvo.Direction = GATE_DIRECTION_IN
	nvo.InGateId = router.InGateId
	nvo.InStationId = router.InStationId
	nvo.InStationName = router.InStationName
	nvo.InTime = router.InTime.Unix()
	return PublishNotification(router.UserId, NOTIFICATION_ROUTER, log4g.ToJsonString(nvo))
}

func CreateGateOutNotification(router *model.RouterInfo) error {
	nvo := new(vo.RouterNotificationVo)
	nvo.Direction = GATE_DIRECTION_OUT
	nvo.InGateId = router.InGateId
	nvo.InStationId = router.InStationId
	nvo.InStationName = router.InStationName
	nvo.InTime = router.InTime.Unix()
	nvo.OutGateId = router.OutGateId
	nvo.OutStationId = router.OutStationId
	nvo.OutStationName = router.OutStationName
	nvo.OutTime = router.OutTime.Unix()
	nvo.Money = router.Money
	return PublishNotification(router.UserId, NOTIFICATION_ROUTER, log4g.ToJsonString(nvo))
}

func getNotificationKey(category, contentId string) string {
	return fmt.Sprintf("notification:%s:%s", category, contentId)
}

func PublishNotification(userId, category, content string) error {
	contentId := util.NewUuid()
	key := getNotificationKey(category, contentId)
	err := redisx.Client.Set(key, content, 0).Err()
	if err != nil {
		log4g.Error(err)
		return err
	}
	notification := new(model.Notification)
	notification.UserId = userId
	notification.Category = category
	notification.ContentId = contentId
	err = dao.InsertNotification(notification)
	if err != nil {
		redisx.Client.Del(key)
	} else {
		log4g.Debug("publish %s notification for user[%s]\n%s", category, userId, content)
	}
	return err
}

func GetRouterNotification(userId string) (nvo *vo.RouterNotificationVo, err error) {
	var notification *model.Notification
	notification, err = dao.GetNotification(userId, NOTIFICATION_ROUTER)
	if err != nil {
		return
	}
	key := getNotificationKey(notification.Category, notification.ContentId)
	var jstrNVO string
	jstrNVO, err = redisx.Client.Get(key).Result()
	if err != nil {
		if redis.Nil == err {
			log4g.Error("redis nil for key [%s]", key)
			dao.ConsumeNotification(notification.Id)
		}
		err = dbx.ErrNotFound
		return
	}
	log4g.Debug("notification content: %s", jstrNVO)
	nvo = new (vo.RouterNotificationVo)
	err = json.Unmarshal([]byte(jstrNVO), nvo)
	if err == nil {
		nvo.NotificationId = notification.Id
		log4g.Debug(nvo)
	} else {
		log4g.Error(err)
		log4g.Error(jstrNVO)
	}
	return
}
