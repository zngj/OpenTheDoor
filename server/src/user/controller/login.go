package controller

import (
	"common/cmnmsg"
	"encoding/json"
	"fmt"
	"github.com/carsonsx/log4g"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strings"
	"user/service"
	"user/vo"
	"errors"
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
		render.WriteJSON(c.Writer, cmnmsg.NewErrorResponse(err))
		return
	}
	if login.Code == "" {
		render.WriteJSON(c.Writer, cmnmsg.NewEmptyArgResponse("code"))
		return
	}

	session, err := code2session(login.Code)
	if err != nil {
		render.WriteJSON(c.Writer, cmnmsg.NewErrorResponse(err))
		return
	}

	if session.Openid == "" || session.Session_key == "" {
		render.WriteJSON(c.Writer, cmnmsg.NewIllegalArgResponse("code"))
		return
	}

	//save session to redis
	log4g.Debug("openId=" + session.Openid)
	log4g.Debug("session_key=" + session.Session_key)
	log4g.Debug("unionid=" + session.Unionid)

	var result vo.WxappLoginToken
	result.Token = strings.Replace(uuid.New().String(), "-", "", -1)

	if err = service.SaveLoginSession(result.Token, session); err != nil {
		render.WriteJSON(c.Writer, cmnmsg.NewErrorResponse(err))
		return
	}

	render.WriteJSON(c.Writer, cmnmsg.NewDataResponse(&result))

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
		//err = nil
		//session.Openid = "openid_for_testing_1111111"
		//session.Session_key = "session_key_for_testing_11111111"
	}
	return
}
