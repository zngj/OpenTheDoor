package test_test

import (
	"gate/msg"
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"testing"
	"common/util"
	"time"
)

var gateId = "010100101"

func TestLogin(t *testing.T) {
	start(login)
}

func login()  {
	dispatcher.AddHandler(loginResult, msg.S2C_GATE_LOGIN)
	header := msg.NewSGHeader(gateId)
	rp := new(net4g.RawPack)
	rp.Id = msg.C2S_GATE_LOGIN
	agent.Write(rp, header)
}

func loginResult(agent net4g.NetAgent) {

	if gl := agent.Msg().(*msg.S2CGateLogin); gl.ErrCode == 0 {
		log4g.Info("gate %s login success", gl.GateId)
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
	dispatcher.AddHandler(rsaKeyResult, msg.S2C_RSA_KEY)
	header := msg.NewSGHeader(gateId)
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
	header := msg.NewSGHeader(gateId)
	ve := new(msg.C2SVerifyEvidence)
	ve.EvidenceId = "0d61539e6b954dc5943fd2f1a33819aa"
	agent.Write(ve, header)
}

func verifyEvidenceResult(agent net4g.NetAgent) {
	result := agent.Msg().(*msg.S2CVerifyEvidence)
	log4g.Debug("* verify evidence result: [%d]%s", result.ErrCode, result.ErrMsg)
	net4g.TestDone()
}

func TestSubmitEvidence(t *testing.T) {
	gateId = "010100202"
	start(login, submitEvidence)
}

func submitEvidence() {
	dispatcher.AddHandler(submitEvidenceResult, msg.S2C_SUBMIT_EVIDENCE)
	header := msg.NewSGHeader(gateId)
	ue := new(msg.C2SSubmitEvidence)
	ue.EvidenceId = "ff2ec54dc7434237959b661c472584da"
	ue.ScanTime = util.TimeToUnixMilli(time.Now())
	agent.Write(ue, header)
}

func submitEvidenceResult(agent net4g.NetAgent) {
	log4g.Debug("* upload user evidence result: %v", agent.Msg().(*msg.S2CSubmitEvidence).ErrCode)
	net4g.TestDone()
}