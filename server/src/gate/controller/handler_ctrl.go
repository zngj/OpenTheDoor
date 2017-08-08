package controller

import (
	"gate/msg"
	"github.com/carsonsx/net4g"
)

var Dispatcher = net4g.NewDispatcher("gate", 10)

func OnInit() {

	Dispatcher.OnConnectionCreated(onConnectionCreated)

	Dispatcher.AddHandler(gateLoginFn, msg.C2S_GATE_LOGIN)
	Dispatcher.AddHandler(rsaKeyFn, msg.C2S_RSA_KEY)
	Dispatcher.AddHandler(verifyEvidenceFn, msg.C2S_VERIFY_EVIDENCE)
	Dispatcher.AddHandler(submitEvidenceFn, msg.C2S_SUBMIT_EVIDENCE)

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

func getGateIdFromSession(agent net4g.NetAgent) string {
	return agent.Session().Key()
}

func write(agent net4g.NetAgent, v interface{}) {
	agent.Write(v, msg.NewSGHeader(getGateIdFromSession(agent)))
}