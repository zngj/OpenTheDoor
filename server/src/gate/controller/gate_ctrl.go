package controller

import (
	"github.com/carsonsx/net4g"
	"gate/msg"
	"smartgate/service"
	"common/errcode"
)

func verifyEvidenceFn(agent net4g.NetAgent)  {
	if !checkLogin(agent) {
		return
	}
	evidence := agent.Msg().(*msg.C2SVerifyEvidence)
	write(agent, verifyEvidence(evidence.EvidenceId, getGateId(agent)))
}

func verifyEvidence(evidenceId, gateId string) *msg.S2CVerifyEvidence {
	verifyResult := new(msg.S2CVerifyEvidence)
	if evidenceId == "" {
		verifyResult.ErrCode = errcode.CODE_COMMON_EMPTY_ARG
		verifyResult.ErrMsg = errcode.GetMsg(verifyResult.ErrCode)
	} else if len(evidenceId) != 32 {
		verifyResult.ErrCode = errcode.CODE_GATE_INVALID_EVIDENCE
		verifyResult.ErrMsg = errcode.GetMsg(verifyResult.ErrCode)
	} else {
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
	verifyResult := verifyEvidence(evidence.EvidenceId, gateId)
	if verifyResult.ErrCode > 0 {
		result.ErrCode = verifyResult.ErrCode
		result.ErrMsg = verifyResult.ErrMsg
	} else {
		// TODO save to redis then go ?
		go service.SubmitEvidence(evidence, gateId)
	}
	write(agent, result)
}
