package test_test

import (
	"gate/msg"
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"testing"
	"time"
)

var gateId = "010100101"

func TestLogin(t *testing.T) {
	start(login)
}

func login()  {
	dispatcher.AddHandler(loginResult, msg.GATE_LOGIN)
	header := msg.NewSGHeader(gateId)
	rp := new(net4g.RawPack)
	rp.Id = msg.GATE_LOGIN
	agent.Write(rp, header)
}

func loginResult(agent net4g.NetAgent) {

	if gl := agent.Msg().(*msg.S2CGateLogin); gl.ErrCode == 0 {
		log4g.Info("[client]gate %s login success", gl.GateId)
		log4g.JsonDebug(agent.Msg())
		net4g.TestDone()
	}
}

func TestNotLogin(t *testing.T) {
	start(notLogin)
}

func notLogin()  {
	rsaKey()
}

func TestRsaKey(t *testing.T) {
	start(login, rsaKey)
}

func rsaKey() {
	dispatcher.AddHandler(rsaKeyResult, msg.RSA_KEY)
	header := msg.NewSGHeader(gateId)
	rp := new(net4g.RawPack)
	rp.Id = msg.RSA_KEY
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
	dispatcher.AddHandler(verifyEvidenceResult, msg.VERIFY_EVIDENCE)
	header := msg.NewSGHeader(gateId)
	ve := new(msg.C2SVerifyEvidence)
	ve.EvidenceKey = "GoQSYpFTFiN/bfwj19iMpQlR/ilzBFYaNn2i2EejPyGfQFxGAhZic69Jn4yMeV0ohcba3Tvn1Dv2CyIK/eOG9A5eir9V10ZVk5j60wOhV4qMJ8QiHxqjYCbFHUAivF0H8l10mR3rU4QJkD9iymFBT7jF3uMp+qMox/p541bxRHg="
	agent.Write(ve, header)
}

func verifyEvidenceResult(agent net4g.NetAgent) {
	result := agent.Msg().(*msg.S2CVerifyEvidence)
	log4g.Debug("* verify evidence result: [%d]%s", result.ErrCode, result.ErrMsg)
	net4g.TestDone()
}

func TestSubmitEvidence(t *testing.T) {
	//gateId = "010100202"
	start(login, submitEvidence)
}

func submitEvidence() {
	dispatcher.AddHandler(submitEvidenceResult, msg.SUBMIT_EVIDENCE)
	header := msg.NewSGHeader(gateId)
	ue := new(msg.C2SSubmitEvidence)
	ue.EvidenceKey = "RQ6nROMjHQ4TgWRhMZyqM5wPh2/hgaw6Et8SyTbJ2yqMgjQUAy/q1Bz8yqyXstZGa8oI2oEJs9koxzyHBf+I06jo22CqXOmJFwvX+JaFY8XlgQ+7eCa3zOn9NPXSnGJxcVUFD+20bQJRhti4T7dhsZT+y0/lT6ZNKSsnOEqjFbE="
	ue.ScanTime = time.Now().Unix()
	agent.Write(ue, header)
}

func submitEvidenceResult(agent net4g.NetAgent) {
	log4g.Debug("* upload user evidence result: %v", agent.Msg().(*msg.S2CSubmitEvidence).ErrCode)
	net4g.TestDone()
}