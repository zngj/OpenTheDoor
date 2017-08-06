package controller

import (
	"github.com/gin-gonic/gin"
	"usercenter/token"
	"common/cmnmsg"
	"smartgate/service"
	"smartgate/vo"
)

func WalletInfo(c *gin.Context) {
	userId, err := token.GetUserIdFromHeader(c.Request.Header)
	if err != nil {
		cmnmsg.WriteErrorResponse(c.Writer, err)
		return
	}
	wallet, err := service.GetWallet(userId)
	if err != nil {
		cmnmsg.WriteErrorResponse(c.Writer, err)
		return
	}
	var walletVo vo.WalletVO
	walletVo.Balance = wallet.Balance
	walletVo.WxpayQuick = wallet.WxpayQuick

	cmnmsg.WriteDataResponse(c.Writer, &walletVo)
}


