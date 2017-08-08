package controller

import (
	"github.com/gin-gonic/gin"
	"usercenter/token"
	"smartgate/service"
	"smartgate/vo"
	"common/errcode"
)

func WalletInfo(c *gin.Context) {
	userId, err := token.GetUserIdFromHeader(c.Request.Header)
	if err != nil {
		errcode.WriteErrorResponse(c.Writer, err)
		return
	}
	wallet, err := service.GetWallet(userId)
	if err != nil {
		errcode.WriteErrorResponse(c.Writer, err)
		return
	}
	var walletVo vo.WalletVO
	walletVo.Balance = wallet.Balance
	walletVo.WxpayQuick = wallet.WxpayQuick

	errcode.WriteDataResponse(c.Writer, &walletVo)
}


