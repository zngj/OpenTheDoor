package controller

import (
	"github.com/carsonsx/net4g"
	"gate/msg"
	"smartgate/service"
	"common/errcode"
	"smartgate/codec"
)

func verifyEvidenceFn(agent net4g.NetAgent)  {
	if !checkLogin(agent) {
		return
	}
	evidence := agent.Msg().(*msg.C2SVerifyEvidence)

	write(agent, verifyEvidence(evidence.EvidenceKey, getGateId(agent)))
}

func verifyEvidence(evidenceEncryptKey, gateId string) (verifyResult *msg.S2CVerifyEvidence) {
	verifyResult = new(msg.S2CVerifyEvidence)
	if evidenceEncryptKey == "" {
		verifyResult.ErrCode = errcode.CODE_COMMON_EMPTY_ARG
		verifyResult.ErrMsg = errcode.GetMsg(verifyResult.ErrCode)
		return
	}
	evidenceKey, err := codec.PublicDecrypt(evidenceEncryptKey)
	if err != nil {
		verifyResult.ErrCode = errcode.CODE_COMMON_ERROR
		verifyResult.ErrMsg = err.Error()
		return
	}
	if len(evidenceKey) != 42 {
		verifyResult.ErrCode = errcode.CODE_GATE_INVALID_EVIDENCE
		verifyResult.ErrMsg = errcode.GetMsg(verifyResult.ErrCode)
		return
	}
	evidenceId := evidenceKey[0:len(evidenceKey)-10]
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

func submitEvidenceFn(agent net4g.NetAgent)  {
	if !checkLogin(agent) {
		return
	}
	evidence := agent.Msg().(*msg.C2SSubmitEvidence)
	result := new(msg.S2CSubmitEvidence)
	gateId := getGateId(agent)
	verifyResult := verifyEvidence(evidence.EvidenceKey, gateId)
	if verifyResult.ErrCode > 0 {
		result.ErrCode = verifyResult.ErrCode
		result.ErrMsg = verifyResult.ErrMsg
	} else {
		// TODO save to redis then go ?
		go service.SubmitEvidence(evidence, gateId)
	}
	write(agent, result)
}
