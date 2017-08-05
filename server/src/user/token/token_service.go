package token

import (
	"github.com/carsonsx/log4g"
	"common/redisx"
	"user/util"
)

func IsValidToken(token string) (bool, error) {
	key := util.GetTokenKey(token)
	exist, err := redisx.Client.HExists(key, "userid").Result()
	if err != nil {
		log4g.Error(err)
		return false, err
	}
	return exist, nil
}