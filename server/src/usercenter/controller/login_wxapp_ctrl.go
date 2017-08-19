package controller

import (
	"common/sg"
	"common/util"
	"common/vo"
	"fmt"
	"github.com/carsonsx/log4g"
	"github.com/gin-gonic/gin"
	"usercenter/service"
	"common/httpx"
	"common/errcode"
)

const (
	WEAPP_URL_FORMAT = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	WEAPP_APPID      = "wx7b727b2cccb9c916"
	WEAPP_SECRET     = "9166cb4364d7311b09f5086a8ad84d10"
)

func WeappLogin(c *gin.Context) {
	sgc := sg.Context(c)
	var login vo.WxappLogin
	if sgc.CheckError(c.Bind(&login)) {
		return
	}
	if sgc.CheckParamEmpty(login.Code, "code") {
		return
	}
	session, err := code2session(login.Code)
	if sgc.CheckError(err) {
		return
	}
	if sgc.CheckParamEmpty(session.Openid, "code") || sgc.CheckParamEmpty(session.SessionKey, "code") {
		return
	}

	accessToken := util.NewUuid()
	if sgc.CheckError(service.SaveWxappLoginSession(accessToken, session)) {
		return
	}

	result := new(vo.LoginToken)
	result.AccessToken = accessToken
	result.ExpiresIn = session.ExpiresIn
	sgc.WriteData(result)
}

func code2session(code string) (session *vo.WxSession, err error) {
	log4g.Debug("login weapp with code %s", code)
	url := fmt.Sprintf(WEAPP_URL_FORMAT, WEAPP_APPID, WEAPP_SECRET, code)
	session = new(vo.WxSession)
	session.Client = "weapp"
	err = httpx.Get(url, nil, session)
	if session.Errcode > 0 {
		err = errcode.NewError(session.Errcode)
		log4g.Error(err)
		//for testing
		//err = nil
		//session.SessionKey = "ot96OBqsvSa3WLFBz4U+gw=="
		//session.ExpiresIn = 7200
		//session.Openid = "oVecO0Ze4kNxMGymF05d1uiIcmqA"
	}
	log4g.Debug(log4g.JsonFunc(session))
	return
}
