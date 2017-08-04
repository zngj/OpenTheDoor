package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"common/cmnmsg"
	"user/vo"
)

func VerifyToken(c *gin.Context) {

	var verify vo.VerifyToken
	err := c.Bind(&verify)
	if err != nil {
		render.WriteJSON(c.Writer, cmnmsg.NewErrorResponse(err))
		return
	}
	if verify.Token == "" {
		render.WriteJSON(c.Writer, cmnmsg.NewWrongArgResponse("code"))
		return
	}


}