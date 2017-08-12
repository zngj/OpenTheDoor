package token

import (
	"github.com/gin-gonic/gin"
	"common/vo"
	"github.com/carsonsx/log4g"
	"common/errcode"
	"common/sg"
)

const HEADER_ACCESS_TOKEN  = "Access-Token"

func VerifyToken(c *gin.Context) {
	sgc := sg.Context(c)
	if CheckToken(c) {
		sgc.WriteSuccess()
	}
}

func CheckToken(c *gin.Context) bool {
	sgc := sg.Context(c)
	var verify vo.VerifyToken
	verify.AccessToken = c.Request.Header.Get(HEADER_ACCESS_TOKEN)
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
	valid, err := IsValid(verify.AccessToken)
	if sgc.CheckError(err) {
		return false
	}
	if !valid {
		sgc.Write(errcode.CODE_UC_TOKEN_EXPIRED)
		return false
	}
	return true
}

func VerifyTokenFn(c *gin.Context) {
	if !CheckToken(c) {
		log4g.Info("access token is expired")
		c.Abort()
	}
}