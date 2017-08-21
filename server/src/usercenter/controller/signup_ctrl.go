package controller

import (
	"github.com/gin-gonic/gin"
	"common/sg"
	"common/vo"
	"usercenter/dao"
	"common/errcode"
	"github.com/carsonsx/log4g"
	"crypto/md5"
	"fmt"
)

func SignUp(c *gin.Context) {
	sgc := sg.Context(c)
	signUpVo := new(vo.SignUpVo)
	if sgc.CheckError(c.BindJSON(signUpVo)) {
		return
	}
	if sgc.CheckParamEmpty(signUpVo.PhoneNumber, "mobile") || sgc.CheckParamEmpty(signUpVo.Password, "password") {
		return
	}

	log4g.Info("sign up with phone number: %s", signUpVo.PhoneNumber)

	exist, err := dao.NewUserDao().IsPhoneNumberExist(signUpVo.PhoneNumber)
	if sgc.CheckError(err) {
		return
	}
	if exist {
		err := errcode.SGErrMobileDuplicate
		log4g.Error(err)
		sgc.WriteError(err)
		return
	}
	signUpVo.Password = fmt.Sprintf("%x", md5.Sum([]byte(signUpVo.Password)))
	sgc.WriteSuccessOrError(dao.NewUserDao().Insert(signUpVo.PhoneNumber, signUpVo.Password))
}

func CheckPhoneNumber(c *gin.Context) {
	sgc := sg.Context(c)
	param := "phone_number"
	phoneNumber := c.Query(param)
	if sgc.CheckParamEmpty(phoneNumber, param) {
		return
	}
	log4g.Info("check phone number: %s", phoneNumber)
	exist, err := dao.NewUserDao().IsPhoneNumberExist(phoneNumber)
	if sgc.CheckError(err) {
		return
	}
	if exist {
		err := errcode.SGErrMobileDuplicate
		log4g.Error(err)
		sgc.WriteError(err)
		return
	}
	sgc.WriteSuccess()
}
