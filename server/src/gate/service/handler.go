package service

import (
	"gate/msg"
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/log4g"
)

const GATE_NO_KEY  = "gate_no"

func init()  {
	msg.Dispatcher.AddHandler(setGateNo, msg.RsaKeyC2SType)
	msg.Dispatcher.AddHandler(userIn, msg.UserInC2SType)
	msg.Dispatcher.AddHandler(userOut, msg.UserOutC2SType)
	msg.Dispatcher.OnConnectionClosed(onConnectionClosed)
}

func onConnectionClosed(agent net4g.NetAgent) {
	log4g.Warn("the gate  was offline!", agent.Session().GetString(GATE_NO_KEY))
}

func setGateNo(agent net4g.NetAgent)  {
	gate := agent.Msg().(*msg.RsaKeyC2S)
	agent.Session().Set(GATE_NO_KEY, gate.No)
	agent.Write(new(msg.RsaKeyS2C))
}

func userIn(agent net4g.NetAgent)  {
	agent.Write(new(msg.UserInS2C))
}

func userOut(agent net4g.NetAgent)  {
	agent.Write(new(msg.UserOutS2C))
}