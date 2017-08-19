package service

import (
	"common/model"
	"etransin/dao"
)
func CreateNotification(userId string, notificationType int8) error {
	notification := new(model.Notification)
	notification.UserId = userId
	notification.Type = notificationType
	return dao.NewNotificationDao().Insert(notification)
}
