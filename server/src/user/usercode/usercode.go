package usercode

import "common/cmnmsg"

// user code between -1000 and -1999

const (
	CODE_USER_TOKEN_EXPIRED = 1000 // 登录失效
)

var msg_map = map[int]string {
	CODE_USER_TOKEN_EXPIRED: "token was expired",
}

func init()  {
	cmnmsg.AddMsg(msg_map)
}