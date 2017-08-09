package errcode

import (
	"common/httpx"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"common/dbx"
)

func NewResponse(code int, tag ...string) *httpx.Response {
	return NewResponseWithMsg(code, GetMsg(code), tag...)
}

func NewResponseWithMsg(code int, msg string, tag ...string) *httpx.Response {
	//if code != 0 {
	//	if log4g.IsDebugEnabled() {
	//		log4g.Error("\n" + string(debug.Stack()))
	//	}
	//}
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
	code := CODE_COMMON_ERROR
	msg := err.Error()
	if err == dbx.ErrNotFound {
		code = CODE_COMMON_NOT_FOUND
		msg = GetMsg(CODE_COMMON_NOT_FOUND)
	}
	render.WriteJSON(rw, NewResponseWithMsg(code, msg))
}

func WriteSuccessOrErrorResponse(rw gin.ResponseWriter, err error) {
	if err != nil {
		WriteErrorResponse(rw, err)
	} else {
		WriteSuccessResponse(rw)
	}
}

func WriteEmptyArgResponse(rw gin.ResponseWriter, input ...string) {
	render.WriteJSON(rw, NewResponse(CODE_COMMON_EMPTY_ARG, input...))
}
//
//func NewEmptyArgResponse(input ...string) *httpx.Response {
//	return NewResponse(CODE_COMMON_EMPTY_ARG, input...)
//}

func WriteWrongArgResponse(rw gin.ResponseWriter, input ...string) {
	render.WriteJSON(rw, NewResponse(CODE_COMMON_WRONG_ARG, input...))
}

//func NewWrongArgResponse(input ...string) *httpx.Response {
//	return NewResponse(CODE_COMMON_WRONG_ARG, input...)
//}


func WriteIllegalArgResponse(rw gin.ResponseWriter, input ...string) {
	render.WriteJSON(rw, NewResponse(CODE_COMMON_ILLEGAL_ARG, input...))
}

//func NewIllegalArgResponse(input ...string) *httpx.Response {
//	return NewResponse(CODE_COMMON_ILLEGAL_ARG, input...)
//}

func WriteDataResponse(rw gin.ResponseWriter, data interface{}) {
	res := NewResponse(CODE_COMMON_SUCCESS)
	res.Data = data
	render.WriteJSON(rw, res)
}