package controller

import (
	"common/sg"
	"github.com/gin-gonic/gin"
	"etransin/service"
	"time"
	"common/tokenutil"
	"common/vo"
)

type TestRouterInVo struct {
	UserId string `json:"user_id"`
	GateId string `json:"gate_id"`
}

func TestRouterIn(c *gin.Context) {
	sgc := sg.Context(c)
	userId, _ := tokenutil.GetUserId(c)
	var testRouterVo vo.TestRouterVo
	c.BindJSON(&testRouterVo)
	if userId == "" {
		userId = testRouterVo.UserId
	}
	var key string
	var err error
	if userId != "" {
		_, key, err = service.CreateEvidenceWithEncrypt(userId, service.GATE_DIRECTION_IN)
		if sgc.CheckError(err) {
			return
		}
	} else {
		key = testRouterVo.EvidenceKey
	}
	if sgc.CheckParamCorrect(key != "") {
		return
	}
	gateId := "010100101"
	t := testRouterVo.ScanTime
	if t == 0 {
		t = time.Now().Unix()
	}

	code, err := service.VerifyEvidenceKey(key, gateId)
	if err != nil {
		sgc.WriteError(err)
		return
	}
	if code > 0 {
		sgc.Write(code)
		return
	}
	sgc.WriteSuccessOrError(service.SubmitEvidenceKey(key, t, gateId))
}

func TestRouterOut(c *gin.Context) {
	sgc := sg.Context(c)
	userId, _ := tokenutil.GetUserId(c)
	var testRouterVo vo.TestRouterVo
	c.BindJSON(&testRouterVo)
	if userId == "" {
		userId = testRouterVo.UserId
	}
	var key string
	var err error
	if userId != "" {
		_, key, err = service.CreateEvidenceWithEncrypt(userId, service.GATE_DIRECTION_OUT)
		if sgc.CheckError(err) {
			return
		}
	} else {
		key = testRouterVo.EvidenceKey
	}
	if sgc.CheckParamCorrect(key != "") {
		return
	}
	gateId := "010100202"
	t := testRouterVo.ScanTime
	if t == 0 {
		t = time.Now().Unix()
	}

	code, err := service.VerifyEvidenceKey(key, gateId)
	if err != nil {
		sgc.WriteError(err)
		return
	}
	if code > 0 {
		sgc.Write(code)
		return
	}

	sgc.WriteSuccessOrError(service.SubmitEvidenceKey(key, time.Now().Unix(), gateId))
}