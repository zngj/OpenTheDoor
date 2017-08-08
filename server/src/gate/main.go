package main

import (
	"gate/msg"
	"github.com/carsonsx/net4g"
	"gate/controller"
)

func main() {

	net4g.NetConfig.BeginBytes = []byte{0x10, 0x02}
	net4g.NetConfig.EndBytes = []byte{0x10, 0x03}

	msg.OnInit()
	controller.OnInit()

	net4g.NewTcpServer("gate", ":8083").
		//EnableHeartbeat().
		SetSerializer(msg.Serializer).
		AddDispatchers(controller.Dispatcher).
		Start().
		Wait()

}
