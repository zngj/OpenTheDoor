package dao

import (
	"time"
	"common/dbx"
	"common/model"
)

func NewWalletDao() *walletDao {
	d := new(walletDao)
	d.dao = new(dbx.Dao)
	return d
}

type walletDao struct {
	dao *dbx.Dao
}

func (d *walletDao) GetByUserId(userId string) (wallet *model.WalletInfo, err error) {
	var balance float32
	var wxpayQuick bool
	err = d.dao.Query("select balance,wxpay_quick from sg_wallet_info where user_id = ?", userId).
		One(&balance, &wxpayQuick)
	if err != nil {
		return
	}
	wallet = new(model.WalletInfo)
	wallet.UserId = userId
	wallet.Balance = balance
	wallet.WxpayQuick = wxpayQuick
	return
}

func (d *walletDao) Insert(wallet *model.WalletInfo) error {
	return d.dao.Exec("insert sg_wallet_info (user_id,balance,insert_time) values (?,?,?)", wallet.UserId, wallet.Balance, time.Now())
}

func (d *walletDao) Consume(userId string, amount float32) error {
	return d.dao.Exec("update sg_wallet_info set balance = balance-?, update_time=? where user_id=?", amount, time.Now(), userId)
}

func WalletCharge(userId string, amount float32) error {
	return dbx.NewDao().Exec("update sg_wallet_info set balance = balance+?, update_time=? where user_id=?", amount, time.Now(), userId)
}