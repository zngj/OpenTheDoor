package cmnmsg

import "common/httpx"

func NewResponse(code int, input ...string) *httpx.Response {
	return NewResponseWithMsg(code, GetMsg(code), input...)
}

func NewResponseWithMsg(code int, msg string, input ...string) *httpx.Response {
	res := &httpx.Response{Code: code, Msg: msg}
	if len(input) < 0 {
		res.Input = input[0]
	}
	return res
}

func NewSuccessResponse() *httpx.Response {
	return NewResponse(CODE_COMMON_SUCCESS)
}

func NewErrorResponse(err error) *httpx.Response {
	return NewResponseWithMsg(CODE_COMMON_ERROR, err.Error())
}

func NewEmptyArgResponse(input ...string) *httpx.Response {
	return NewResponse(CODE_COMMON_EMPTY_ARG, input...)
}

func NewWrongArgResponse(input ...string) *httpx.Response {
	return NewResponse(CODE_COMMON_WRONG_ARG, input...)
}

func NewIllegalArgResponse(input ...string) *httpx.Response {
	return NewResponse(CODE_COMMON_ILLEGAL_ARG, input...)
}

func NewDataResponse(data interface{}) *httpx.Response {
	res := NewSuccessResponse()
	res.Data = data
	return res
}