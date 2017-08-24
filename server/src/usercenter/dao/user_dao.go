package dao

import (
	"common/sqlx"
	"common/model"
	"time"
)

func NewUserDao(sqlxDao ...*sqlx.Dao) *userDao {
	d := new(userDao)
	if len(sqlxDao) > 0 {
		d.sqlxDao = sqlxDao[0]
	} else {
		d.sqlxDao = new(sqlx.Dao)
	}
	return d
}

type userDao struct {
	sqlxDao *sqlx.Dao
}

func (d *userDao) IsPhoneNumberExist(phoneNumber string) (exist bool, err error) {
	sql := "select id from uc_user_info where phone_number=?"
	return d.sqlxDao.Query(sql, phoneNumber).Exist()
}

func (d *userDao) Insert(userId, phoneNumber, password string) error {
	sql := "insert into uc_user_info (id,channel,phone_number,password,insert_time) values (?,?,?,?,?)"
	return d.sqlxDao.Exec(sql, userId, "self", phoneNumber, password, time.Now())
}

func (d *userDao) GetByPhoneNumber(phoneNumber string, user *model.User) error {
	sql := "select * from uc_user_info where phone_number=?"
	return d.sqlxDao.Query(sql, phoneNumber).One(user)
}

func (d *userDao) GetById(id string, user *model.User) error {
	sql := "select * from uc_user_info where id=?"
	return d.sqlxDao.Query(sql, id).One(user)
}

