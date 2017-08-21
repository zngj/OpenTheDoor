package dao

import (
	"common/model"
	"common/sqlx"
	"time"
)

func NewNotificationDao(dao ...*sqlx.Dao) *notificationDao {
	d := new(notificationDao)
	if len(dao) > 0 {
		d.dao = dao[0]
	} else {
		d.dao = new(sqlx.Dao)
	}
	return d
}

type notificationDao struct {
	dao *sqlx.Dao
}

func (d *notificationDao) Insert(notification *model.Notification) error {
	sql := "insert sg_sys_notification (user_id,type,insert_time) values (?,?,?)"
	return sqlx.NewDao().Exec(sql, notification.UserId, notification.Type, time.Now())
}

func (d *notificationDao) Current(userId string, n *model.Notification) error {
	sql := "select * from sg_sys_notification where received=0 and user_id=? order by id asc"
	return d.dao.Query(sql, userId).One(n)
}

func (d *notificationDao) Consume(id uint64) error {
	return d.dao.Exec("update sg_sys_notification set received=1 where id=?", id)
}
