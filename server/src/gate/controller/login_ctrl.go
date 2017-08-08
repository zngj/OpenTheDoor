package controller

import (
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/log4g"
	"gate/msg"
)

const _IS_LOGIN_KEY  = "is_login"

func gateLoginFn(agent net4g.NetAgent)  {

	log4g.Debug(agent.Header())
	log4g.Debug(agent.Msg())

	gateId := getGateIdFromHeader(agent)

	gateLogin := new(msg.S2CGateLogin)
	if gateId == "010100101" || gateId == "010100202" {
		gateLogin.Code = 0
		if gateId == "010100101" {
			gateLogin.GateDirection = "in"
			gateLogin.StationName = "五一广场"
			gateLogin.CityName = "长沙"
		} else if gateId == "010100202" {
			gateLogin.GateDirection = "out"
			gateLogin.StationName = "黄兴广场"
			gateLogin.CityName = "长沙"
		}
		agent.Key(gateId)
		agent.Session().Set(_IS_LOGIN_KEY, true)
		log4g.Info("gate %s login success", gateId)
	} else {
		log4g.Error("gate %s is not existed", gateId)
		gateLogin.Code = 1 //gate id不存在
	}

	//response
	header := msg.NewSGHeader(gateId)
	header.GateId = gateId
	agent.Write(gateLogin, header)

}

func _checkLogin(agent net4g.NetAgent) bool {
	if !agent.Session().GetBool(_IS_LOGIN_KEY) {
		agent.Write(net4g.NewRawPackById(msg.S2C_NOT_LOGIN), msg.NewSGHeader(getGateIdFromHeader(agent)))
		return false
	}
	return true
}

