package dao

import (
	"common/model"
	"common/sqlx"
	"time"
	"github.com/carsonsx/log4g"
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
	id, err := sqlx.NewDao().Insert(sql, notification.UserId, notification.Type, time.Now())
	if err != nil {
		log4g.Error(err)
		return err
	}
	notification.Id = uint64(id)
	return nil
}

func (d *notificationDao) One(userId string, n *model.Notification) error {
	sql := "select * from sg_sys_notification where received=0 and user_id=? order by id asc"
	return d.dao.Query(sql, userId).One(n)
}

func (d *notificationDao) Consume(id uint64) error {
	return d.dao.Exec("update sg_sys_notification set received=1 where id=?", id)
}
