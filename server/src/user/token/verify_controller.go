package token

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"common/cmnmsg"
	"user/vo"
	"user/usercode"
	"github.com/carsonsx/log4g"
)

func VerifyToken(c *gin.Context) {
	if CheckToken(c) {
		render.WriteJSON(c.Writer, cmnmsg.NewSuccessResponse())
	}
}

func CheckToken(c *gin.Context) bool {
	var verify vo.VerifyToken
	err := c.Bind(&verify)
	if err != nil {
		log4g.Error(err)
		render.WriteJSON(c.Writer, cmnmsg.NewErrorResponse(err))
		return false
	}
	if verify.AccessToken == "" {
		resp := cmnmsg.NewEmptyArgResponse("code")
		log4g.Error(resp.Msg)
		render.WriteJSON(c.Writer, resp)
		return false
	}
	valid, err := IsValidToken(verify.AccessToken)
	if err != nil {
		log4g.Error(err)
		render.WriteJSON(c.Writer, cmnmsg.NewErrorResponse(err))
		return false
	}
	if !valid {
		resp :=  usercode.NewUserTokenExpiredResponse()
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