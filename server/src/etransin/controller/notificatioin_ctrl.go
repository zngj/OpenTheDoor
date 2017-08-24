package controller

import (
	"common/sg"
	"github.com/gin-gonic/gin"
	"etransin/dao"
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
	var nvo vo.NotificationVo
	if sgc.CheckError(c.BindJSON(&nvo)) {
		return
	}
	if sgc.CheckParamCorrect(nvo.Id > 0, "id") {
		return
	}
	sgc.WriteSuccessOrError(dao.NewNotificationDao().Consume(nvo.Id))
}
