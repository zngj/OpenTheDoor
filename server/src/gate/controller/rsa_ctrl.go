package controller

import (
	"common/codec"
	"gate/msg"
	"github.com/carsonsx/net4g"
)

func rsaKeyFn(agent net4g.NetAgent) {
	if !checkLogin(agent) {
		return
	}
	rsaKey := new(msg.S2CRsaKey)
	rsaKey.Key = string(codec.Public_Key)
	write(agent, rsaKey)
}
