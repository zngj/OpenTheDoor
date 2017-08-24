package service

import (
	"common/model"
	"etransin/dao"
	"common/httpx"
	"common/sg"
	"fmt"
)
func CreateNotification(userId string, notificationType int8) error {
	notification := new(model.Notification)
	notification.UserId = userId
	notification.Type = notificationType
	err := dao.NewNotificationDao().Insert(notification)
	if err != nil {
		return err
	}
	var res sg.Response
	return httpx.Get(fmt.Sprintf("http://localhost:8084/notify?user_id=%s&id=%d&type=%d", userId, notification.Id, notificationType), nil, &res)
}
