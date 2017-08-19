package controller

import (
	"common/sg"
	"github.com/gin-gonic/gin"
	"etransin/dao"
	"strconv"
	"common/tokenutil"
	"common/model"
	"common/vo"
)

func ExploreNotification(c *gin.Context) {
	sgc := sg.Context(c)
	userId, err := tokenutil.GetUserId(c)
	if sgc.CheckError(err) {
		return
	}
	var notification model.Notification
	if sgc.CheckError(dao.NewNotificationDao().Explore(userId, &notification)) {
		return
	}
	var notificationVo vo.NotificationVo
	notificationVo.Id = notification.Id
	notificationVo.Type = notification.Type
	sgc.WriteData(&notificationVo)
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
	sgc.WriteSuccessOrError(dao.NewNotificationDao().Consume(id))
}
