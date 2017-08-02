package service

import (
	"gate/msg"
	"github.com/carsonsx/net4g"
	"github.com/carsonsx/log4g"
)

const GATE_NO_KEY  = "gate_no"

func init()  {
	msg.Dispatcher.AddHandler(setGateNo, msg.SetGateNoType)
	msg.Dispatcher.AddHandler(userIn, msg.UserInType)
	msg.Dispatcher.AddHandler(userOut, msg.UserOutType)
	msg.Dispatcher.OnConnectionClosed(onConnectionClosed)
}

func onConnectionClosed(session net4g.NetSession) {
	log4g.Warn("the gate  was offline!", session.GetString(GATE_NO_KEY))
}

func setGateNo(agent net4g.NetAgent)  {
	gate := agent.Msg().(*msg.SetGateNo)
	agent.Session().Set(GATE_NO_KEY, gate.No)
	agent.Write(new(msg.SetGateNoSuccess))
}

func userIn(agent net4g.NetAgent)  {
	agent.Write(new(msg.UserInSuccess))
}

func userOut(agent net4g.NetAgent)  {
	agent.Write(new(msg.UserOutSuccess))
}