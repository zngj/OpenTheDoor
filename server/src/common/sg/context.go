package sg

import (
	"common/errcode"
	"github.com/gin-gonic/gin"
	"common/dbx"
	"github.com/gin-gonic/gin/render"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
	Tag  string      `json:"tag,omitempty"`
}

func NewResponse(code int, tag ...string) *Response {
	return NewResponseWithMsg(code, errcode.GetMsg(code), tag...)
}

func NewResponseWithMsg(code int, msg string, tag ...string) *Response {
	//if code != 0 {
	//	if log4g.IsDebugEnabled() {
	//		log4g.Error("\n" + string(debug.Stack()))
	//	}
	//}
	res := &Response{Code: code, Msg: msg}
	if len(tag) > 0 {
		res.Tag = tag[0]
	}
	return res
}


func Context(c *gin.Context) *context {
	w := new(context)
	w.gc = c
	return w
}

type context struct {
	gc *gin.Context
}

func (c *context) CheckParamEmpty(arg string, tag ...string) bool {
	if arg == "" {
		c.WriteParamEmpty(tag...)
		return true
	}
	return false
}

func (c *context) WriteParamEmpty(tag ...string) {
	c.Write(errcode.CODE_COMMON_EMPTY_ARG, tag...)
}

func (c *context) CheckParamEqual(paramValue, expectValue string, tag ...string) bool {
	if paramValue != expectValue {
		c.WriteParamWrong(tag...)
		return true
	}
	return false
}

func (c *context) CheckParamCorrect(correct bool, tag ...string) bool {
	if correct {
		return false
	}
	c.WriteParamWrong(tag...)
	return true
}

func (c *context) WriteParamWrong(tag ...string) {
	c.Write(errcode.CODE_COMMON_WRONG_ARG, tag...)
}

func (c *context) CheckError(err error) bool {
	if err == nil {
		return false
	}
	c.WriteError(err)
	return true
}

func (c *context) WriteSuccess() {
	c.Write(errcode.CODE_COMMON_SUCCESS)
}

func (c *context) WriteError(err error) {
	code := errcode.CODE_COMMON_ERROR
	msg := err.Error()
	if err == dbx.ErrNotFound {
		code = errcode.CODE_COMMON_NOT_FOUND
		msg = errcode.GetMsg(errcode.CODE_COMMON_NOT_FOUND)
	} else if sgerr, ok := err.(*errcode.SGError); ok {
		code = sgerr.Code()
		msg = sgerr.Error()
	}
	c.WriteWithMsg(code, msg)
}

func (c *context) WriteData(data interface{}) {
	res := NewResponse(errcode.CODE_COMMON_SUCCESS)
	res.Data = data
	c.WriteResponse(res)
}
func (c *context) WriteSuccessOrError(err error) {
	if err != nil {
		c.WriteError(err)
	} else {
		c.WriteSuccess()
	}
}

func (c *context) WriteDataOrError(data interface{}, err error) {
	if err != nil {
		c.WriteError(err)
	} else {
		c.WriteData(data)
	}
}

func (c *context) Write(code int, tag ...string) {
	c.WriteResponse(NewResponse(code, tag...))
}

func (c *context) WriteWithMsg(code int, msg string, tag ...string) {
	c.WriteResponse(NewResponseWithMsg(code, msg, tag...))
}

func (c *context) WriteResponse(res *Response) {
	render.WriteJSON(c.gc.Writer, res)
}

