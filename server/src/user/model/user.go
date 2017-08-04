package model

import "time"

type User struct {
	Id         int        `json:"id"`        // 用户id
	Channel    string     `json:"channel"`   // 渠道
	OpenId     string     `json:"openId"`    // 渠道用户id
	Nick       string     `json:"nick"`      // 昵称
	Mobile     string     `json:"mobile"`    // 昵称
	Email      string     `json:"emial"`     // 昵称
	Signature  string     `json:"signature"` // 个性签名
	Avatar     string     `json:"avatar"`    // 个人头像
	Sex        byte       `json:"sex"`       // 性别：1-男；2-女；3-未知
	CreateTime *time.Time `json:"create_time,omitempty"`
	UpdateTime *time.Time `json:"update_time,omitempty"`
}
