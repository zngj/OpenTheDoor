package errcode

import (
	"common/httpx"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func NewResponse(code int, tag ...string) *httpx.Response {
	return NewResponseWithMsg(code, GetMsg(code), tag...)
}

func NewResponseWithMsg(code int, msg string, tag ...string) *httpx.Response {
	res := &httpx.Response{Code: code, Msg: msg}
	if len(tag) > 0 {
		res.Tag = tag[0]
	}
	return res
}

func WriteResponse(rw gin.ResponseWriter, code int, tag ...string) {
	render.WriteJSON(rw, NewResponse(code, tag...))
}

func WriteResponseWithMsg(rw gin.ResponseWriter, code int, msg string, tag ...string) {
	render.WriteJSON(rw, NewResponseWithMsg(code, msg, tag...))
}

func WriteSuccessResponse(rw gin.ResponseWriter) {
	render.WriteJSON(rw, NewResponse(CODE_COMMON_SUCCESS))
}

func WriteErrorResponse(rw gin.ResponseWriter, err error) {
	render.WriteJSON(rw, NewResponseWithMsg(CODE_COMMON_ERROR, err.Error()))
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

func WriteDataResponse(rw gin.ResponseWriter, data interface{}) {
	res := NewResponse(CODE_COMMON_SUCCESS)
	res.Data = data
	render.WriteJSON(rw, res)
}