package controller

import (
	"common/errcode"
	"common/model"
	"common/sqlx"
	"gate/msg"
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"etransin/dao"
)

const _IS_LOGIN_KEY = "is_login"

func gateLoginFn(agent net4g.NetAgent) {

	gateId := getGateIdFromHeader(agent)
	gateLogin := new(msg.S2CGateLogin)
	var gate model.GateInfo
	err := dao.NewGateDao().Get(gateId, &gate)
	if err == sqlx.ErrNotFound {
		gateLogin.ErrCode = errcode.CODE_GATE_INVALID_GATE
		gateLogin.ErrMsg = errcode.GetMsg(errcode.CODE_GATE_INVALID_GATE)
	} else if err != nil {
		gateLogin.ErrCode = errcode.CODE_COMMON_ERROR
		gateLogin.ErrMsg = err.Error()
	} else {
		gateLogin.GateId = gateId
		gateLogin.GateDirection = gate.Direction
		gateLogin.StationName = gate.StationName
		gateLogin.CityName = gate.CityName
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
