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
	"common/util"
	"usercenter/service"
	"github.com/google/uuid"
	"strings"
	"common/tokenutil"
	"common/model"
)

func SignUp(c *gin.Context) {
	sgc := sg.Context(c)
	signUpVo := new(vo.SignUpVo)
	if sgc.CheckError(c.BindJSON(signUpVo)) {
		return
	}
	if sgc.CheckParamEmpty(signUpVo.PhoneNumber, "phone_number") || sgc.CheckParamEmpty(signUpVo.Password, "password") {
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
	userId := strings.Replace(uuid.New().String(), "-", "", -1)
	if sgc.CheckError(dao.NewUserDao().Insert(userId, signUpVo.PhoneNumber, signUpVo.Password)) {
		return
	}

	result := new(vo.LoginToken)
	result.AccessToken = util.NewUuid()
	result.ExpiresIn = 7200
	sgc.WriteDataOrError(result, service.SaveLoginSession(result.AccessToken, userId, result.ExpiresIn))
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

func Me(c *gin.Context) {
	sgc := sg.Context(c)
	userId, err := tokenutil.GetUserId(c)
	if sgc.CheckError(err) {
		return
	}
	var user model.User
	if sgc.CheckError(dao.NewUserDao().GetById(userId, &user)) {
		return
	}
	var userVo vo.UserVo
	userVo.Id = user.Id
	userVo.PhoneNumber = user.PhoneNumber.String()
	userVo.NickName = user.NickName.String()
	userVo.Email = user.Email.String()
	userVo.Signature = user.Signature.String()
	userVo.Avatar = user.Avatar.String()
	userVo.Sex = user.Sex.Int8()
	userVo.InsertTime = user.InsertTime
	userVo.UpdateTime = user.UpdateTime
	sgc.WriteData(&userVo)
}