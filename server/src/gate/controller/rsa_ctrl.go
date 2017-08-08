package controller

import (
	"github.com/carsonsx/net4g"
	"gate/msg"
	"smartgate/codec"
)

func rsaKeyFn(agent net4g.NetAgent)  {
	if !_checkLogin(agent) {
		return
	}
	rsaKey := new(msg.S2CRsaKey)
	rsaKey.Key = string(codec.Private_Key)
	write(agent, rsaKey)
}