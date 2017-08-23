package dao

import (
	"common/model"
	"common/sqlx"
)

func NewGateDao(sqlxDao ...*sqlx.Dao) *gateDao {
	d := new(gateDao)
	if len(sqlxDao) > 0 {
		d.sqlxDao = sqlxDao[0]
	} else {
		d.sqlxDao = new(sqlx.Dao)
	}
	return d
}

type gateDao struct {
	sqlxDao *sqlx.Dao
}

func (d *gateDao) Get(id string, gate *model.GateInfo) error {
	sql := "select id,direction,station_id,station_name,city_id,city_name from sg_gate_info where id = ?"
	return d.sqlxDao.Query(sql, id).One(gate)
}
