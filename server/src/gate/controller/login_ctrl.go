package controller

import (
	"github.com/carsonsx/net4g"
	"gate/msg"
	"common/dbx"
	"common/errcode"
	"github.com/carsonsx/log4g"
	"smartgate/dao"
)

const _IS_LOGIN_KEY  = "is_login"

func gateLoginFn(agent net4g.NetAgent)  {

	gateId := getGateIdFromHeader(agent)
	gateLogin := new(msg.S2CGateLogin)

	gateInfo, err := dao.GetGateInfo(gateId)
	if err == dbx.ErrNotFound {
		gateLogin.ErrCode = errcode.CODE_GATE_INVALID_GATE
		gateLogin.ErrMsg = errcode.GetMsg(errcode.CODE_GATE_INVALID_GATE)
	} else if err != nil {
		gateLogin.ErrCode = errcode.CODE_COMMON_ERROR
		gateLogin.ErrMsg = err.Error()
	} else {
		gateLogin.GateId = gateId
		gateLogin.GateDirection = gateInfo.Direction
		gateLogin.StationName = gateInfo.StationName
		gateLogin.CityName = gateInfo.CityName
		agent.Session().Set(_IS_LOGIN_KEY, true)
		agent.Key(gateId)
		log4g.Info("* gate %s login success", gateId)
	}

	//response
	header := msg.NewSGHeader(gateId)
	header.GateId = gateId
	agent.Write(gateLogin, header)

}

func checkLogin(agent net4g.NetAgent) bool {
	if !agent.Session().GetBool(_IS_LOGIN_KEY) {
		agent.Write(net4g.NewRawPack(msg.NOT_LOGIN), msg.NewSGHeader(getGateIdFromHeader(agent)))
		return false
	}
	return true
}

