package test_test

import (
	"github.com/carsonsx/net4g"
	"gate/msg"
	"github.com/carsonsx/log4g"
)

var dispatcher = net4g.NewDispatcher("client")
var serializer = msg.NewGateSerializer()
var agent net4g.NetAgent

func onInit()  {

	serializer.DeserializeId(new(msg.S2CGateLogin), msg.GATE_LOGIN)
	serializer.DeserializeId(new(msg.S2CRsaKey), msg.RSA_KEY)
	serializer.SerializeId(new(msg.C2SVerifyEvidence), msg.VERIFY_EVIDENCE)
	serializer.DeserializeId(new(msg.S2CVerifyEvidence), msg.VERIFY_EVIDENCE)
	serializer.SerializeId(new(msg.C2SSubmitEvidence), msg.SUBMIT_EVIDENCE)
	serializer.DeserializeId(new(msg.S2CSubmitEvidence), msg.SUBMIT_EVIDENCE)

}

func notLoginResult(agent net4g.NetAgent) {
	log4g.Debug("[client]gate not login")
	net4g.TestDone()
}

func connect(callback func()) {

	log4g.SetLevel(log4g.LEVEL_TRACE)

	net4g.NetConfig.ReadMode = net4g.READ_MODE_BEGIN_END
	net4g.NetConfig.BeginBytes = []byte{0x10, 0x02}
	net4g.NetConfig.EndBytes = []byte{0x10, 0x03}
	net4g.NetConfig.LengthSize = 1
	net4g.NetConfig.IdSize = 1
	net4g.NetConfig.Print()

	onInit()

	dispatcher.AddHandler(notLoginResult, msg.NOT_LOGIN)
	dispatcher.OnConnectionCreated(func(a net4g.NetAgent) {
		agent = a
		callback()
	})
	addr := ":8083"
	//addr = "sgu.youstars.com.cn:8083"
	net4g.NewTcpClient(net4g.NewNetAddrFn(addr)).
		SetSerializer(serializer).
		AddDispatchers(dispatcher).
		DisableAutoReconnect().
	//EnableHeartbeat().
		Connect()

	net4g.TestWait()

	log4g.Debug("?????????????????????????????")

	agent.Close()

}

func start(calls ...func())  {
	connect(func() {
		net4g.TestCall(calls...)
	})
}
