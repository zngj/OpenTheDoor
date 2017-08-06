package token

import (
	"github.com/carsonsx/log4g"
	"common/redisx"
	"usercenter/util"
	"time"
)

func IsValid(token string) (bool, error) {
	key := util.GetAccessTokenKey(token)
	exist, err := redisx.Client.HExists(key, "userid").Result()
	if err != nil {
		log4g.Error(err)
		return false, err
	}
	return exist, nil
}

func Save(userId, accessToken, sessionKey, unionid string, now time.Time, expiresAt time.Time) error {
	fields := make(map[string]interface{})
	fields["userid"] = userId
	fields["session_key"] = sessionKey
	if unionid != "" {
		fields["unionid"] = unionid
	}
	fields["login_time"] = now.UnixNano()
	key := util.GetAccessTokenKey(accessToken)
	err := redisx.Client.HMSet(key, fields).Err()
	if err != nil {
		log4g.Error(err)
		return err
	}
	err = redisx.Client.ExpireAt(key, expiresAt).Err()
	if err != nil {
		log4g.Error(err)
	}
	return err
}

func GetUserId(accessToken string) (userId string, err error) {
	key := util.GetAccessTokenKey(accessToken)
	cmd := redisx.Client.HGet(key, "userid")
	err = cmd.Err()
	if err != nil {
		log4g.Error(err)
		return
	}
	userId, err = cmd.Result()
	if err != nil {
		log4g.Error(err)
	}
	log4g.Debug("get userId %s by key %s from redis", userId, key)
	return
}