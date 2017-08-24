package vo

import "time"

type UserVo struct {
	Id          string     `json:"id"`           // 用户id
	NickName    string     `json:"nick_name"`    // 昵称
	PhoneNumber string     `json:"phone_number"` // 手机
	Email       string     `json:"email"`        // 邮件
	Signature   string     `json:"signature"`    // 个性签名
	Avatar      string     `json:"avatar"`       // 个人头像
	Sex         int8        `json:"sex"`          // 性别：1-男；2-女；0-未知
	InsertTime  *time.Time `json:"insert_time"`
	UpdateTime  *time.Time `json:"update_time,omitempty"`
}
