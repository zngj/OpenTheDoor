package model

import (
	"common/sqlx"
	"time"
)

type User struct {
	Id          string          `db:"id"`           // 用户id
	Channel     string          `db:"channel"`      // 渠道
	OpenId      sqlx.NullString `db:"open_id"`      // 渠道用户id
	NickName    sqlx.NullString `db:"nick_name"`    // 昵称
	PhoneNumber sqlx.NullString `db:"phone_number"` // 手机
	Email       sqlx.NullString `db:"email"`        // 邮件
	Password    sqlx.NullString `db:"password"`
	Signature   sqlx.NullString `db:"signature"` // 个性签名
	Avatar      sqlx.NullString `db:"avatar"`    // 个人头像
	Sex         sqlx.NullInt64  `db:"sex"`       // 性别：1-男；2-女；3-未知
	InsertTime  *time.Time      `db:"insert_time"`
	UpdateTime  *time.Time      `db:"update_time"`
}
