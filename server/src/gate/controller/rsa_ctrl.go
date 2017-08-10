package controller

import (
	"github.com/carsonsx/net4g"
	"gate/msg"
	"common/codec"
)

func rsaKeyFn(agent net4g.NetAgent)  {
	if !checkLogin(agent) {
		return
	}
	rsaKey := new(msg.S2CRsaKey)
	rsaKey.Key = string(codec.Private_Key)
	write(agent, rsaKey)
}