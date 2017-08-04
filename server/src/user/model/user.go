package model

type User struct {
	Id                int              `json:"id"`
	Channel           *string          `json:"channel"` // 渠道
	Uid               *string          `json:"uid"`    // 渠道用户id

	Nick              *string          `json:"nick"`              // 昵称
	Signature         *string          `json:"signature"`         // 个性签名
	Avatar            *string          `json:"avatar"`            // 个人头像
	Sex               byte             `json:"sex"`               // 性别：1-男；2-女；3-未知

	Point             int              `json:"point"`             // 积分
	LastLat           float32          `json:"last_lat"`          // 最近一次定位的维度
	LastLong          float32          `json:"last_long"`         // 最近一次定位的经度
	LastHeartbeatTime *model.TimeStamp `json:"last_locate_time"`  // 最近一次心跳时间
	CreateTime        *model.TimeStamp `json:"create_time,omitempty"`
	UpdateTime        *model.TimeStamp `json:"update_time,omitempty"`
}
