package controller

import (
	"github.com/carsonsx/net4g"
	"gate/msg"
	"gate/service"
	"common/errcode"
)

func verifyEvidenceFn(agent net4g.NetAgent)  {
	if !checkLogin(agent) {
		return
	}
	evidence := agent.Msg().(*msg.C2SVerifyEvidence)
	verifyResult := new(msg.S2CVerifyEvidence)
	if evidence.EvidenceId == "" {
		verifyResult.ErrCode = errcode.CODE_COMMON_EMPTY_ARG
		verifyResult.ErrMsg = errcode.GetMsg(verifyResult.ErrCode)
	} else if len(evidence.EvidenceId) != 32 {
		verifyResult.ErrCode = errcode.CODE_GATE_INVALID_EVIDENCE
		verifyResult.ErrMsg = errcode.GetMsg(verifyResult.ErrCode)
	} else {
		var err error
		verifyResult.ErrCode, err = service.VerifyEvidence(evidence.EvidenceId, getGateId(agent))
		if err != nil {
			verifyResult.ErrCode = errcode.CODE_COMMON_ERROR
			verifyResult.ErrMsg = err.Error()
		} else {
			if verifyResult.ErrCode > 0 {
				verifyResult.ErrMsg = errcode.GetMsg(verifyResult.ErrCode)
			}
		}
	}
	write(agent, verifyResult)
}

func submitEvidenceFn(agent net4g.NetAgent)  {
	if !checkLogin(agent) {
		return
	}




	result := new(msg.S2CSubmitEvidence)
	write(agent, result)
}
