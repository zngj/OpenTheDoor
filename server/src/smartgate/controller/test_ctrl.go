package controller

import (
	"github.com/gin-gonic/gin"
	"common/sg"
	"smartgate/service"
	"time"
)

type TestRouterInVo struct {
	UserId string `json:"user_id"`
	GateId string `json:"gate_id"`
}

func TestRouterIn(c *gin.Context) {
	sgc := sg.Context(c)
	param := "userid"
	userId := c.Param(param)
	if sgc.CheckParamEmpty(userId, param) {
		return
	}
	_, key, err := service.CreateEvidenceWithEncrypt(userId, service.GATE_DIRECTION_IN)
	if sgc.CheckError(err) {
		return
	}
	sgc.WriteSuccessOrError(service.SubmitEvidenceKey(key, time.Now().Unix(), "010100101"))
}

func TestRouterOut(c *gin.Context) {
	sgc := sg.Context(c)
	param := "userid"
	userId := c.Param(param)
	if sgc.CheckParamEmpty(userId, param) {
		return
	}
	_, key, err := service.CreateEvidenceWithEncrypt(userId, service.GATE_DIRECTION_OUT)
	if sgc.CheckError(err) {
		return
	}
	sgc.WriteSuccessOrError(service.SubmitEvidenceKey(key, time.Now().Unix(), "010100202"))
}
