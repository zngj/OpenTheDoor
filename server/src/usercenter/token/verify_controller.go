package token

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"usercenter/vo"
	"github.com/carsonsx/log4g"
	"common/errcode"
)

const HEADER_ACCESS_TOKEN  = "Access-Token"

func VerifyToken(c *gin.Context) {
	if CheckToken(c) {
		errcode.WriteSuccessResponse(c.Writer)
	}
}

func CheckToken(c *gin.Context) bool {
	var verify vo.VerifyToken
	verify.AccessToken = c.Request.Header.Get(HEADER_ACCESS_TOKEN)
	if verify.AccessToken == "" {
		err := c.Bind(&verify)
		if err != nil {
			log4g.Error(err)
			errcode.WriteErrorResponse(c.Writer, err)
			return false
		}
	} else {
		log4g.Debug("found access_token in header: %s", verify.AccessToken)
	}
	if verify.AccessToken == "" {
		errcode.WriteEmptyArgResponse(c.Writer, "code")
		return false
	}
	valid, err := IsValid(verify.AccessToken)
	if err != nil {
		log4g.Error(err)
		errcode.WriteErrorResponse(c.Writer, err)
		return false
	}
	if !valid {
		resp :=  errcode.NewResponse(errcode.CODE_UC_TOKEN_EXPIRED)
		log4g.Info(resp.Msg)
		render.WriteJSON(c.Writer, resp)
		return false
	}
	return true
}

func VerifyTokenFn(c *gin.Context) {
	log4g.Debug("checking token...")
	if !CheckToken(c) {
		c.Abort()
	}
}