package service

import (
	"common/model"
	"common/sqlx"
	"etransin/dao"
)

func GetWallet(userId string, wallet *model.WalletInfo) error {
	dao := dao.NewWalletDao()
	err := dao.GetByUserId(userId, wallet)
	if err == sqlx.ErrNotFound {
		wallet.UserId = userId
		wallet.Balance = 100
		err = dao.Insert(wallet)
	}
	return err
}

func ConsumeWallet(userId string, money float32) error {
	return dao.NewWalletDao().Consume(userId, money)
}
