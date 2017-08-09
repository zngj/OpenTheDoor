package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"smartgate/dao"
	"smartgate/service"
	"common/sg"
)

func GetRouterNotification(c *gin.Context) {
	sgc := sg.Context(c)
	param := "category"
	category := c.Param(param)
	if sgc.CheckParamEmpty(category) {
		return
	}
	if sgc.CheckParamEqual(category, service.NOTFICATION_ROUTER, param) {
		return
	}
	userId, err := GetUserId(c.Request.Header)
	if sgc.CheckError(err) {
		return
	}
	sgc.WriteDataOrError(service.GetRouterNotification(userId))
}

func ConsumeRouterNotification(c *gin.Context) {
	sgc := sg.Context(c)
	param := "id"
	strId := c.Param(param)
	if sgc.CheckParamEmpty(strId, param) {
		return
	}
	id, err := strconv.ParseUint(strId, 10, 64)
	if sgc.CheckError(err) {
		return
	}
	sgc.WriteSuccessOrError(dao.ConsumeNotification(id))
}