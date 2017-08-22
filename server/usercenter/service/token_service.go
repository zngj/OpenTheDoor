package service

import (
	"common/redisx"
	"github.com/carsonsx/log4g"
	"time"
	"common/vo"
	"common/tokenutil"
)

func IsValid(_token string) (bool, error) {
	key := tokenutil.GetAccessTokenKey(_token)
	exist, err := redisx.Client.HExists(key, "userid").Result()
	if err != nil {
		log4g.Error(err)
		return false, err
	}
	return exist, nil
}

func SaveWxToken(accessToken, userId string, session *vo.WxSession, now time.Time) error {
	fields := make(map[string]interface{})
	if session.AccessToken != "" {
		fields["access_token"] = session.AccessToken
	}
	if session.RefreshToken != "" {
		fields["refresh_token"] = session.RefreshToken
	}
	if session.SessionKey != "" {
		fields["session_key"] = session.SessionKey
	}
	if session.Unionid != "" {
		fields["unionid"] = session.Unionid
	}
	if session.Openid != "" {
		fields["openid"] = session.Openid
	}
	if session.Scope != "" {
		fields["scope"] = session.Scope
	}
	return SaveToken(accessToken, userId, session.Client, now, session.ExpiresIn, fields)
}


func SaveToken(accessToken, userId, client string, now time.Time, expiresIn int, extra ...map[string]interface{}) error {
	fields := make(map[string]interface{})
	fields["userid"] = userId
	fields["client"] = client
	fields["login_time"] = now.UnixNano()
	if len(extra) > 0 && extra[0] != nil && len(extra[0]) > 0 {
		for k, v := range extra[0] {
			fields[k] = v
		}
	}
	key := tokenutil.GetAccessTokenKey(accessToken)
	log4g.Debug("[hmset] key:%s, fields:%v", key, fields)
	err := redisx.Client.HMSet(key, fields).Err()
	if err != nil {
		log4g.Error(err)
		return err
	}
	expiresAt := now.Add(time.Duration(expiresIn)*time.Second)
	err = redisx.Client.ExpireAt(key, expiresAt).Err()
	if err != nil {
		log4g.Error(err)
	}
	return err
}