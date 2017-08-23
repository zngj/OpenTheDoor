package tokenutil

import (
	"common/sg"
	"common/httpx"
	"common/errcode"
	"common/redisx"
	"github.com/gin-gonic/gin"
	"github.com/carsonsx/log4g"
)


const HEADER_ACCESS_TOKEN = "Access-Token"

func GetAccessTokenKey(token string) string {
	return "access_token:" + token
}

func VerifyToken(c *gin.Context) {

	log4g.Debug("request uri: %s", c.Request.RequestURI)

	sgc := sg.Context(c)
	res := new (sg.Response)
	accessToken := c.Request.Header.Get(HEADER_ACCESS_TOKEN)
	if sgc.CheckParamEmpty(accessToken) {
		c.Abort()
		return
	}
	err := Get("http://localhost:8081/check_token", accessToken, res)
	if err == nil && res.Code > 0 {
		err = errcode.NewError2(res.Code, res.Msg)
	}
	if err != nil {
		sgc.WriteError(err)
		c.Abort()
	}
}

func GetUserId(c *gin.Context) (userId string, err error) {
	accessToken := c.Request.Header.Get(HEADER_ACCESS_TOKEN)
	key := GetAccessTokenKey(accessToken)
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
	log4g.Debug("get user %s by key %s from redis", userId, key)
	return
}

func Get(url, accessToken string, v interface{}) error {
	header := make(map[string]string)
	header[HEADER_ACCESS_TOKEN] = accessToken
	return httpx.Get(url, header, v)
}