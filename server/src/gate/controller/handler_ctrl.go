package controller

import (
	"gate/msg"
	"github.com/carsonsx/net4g"
)

var Dispatcher = net4g.NewDispatcher("gate")

func OnInit() {
	Dispatcher.OnConnectionCreated(onConnectionCreated)
	Dispatcher.AddHandler(gateLoginFn, msg.GATE_LOGIN)
	Dispatcher.AddHandler(rsaKeyFn, msg.RSA_KEY)
	Dispatcher.AddHandler(verifyEvidenceFn, msg.VERIFY_EVIDENCE)
	Dispatcher.AddHandler(submitEvidenceFn, msg.SUBMIT_EVIDENCE)
}

func onConnectionCreated(agent net4g.NetAgent) {
	//
}

func getSGHeader(agent net4g.NetAgent) *msg.SGHeader {
	return agent.Header().(*msg.SGHeader)
}

func getGateIdFromHeader(agent net4g.NetAgent) string {
	return getSGHeader(agent).GateId
}

func getGateId(agent net4g.NetAgent) string {
	return agent.Session().Key()
}

func write(agent net4g.NetAgent, v interface{}) {
	agent.Write(v, msg.NewSGHeader(getGateId(agent)))
}
