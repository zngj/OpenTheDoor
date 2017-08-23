package controller

import (
	"common/errcode"
	"common/sg"
	"common/vo"
	"github.com/carsonsx/log4g"
	"github.com/gin-gonic/gin"
	"usercenter/service"
	"common/tokenutil"
)

func CheckToken(c *gin.Context) {
	sgc := sg.Context(c)
	if _checkToken(c) {
		sgc.WriteSuccess()
	}
}

func _checkToken(c *gin.Context) bool {
	sgc := sg.Context(c)
	var verify vo.VerifyToken
	verify.AccessToken = c.Request.Header.Get(tokenutil.HEADER_ACCESS_TOKEN)
	if verify.AccessToken == "" {
		err := c.Bind(&verify)
		if sgc.CheckError(err) {
			return false
		}
	} else {
		log4g.Debug("found access_token in header: %s", verify.AccessToken)
	}
	if sgc.CheckParamEmpty(verify.AccessToken) {
		return false
	}
	valid, err := service.IsValid(verify.AccessToken)
	if sgc.CheckError(err) {
		return false
	}
	if !valid {
		sgc.Write(errcode.CODE_UC_TOKEN_EXPIRED)
		return false
	}
	return true
}