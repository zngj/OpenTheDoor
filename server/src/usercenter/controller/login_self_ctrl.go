package controller

import (
	"github.com/gin-gonic/gin"
	"common/sg"
	"common/vo"
	"github.com/carsonsx/log4g"
	"common/errcode"
	"fmt"
	"crypto/md5"
	"common/model"
	"common/sqlx"
	"usercenter/dao"
	"common/util"
	"usercenter/service"
)

func Login(c *gin.Context) {
	sgc := sg.Context(c)
	loginVo := new(vo.LoginVo)
	if sgc.CheckError(c.BindJSON(loginVo)) {
		return
	}
	if sgc.CheckParamEmpty(loginVo.PhoneNumber, "phone_number") || sgc.CheckParamEmpty(loginVo.Password, "password") {
		return
	}
	log4g.Info("login with phone number: %s", loginVo.PhoneNumber)
	var user model.User
	if err := dao.NewUserDao().GetByPhoneNumber(loginVo.PhoneNumber, &user); err != nil {
		if err == sqlx.ErrNotFound {
			err = errcode.SGErrMobileNotFound
			log4g.Debug(err)
		}
		sgc.WriteError(err)
		return
	}
	if user.Password.String() != fmt.Sprintf("%x", md5.Sum([]byte(loginVo.Password))) {
		err := errcode.SGErrWrongPassword
		log4g.Error(err)
		sgc.WriteError(err)
		return
	}

	result := new(vo.LoginToken)
	result.AccessToken = util.NewUuid()
	result.ExpiresIn = 7200
	sgc.WriteDataOrError(result, service.SaveLoginSession(result.AccessToken, user.Id, result.ExpiresIn))
}


func TestLogin(c *gin.Context) {

	sgc := sg.Context(c)
	param := "phone_number"
	phoneNumber := c.Query(param)
	if sgc.CheckParamEmpty(phoneNumber, param) {
		return
	}
	log4g.Info("test login with phone number: %s", phoneNumber)
	var user model.User
	if err := dao.NewUserDao().GetByPhoneNumber(phoneNumber, &user); err != nil {
		if err == sqlx.ErrNotFound {
			err = errcode.SGErrMobileNotFound
			log4g.Debug(err)
		}
		sgc.WriteError(err)
		return
	}
	result := new(vo.LoginToken)
	result.AccessToken = util.NewUuid()
	result.ExpiresIn = 7200
	sgc.WriteDataOrError(result, service.SaveLoginSession(result.AccessToken, user.Id, result.ExpiresIn))
}