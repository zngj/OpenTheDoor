package msg

import (
	"github.com/carsonsx/net4g"
	"reflect"
)

const (
	MSG_SET_GATE_NO         = 100
	MSG_SET_GATE_NO_SUCCESS = 200

	MSG_USER_IN         = 101
	MSG_USER_IN_SUCCESS = 201

	MSG_USER_OUT         = 102
	MSG_USER_OUT_SUCCESS = 202
)

var Dispatcher = net4g.NewDispatcher("gate", 1)

var Serializer = net4g.NewJsonSerializer()

func init() {

	net4g.NetConfig.MessageLengthSize = 1
	net4g.NetConfig.MessageIdSize = 1

	Serializer.RegisterId(SetGateNoType, true, MSG_SET_GATE_NO)
	Serializer.RegisterId(SetGateNoSuccessType, true, MSG_SET_GATE_NO_SUCCESS)
	Serializer.RegisterId(UserInType, true, MSG_USER_IN)
	Serializer.RegisterId(UserInSuccessType, true, MSG_USER_IN_SUCCESS)
	Serializer.RegisterId(UserOutType, true, MSG_USER_OUT)
	Serializer.RegisterId(UserOutSuccessType, true, MSG_USER_OUT_SUCCESS)

}

var SetGateNoType = reflect.TypeOf(&SetGateNo{})
var SetGateNoSuccessType = reflect.TypeOf(&SetGateNoSuccess{})
var UserInType = reflect.TypeOf(&UserIn{})
var UserInSuccessType = reflect.TypeOf(&UserInSuccess{})
var UserOutType = reflect.TypeOf(&UserOut{})
var UserOutSuccessType = reflect.TypeOf(&UserOutSuccess{})

type SetGateNo struct {
	No string `json:"no"`
}

type SetGateNoSuccess struct {
}

type UserIn struct {
	QrCode string `json:"qr"`
}

type UserInSuccess struct {
}

type UserOut struct {
	QrCode string `json:"qr"`
}

type UserOutSuccess struct {
}
