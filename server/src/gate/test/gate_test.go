package test_test

import (
	"gate/msg"
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"testing"
)

func TestLogin(t *testing.T) {
	start(login)
}

func login()  {
	dispatcher.AddHandler(loginResult, msg.S2C_GATE_LOGIN)
	header := msg.NewSGHeader("010100101")
	rp := new(net4g.RawPack)
	rp.Id = msg.C2S_GATE_LOGIN
	agent.Write(rp, header)
}

func loginResult(agent net4g.NetAgent) {
	log4g.Debug(agent.Msg())
	if agent.Msg().(*msg.S2CGateLogin).Code == 0 {
		net4g.TestDone()
	}
}

func TestNotLogin(t *testing.T) {
	start(netLogin)
}

func netLogin()  {
	dispatcher.AddHandler(notLoginResult, msg.S2C_NOT_LOGIN)
	rsaKey()
}

func notLoginResult(agent net4g.NetAgent) {
	log4g.Debug("* gate not login")
	net4g.TestDone()
}

func TestRsaKey(t *testing.T) {
	start(login, rsaKey)
}

func rsaKey() {
	dispatcher.AddHandler(rsaKeyResult, msg.S2C_RSA_KEY)
	header := msg.NewSGHeader("010100101")
	rp := new(net4g.RawPack)
	rp.Id = msg.C2S_RSA_KEY
	agent.Write(rp, header)
}

func rsaKeyResult(agent net4g.NetAgent) {
	log4g.Debug(agent.Msg().(*msg.S2CRsaKey).Key)
	net4g.TestDone()
}

func TestVerifyEvidence(t *testing.T) {
	start(login, verifyEvidence)
}

func verifyEvidence() {
	dispatcher.AddHandler(verifyEvidenceResult, msg.S2C_VERIFY_EVIDENCE)
	header := msg.NewSGHeader("010100101")
	ve := new(msg.C2SVerifyEvidence)
	ve.EvidenceKey = "1111"
	agent.Write(ve, header)
}

func verifyEvidenceResult(agent net4g.NetAgent) {
	log4g.Debug("* verify evidence result: %d", agent.Msg().(*msg.S2CVerifyEvidence).Code)
	net4g.TestDone()
}

func TestUserEvidence(t *testing.T) {
	start(login, userEvidence)
}

func userEvidence() {
	dispatcher.AddHandler(userEvidenceResult, msg.S2C_USER_EVIDENCE)
	header := msg.NewSGHeader("010100101")
	ue := new(msg.C2SUserEvidence)
	ue.EvidenceKey = "1111"
	agent.Write(ue, header)
}

func userEvidenceResult(agent net4g.NetAgent) {
	log4g.Debug("* upload user evidence result: %v", agent.Msg().(*msg.S2CUserEvidence).Success)
	net4g.TestDone()
}