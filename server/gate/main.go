package main

import (
	"gate/controller"
	"gate/msg"
	"github.com/carsonsx/log4g"
	"github.com/carsonsx/net4g"
	"runtime/debug"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			log4g.Error("********************* Message Handler Panic *********************")
			log4g.Error(r)
			log4g.Error(string(debug.Stack()))
			log4g.Error("********************* Message Handler Panic *********************")
		}
	}()

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
