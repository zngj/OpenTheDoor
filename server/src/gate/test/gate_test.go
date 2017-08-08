package test_test

import (
	"gate/msg"
	"github.com/carsonsx/net4g"
	"testing"
	"github.com/carsonsx/log4g"
)


func gateLoginSuccess(agent net4g.NetAgent)  {

	log4g.Debug(agent.Msg())

}


func TestLogin(t *testing.T) {

	dispatcher.AddHandler(gateLoginSuccess, msg.S2C_GATE_LOGIN)

	connect(func(agent net4g.NetAgent) {
		header := msg.NewSGHeader("010100101")
		rp := new(net4g.RawPack)
		rp.Id = msg.C2S_GATE_LOGIN
		agent.Write(rp, header)
	})

}
