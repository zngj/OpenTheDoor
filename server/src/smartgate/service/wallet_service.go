package service

import (
	"smartgate/model"
	"smartgate/dao"
	"common/mysqlx"
)

func GetWallet(userId string) (wallet *model.WalletInfo, err error) {
	dao := dao.NewWalletDao()
	wallet, err = dao.GetByUserId(userId)
	if err == mysqlx.ErrNotFound {
		wallet = new(model.WalletInfo)
		wallet.UserId = userId
		wallet.Balance = 100
		err = dao.Insert(wallet)
	}
	return
}
