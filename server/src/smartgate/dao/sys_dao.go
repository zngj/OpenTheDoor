package dao

import (
	"common/model"
	"common/dbx"
	"time"
)

func InsertNotification(notification *model.Notification) error {
	sql := "insert sg_sys_notification (user_id,category,content_id,insert_time) values (?,?,?,?)"
	return dbx.NewDao().Exec(sql, notification.UserId, notification.Category, notification.ContentId, time.Now())
}
