package controller

import (
	"github.com/gin-gonic/gin"
	"common/sg"
	"common/vo"
	"github.com/carsonsx/log4g"
	"fmt"
	"common/httpx"
	"common/util"
	"usercenter/service"
	"common/errcode"
)

const (
	WXAPI_ACCESS_TOKEN = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	WXAPI_REFRESH_TOKEN = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	WXAPI_APPID        = "wx7b727b2cccb9c916"
	WXAPI_SECRET       = "9166cb4364d7311b09f5086a8ad84d10"
)

func WxapiLogin(c *gin.Context) {
	sgc := sg.Context(c)
	var login vo.WxappLogin
	if sgc.CheckError(c.Bind(&login)) {
		return
	}
	if sgc.CheckParamEmpty(login.Code, "code") {
		return
	}
	session, err := wx_api_sns_oauth2_access_token(login.Code)
	if sgc.CheckError(err) {
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

func WxapiVerifyToken()  {

	//检查access_token是否失效


	//根据access_token获取微信access_token

	//检查access_token是否失效

	//尝试刷样

}

func wx_api_sns_oauth2_access_token(code string) (session *vo.WxSession, err error) {
	log4g.Debug("login wx api with code %s", code)
	url := fmt.Sprintf(WXAPI_ACCESS_TOKEN, WXAPI_APPID, WXAPI_SECRET, code)
	session = new(vo.WxSession)
	session.Client = "app"
	err = httpx.Get(url, nil, session)
	if session.Errcode > 0 {
		err = errcode.NewError(session.Errcode)
		log4g.Error(err)
		//for testing
		//err = nil
		//session.AccessToken = "ot96OBqsvSa3WLFBz4U+gw=="
		//session.ExpiresIn = 7200
		//session.RefreshToken = "ot96OBqsvSa3WLFBz4U+gw=="
		//session.Openid = "oVecO0Ze4kNxMGymF05d1uiIcmqA"
		//session.Scope = ""
	}
	log4g.Debug(log4g.JsonFunc(session))
	return

}

