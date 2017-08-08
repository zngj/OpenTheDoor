package test_test

import (
	"github.com/carsonsx/net4g"
	"gate/msg"
)

var dispatcher = net4g.NewDispatcher("client", 1)
var serializer = msg.NewGateSerializer()

func connect(onCreate func(agent net4g.NetAgent)) {

	msg.InitSerializer(serializer)

	net4g.NetConfig.BeginBytes = []byte{0x10, 0x02}
	net4g.NetConfig.EndBytes = []byte{0x10, 0x03}
	net4g.NetConfig.LengthSize = 1
	net4g.NetConfig.IdSize = 1

	dispatcher.OnConnectionCreated(onCreate)

	addr := ":8083"
	//addr := "sgu.youstars.com.cn:8083"
	net4g.NewTcpClient(net4g.NewNetKeyAddrFn("gate_client", addr)).
		SetSerializer(serializer).
		AddDispatchers(dispatcher).
		DisableAutoReconnect().
	//EnableHeartbeat().
		Start().Wait()

}
