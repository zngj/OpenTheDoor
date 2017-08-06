package controller

import (
	"common/cmnmsg"
	"encoding/json"
	"fmt"
	"github.com/carsonsx/log4g"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"io/ioutil"
	"net/http"
	"usercenter/service"
	"usercenter/vo"
	"errors"
	"common/util"
)

const (
	WEAPP_URL_FORMAT = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	WEAPP_APPID      = "wx7b727b2cccb9c916"
	WEAPP_SECRET     = "9166cb4364d7311b09f5086a8ad84d10"
)

func WxappLogin(c *gin.Context) {

	var login vo.WxappLogin
	err := c.Bind(&login)
	if err != nil {
		cmnmsg.WriteErrorResponse(c.Writer, err)
		return
	}
	if login.Code == "" {
		render.WriteJSON(c.Writer, cmnmsg.NewEmptyArgResponse("code"))
		return
	}

	session, err := code2session(login.Code)
	if err != nil {
		cmnmsg.WriteErrorResponse(c.Writer, err)
		return
	}

	if session.Openid == "" || session.Session_key == "" {
		render.WriteJSON(c.Writer, cmnmsg.NewIllegalArgResponse("code"))
		return
	}

	//save session to redis
	log4g.Debug("session_key=" + session.Session_key)
	log4g.Debug("expires_in=%d", session.ExpiresIn)
	log4g.Debug("openid=" + session.Openid)
	log4g.Debug("unionid=" + session.Unionid)


	accessToken := util.NewUuid()
	if err = service.SaveLoginSession(accessToken, session); err != nil {
		cmnmsg.WriteErrorResponse(c.Writer, err)
		return
	}

	var result vo.WxappLoginToken
	result.AccessToken = accessToken
	result.ExpiresIn = session.ExpiresIn
	cmnmsg.WriteDataResponse(c.Writer, &result)

}

func code2session(code string) (session *vo.WxappSession, err error) {

	log4g.Debug("login with code %s", code)

	url := fmt.Sprintf(WEAPP_URL_FORMAT, WEAPP_APPID, WEAPP_SECRET, code)
	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		log4g.Error(err)
		return
	}
	defer resp.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log4g.Error(err)
		return
	}
	log4g.Debug(string(body))
	session = new(vo.WxappSession)
	err = json.Unmarshal(body, session)
	if err != nil {
		log4g.Error(err)
		return
	}
	if session.Errcode > 0 {
		log4g.Error(session.Errmsg)
		err = errors.New(session.Errmsg)
		//for testing
		err = nil
		session.Session_key = "ot96OBqsvSa3WLFBz4U+gw=="
		session.ExpiresIn = 7200
		session.Openid = "oVecO0Ze4kNxMGymF05d1uiIcmqA"
	}
	return
}
