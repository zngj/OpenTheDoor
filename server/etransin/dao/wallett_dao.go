package dao

import (
	"common/model"
	"common/sqlx"
	"time"
)

func NewWalletDao(dao ...*sqlx.Dao) *walletDao {
	d := new(walletDao)
	if len(dao) > 0 {
		d.dao = dao[0]
	} else {
		d.dao = new(sqlx.Dao)
	}
	return d
}

type walletDao struct {
	dao *sqlx.Dao
}

func (d *walletDao) GetByUserId(userId string, wallet *model.WalletInfo) error {
	return d.dao.Query("select balance,wxpay_quick from sg_wallet_info where user_id = ?", userId).One(wallet)
}

func (d *walletDao) Insert(wallet *model.WalletInfo) error {
	return d.dao.Exec("insert sg_wallet_info (user_id,balance,insert_time) values (?,?,?)", wallet.UserId, wallet.Balance, time.Now())
}

func (d *walletDao) Consume(userId string, amount float32) error {
	return d.dao.Exec("update sg_wallet_info set balance = balance-?, update_time=? where user_id=?", amount, time.Now(), userId)
}

func (d *walletDao) Charge(userId string, amount float32) error {
	return d.dao.Exec("update sg_wallet_info set balance = balance+?, update_time=? where user_id=?", amount, time.Now(), userId)
}
