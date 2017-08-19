package controller

import (
	"common/model"
	"common/sg"
	"common/vo"
	"github.com/carsonsx/log4g"
	"github.com/gin-gonic/gin"
	"etransin/dao"
	"etransin/service"
	"common/tokenutil"
)

func WalletInfo(c *gin.Context) {
	sgc := sg.Context(c)
	userId, err := tokenutil.GetUserId(c)
	if sgc.CheckError(err) {
		log4g.Error(err)
		return
	}
	var wallet model.WalletInfo
	if sgc.CheckError(service.GetWallet(userId, &wallet)) {
		log4g.Error(err)
		return
	}
	var walletVo vo.WalletVO
	walletVo.Balance = wallet.Balance
	walletVo.WxpayQuick = wallet.WxpayQuick
	log4g.Debug(log4g.JsonFunc(&walletVo))
	sgc.WriteData(&walletVo)
}

func WalletCharge(c *gin.Context) {
	sgc := sg.Context(c)
	var vo vo.WalletChargeVO
	err := c.BindJSON(&vo)
	if sgc.CheckError(err) {
		log4g.Error(err)
		return
	}
	userId, err := tokenutil.GetUserId(c)
	if sgc.CheckError(err) {
		log4g.Error(err)
		return
	}
	sgc.WriteSuccessOrError(dao.NewWalletDao().Charge(userId, vo.Money))
}
