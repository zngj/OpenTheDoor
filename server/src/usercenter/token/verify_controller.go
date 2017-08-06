package token

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"common/cmnmsg"
	"usercenter/vo"
	"usercenter/usercode"
	"github.com/carsonsx/log4g"
	"net/http"
)

const HEADER_ACCESS_TOKEN  = "Access-Token"

func VerifyToken(c *gin.Context) {
	if CheckToken(c) {
		cmnmsg.WriteSuccessResponse(c.Writer)
	}
}

func CheckToken(c *gin.Context) bool {
	var verify vo.VerifyToken
	verify.AccessToken = c.Request.Header.Get(HEADER_ACCESS_TOKEN)
	if verify.AccessToken == "" {
		err := c.Bind(&verify)
		if err != nil {
			log4g.Error(err)
			cmnmsg.WriteErrorResponse(c.Writer, err)
			return false
		}
	} else {
		log4g.Debug("found access_token in header: %s", verify.AccessToken)
	}
	if verify.AccessToken == "" {
		resp := cmnmsg.NewEmptyArgResponse("code")
		log4g.Error(resp.Msg)
		render.WriteJSON(c.Writer, resp)
		return false
	}
	valid, err := IsValid(verify.AccessToken)
	if err != nil {
		log4g.Error(err)
		cmnmsg.WriteErrorResponse(c.Writer, err)
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

func GetUserIdFromHeader(header http.Header) (userId string, err error) {
	accessToken := header.Get(HEADER_ACCESS_TOKEN)
	return GetUserId(accessToken)
}