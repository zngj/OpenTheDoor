package main

import (
	"github.com/carsonsx/net4g"
	"gate/msg"
)

func main()  {

	net4g.NewTcpServer("gate", ":8083").
		EnableHeartbeat().
		SetSerializer(msg.Serializer).
		AddDispatchers(msg.Dispatcher).
		Start().
		Wait()

}
