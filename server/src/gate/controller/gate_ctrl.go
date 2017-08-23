package controller

import (
	"common/codec"
	"common/errcode"
	"gate/msg"
	"github.com/carsonsx/net4g"
	"etransin/service"
	"time"
)

func verifyEvidenceFn(agent net4g.NetAgent) {
	if !checkLogin(agent) {
		return
	}
	evidence := agent.Msg().(*msg.C2SVerifyEvidence)
	evidenceKey, err := codec.PublicDecrypt(evidence.EvidenceKey)
	verifyResult := new(msg.S2CVerifyEvidence)
	if err != nil {
		verifyResult.ErrCode = errcode.CODE_COMMON_ERROR
		verifyResult.ErrMsg = err.Error()
		return
	}
	evidenceId := evidenceKey[0 : len(evidenceKey)-10]
	write(agent, verifyEvidence(evidenceId, getGateId(agent)))
}

func verifyEvidence(evidenceId, gateId string) (verifyResult *msg.S2CVerifyEvidence) {
	verifyResult = new(msg.S2CVerifyEvidence)
	verifyResult.VerifyTime = time.Now().Unix()
	if evidenceId == "" {
		verifyResult.ErrCode = errcode.CODE_COMMON_EMPTY_ARG
		verifyResult.ErrMsg = errcode.GetMsg(verifyResult.ErrCode)
		return
	}

	if len(evidenceId) != 32 {
		verifyResult.ErrCode = errcode.CODE_GATE_INVALID_EVIDENCE
		verifyResult.ErrMsg = errcode.GetMsg(verifyResult.ErrCode)
		return
	}
	var err error
	verifyResult.ErrCode, err = service.VerifyEvidence(evidenceId, gateId)
	if err != nil {
		verifyResult.ErrCode = errcode.CODE_COMMON_ERROR
		verifyResult.ErrMsg = err.Error()
	} else {
		if verifyResult.ErrCode > 0 {
			verifyResult.ErrMsg = errcode.GetMsg(verifyResult.ErrCode)
		}
	}
	return verifyResult
}

func submitEvidenceFn(agent net4g.NetAgent) {
	if !checkLogin(agent) {
		return
	}

	evidence := agent.Msg().(*msg.C2SSubmitEvidence)
	evidenceKey, err := codec.PublicDecrypt(evidence.EvidenceKey)
	verifyResult := new(msg.S2CVerifyEvidence)
	if err != nil {
		verifyResult.ErrCode = errcode.CODE_COMMON_ERROR
		verifyResult.ErrMsg = err.Error()
		return
	}
	evidenceId := evidenceKey[0 : len(evidenceKey)-10]

	result := new(msg.S2CSubmitEvidence)
	gateId := getGateId(agent)
	verifyResult = verifyEvidence(evidenceId, gateId)
	if verifyResult.ErrCode > 0 {
		result.ErrCode = verifyResult.ErrCode
		result.ErrMsg = verifyResult.ErrMsg
	} else {
		// TODO save to redis then go ?
		go service.SubmitEvidence(evidenceId, evidence.ScanTime, gateId)
	}
	write(agent, result)
}
