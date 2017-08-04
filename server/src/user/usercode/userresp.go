package usercode

import (
	"common/httpx"
	"common/cmnmsg"
)

func NewUserTokenExpiredResponse() *httpx.Response {
	return cmnmsg.NewResponse(CODE_USER_TOKEN_EXPIRED)
}

