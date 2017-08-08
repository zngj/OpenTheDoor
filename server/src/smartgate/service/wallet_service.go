package service

import (
	"smartgate/dao"
	"common/dbx"
	"common/model"
)

func GetWallet(userId string) (wallet *model.WalletInfo, err error) {
	dao := dao.NewWalletDao()
	wallet, err = dao.GetByUserId(userId)
	if err == dbx.ErrNotFound {
		wallet = new(model.WalletInfo)
		wallet.UserId = userId
		wallet.Balance = 100
		err = dao.Insert(wallet)
	}
	return
}
