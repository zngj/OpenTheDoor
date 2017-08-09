package sg

import (
	"common/errcode"
	"github.com/gin-gonic/gin"
)

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
		errcode.WriteEmptyArgResponse(c.gc.Writer, tag...)
		return true
	}
	return false
}

func (c *context) WriteParamEmpty(tag ...string) {
	errcode.WriteEmptyArgResponse(c.gc.Writer, tag...)
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
	errcode.WriteWrongArgResponse(c.gc.Writer, tag...)
}

func (c *context) CheckError(err error) bool {
	if err == nil {
		return false
	}
	c.WriteError(err)
	return true
}

func (c *context) WriteSuccess() {
	errcode.WriteSuccessResponse(c.gc.Writer)
}

func (c *context) WriteError(err error) {
	errcode.WriteErrorResponse(c.gc.Writer, err)
}

func (c *context) WriteData(data interface{}) {
	errcode.WriteDataResponse(c.gc.Writer, data)
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
