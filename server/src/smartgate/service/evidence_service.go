package service

import (
	"common/dbx"
	"common/errcode"
	"common/model"
	"common/redisx"
	"common/util"
	"common/vo"
	"gate/msg"
	"github.com/carsonsx/log4g"
	"smartgate/dao"
	"fmt"
	"github.com/go-redis/redis"
	"encoding/json"
	"time"
)

const (
	GATE_DIRECTION_IN  = int8(0)
	GATE_DIRECTION_OUT = int8(1)
	NOTFICATION_ROUTER = "router"
)

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
	ROUTER_STATUS_IN                 = int8(1)
	ROUTER_STATUS_OUT                = int8(2)
	ROUTER_STATUS_ONLY_OUT           = int8(3)
	ROUTER_STATUS_EXCEPTION_ONLY_IN  = int8(4)
	ROUTER_STATUS_EXCEPTION_ONLY_OUT = int8(5)
)

func SubmitEvidence(submitEvidence *msg.C2SSubmitEvidence, gateId string) {
	log4g.Debug("submitting evidence %s", submitEvidence.EvidenceKey)
	gateInfo, err := dao.GetGateInfo(gateId)
	if err != nil {
		//TODO submit evidence later
		return
	}
	if err := _submitEvidence(submitEvidence, gateInfo); err != nil {
		//TODO submit evidence later
	}
}

func _submitEvidence(submitEvidence *msg.C2SSubmitEvidence, gateInfo *model.GateInfo) error {

	evidenceId := submitEvidence.EvidenceKey[0:len(submitEvidence.EvidenceKey)-10]
	evidence, err := dao.GetRouterEvidence(evidenceId)
	if err != nil {
		return err
	}

	scanTime := time.Unix(submitEvidence.ScanTime, 0)
	var ongoingRouter *model.RouterInfo
	ongoingRouter, err = dao.GetOngoingRouterInfo(evidence.UserId)

	if evidence.Direction == GATE_DIRECTION_IN {
		if err == dbx.ErrNotFound {
			router := new(model.RouterInfo)
			router.UserId = evidence.UserId
			router.InStationId = gateInfo.StationId
			router.InStationName = gateInfo.StationName
			router.InGateId = gateInfo.Id
			router.InEvidence = evidenceId
			router.InTime = scanTime
			log4g.Debug(router.InTime)
			router.Status = ROUTER_STATUS_IN
			err = dao.InsertRouteInfoOfIn(router)
			//TODO 系统消息通知
			err = CreateGateInNotification(router)
			if err != nil {
				//TODO 扣费异常
				return err
			}
		} else  if err == nil { //存在
			ongoingRouter.UserId = evidence.UserId
			ongoingRouter.OutStationId = gateInfo.StationId
			ongoingRouter.OutStationName = gateInfo.StationName
			ongoingRouter.OutGateId = gateInfo.Id
			ongoingRouter.OutEvidence = evidenceId
			ongoingRouter.OutTime = scanTime
			ongoingRouter.Status = ROUTER_STATUS_OUT
			err = dao.UpdateRouteInfoOfIn(ongoingRouter)

			//TODO 系统消息通知
			err = CreateGateInNotification(ongoingRouter)
			if err != nil {
				//TODO 扣费异常
				return err
			}
		}

		if err != nil {
			//TODO 添加等下异常
			return err
		}

	} else if evidence.Direction == GATE_DIRECTION_OUT {
		if err == nil { //存在
			ongoingRouter.OutStationId = gateInfo.StationId
			ongoingRouter.OutStationName = gateInfo.StationName
			ongoingRouter.OutGateId = gateInfo.Id
			ongoingRouter.OutEvidence = evidenceId
			ongoingRouter.OutTime = scanTime
			ongoingRouter.Status = ROUTER_STATUS_OUT
			//TODO 计算费用
			ongoingRouter.Money = 2
			err = dao.UpdateRouteInfoOfOut(ongoingRouter)
			if err == nil {
				//TODO 检查余额并扣费
				err = ConsumeWallet(ongoingRouter.UserId, ongoingRouter.Money)
				if err != nil {
					//TODO 扣费异常
					return err
				}
				//TODO 系统消息通知
				err = CreateGateOutNotification(ongoingRouter)
				if err != nil {
					//TODO 扣费异常
					return err
				}
			}
		} else if err == dbx.ErrNotFound { //不存在
			router := new(model.RouterInfo)
			router.UserId = evidence.UserId
			router.OutStationId = gateInfo.StationId
			router.OutStationName = gateInfo.StationName
			router.OutGateId = gateInfo.Id
			router.OutEvidence = evidenceId
			router.OutTime = scanTime
			router.Status = ROUTER_STATUS_ONLY_OUT
			err = dao.InsertRouteInfoOfOut(router)
		}
	} else {
		panic("wrong evidence type")
	}

	if err != nil {
		log4g.Info("handle evidence success.")
	}

	return err
}

func CreateGateInNotification(router *model.RouterInfo) error {
	nvo := new(vo.RouterNotificationVo)
	nvo.Direction = GATE_DIRECTION_IN
	nvo.InGateId = router.InGateId
	nvo.InStationId = router.InStationId
	nvo.InStationName = router.InStationName
	nvo.InTime = router.InTime.Unix()
	return PublishNotification(router.UserId, NOTFICATION_ROUTER, log4g.ToJsonString(nvo))
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
	return PublishNotification(router.UserId, NOTFICATION_ROUTER, log4g.ToJsonString(nvo))
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
	notification, err = dao.GetNotification(userId, NOTFICATION_ROUTER)
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
