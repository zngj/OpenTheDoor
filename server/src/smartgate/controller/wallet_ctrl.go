package controller

import (
	"github.com/gin-gonic/gin"
	"smartgate/service"
	"common/vo"
	"common/sg"
	"smartgate/dao"
)

func WalletInfo(c *gin.Context) {
	sgc := sg.Context(c)
	userId, err := GetUserId(c)
	if sgc.CheckError(err) {
		return
	}
	wallet, err := service.GetWallet(userId)
	if sgc.CheckError(err) {
		return
	}
	var walletVo vo.WalletVO
	walletVo.Balance = wallet.Balance
	walletVo.WxpayQuick = wallet.WxpayQuick
	sgc.WriteData(&walletVo)
}

func WalletCharge(c *gin.Context)  {
	sgc := sg.Context(c)
	var vo vo.WalletChargeVO
	if sgc.CheckError(c.BindJSON(&vo)) {
		return
	}
	userId, err := GetUserId(c)
	if sgc.CheckError(err) {
		return
	}
	sgc.WriteSuccessOrError(dao.WalletCharge(userId, vo.Money))
}

