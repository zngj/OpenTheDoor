package controller

import (
	"github.com/carsonsx/net4g"
	"gate/msg"
)

func verifyEvidenceFn(agent net4g.NetAgent)  {
	if !checkLogin(agent) {
		return
	}
	verify := new(msg.S2CVerifyEvidence)
	verify.Code = 0
	write(agent, verify)
}

func userEvidenceFn(agent net4g.NetAgent)  {
	if !checkLogin(agent) {
		return
	}
	result := new(msg.S2CUserEvidence)
	result.Success = true
	write(agent, result)
}
