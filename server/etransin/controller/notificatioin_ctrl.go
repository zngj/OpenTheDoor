package controller

import (
	"common/sg"
	"github.com/gin-gonic/gin"
	"etransin/dao"
	"strconv"
	"common/tokenutil"
	"common/model"
	"common/vo"
	"common/sqlx"
)

func OneNotification(c *gin.Context) {
	sgc := sg.Context(c)
	userId, err := tokenutil.GetUserId(c)
	if sgc.CheckError(err) {
		return
	}
	var notification model.Notification
	if sgc.CheckErrorIgnore(dao.NewNotificationDao().One(userId, &notification), sqlx.ErrNotFound) {
		return
	}
	if notification.Id > 0 {
		var notificationVo vo.NotificationVo
		notificationVo.Id = notification.Id
		notificationVo.Type = notification.Type
		sgc.WriteData(&notificationVo)
	} else {
		sgc.WriteSuccess()
	}
}

func ConsumeRouterNotification(c *gin.Context) {
	sgc := sg.Context(c)
	param := "id"
	strId := c.PostForm(param)
	if sgc.CheckParamEmpty(strId, param) {
		return
	}
	id, err := strconv.ParseUint(strId, 10, 64)
	if sgc.CheckError(err) {
		return
	}
	sgc.WriteSuccessOrError(dao.NewNotificationDao().Consume(id))
}
