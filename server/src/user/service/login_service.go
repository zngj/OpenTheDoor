package service

import (
	"user/vo"
	"common/redisx"
	"time"
)

func SaveLoginSession(token string, wxsession *vo.WxappSession) error {
	fields := make(map[string]interface{})
	fields["openid"] = wxsession.Openid
	fields["session_key"] = wxsession.Session_key
	fields["unionid"] = wxsession.Unionid
	fields["create_time"] = time.Now()
	return redisx.Client.HMSet(token, fields).Err()
}
