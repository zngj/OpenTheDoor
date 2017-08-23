package dao

import (
	"common/model"
	"common/sqlx"
	"common/sgconst"
	"common/util"
	"time"
	"bytes"
	"github.com/carsonsx/log4g"
)

func NewRouterDao(dao ...*sqlx.Dao) *routerDao {
	d := new(routerDao)
	if len(dao) > 0 {
		d.sqlxDao = dao[0]
	} else {
		d.sqlxDao = new(sqlx.Dao)
	}
	return d
}

type routerDao struct {
	sqlxDao *sqlx.Dao
}

func (d *routerDao) GetOngoing(userId string, router *model.RouterInfo) error {
	sql := "select * from sg_router_info where user_id=? and status <> 2 order by id asc"
	return d.sqlxDao.Query(sql, userId).One(router)
}

func (d *routerDao) GetIn(userId string, router *model.RouterInfo) error {
	sql := "select * from sg_router_info where user_id=? and status = 1"
	return d.sqlxDao.Query(sql, userId).One(router)
}

func (d *routerDao) GetException(userId string, router *model.RouterInfo) error {
	sql := "select * from sg_router_info where user_id=? and status in (4,5)"
	return d.sqlxDao.Query(sql, userId).One(router)
}

func (d *routerDao) InsertIn(router *model.RouterInfo) error {
	sql := "insert sg_router_info (user_id,at_date,group_no,in_station_id,in_station_name,in_gate_id,in_evidence,in_time,status) values (?,?,?,?,?,?,?,?,?)"
	return d.sqlxDao.Exec(sql, router.UserId, time.Now(), router.GroupNo, router.InStationId, router.InStationName, router.InGateId, router.InEvidence, router.InTime, router.Status)
}

func (d *routerDao) UpdateOut(router *model.RouterInfo) error {
	sql := "update sg_router_info set out_station_id=?,out_station_name=?,out_gate_id=?,out_evidence=?,out_time=?,money=?,paid=?,status=? where id=?"
	return d.sqlxDao.Exec(sql, router.OutStationId, router.OutStationName, router.OutGateId, router.OutEvidence, router.OutTime, router.Money, router.Paid, router.Status, router.Id)
}

func (d *routerDao) InsertOut(router *model.RouterInfo) error {
	sql := "insert sg_router_info (user_id,date,group,out_station_id,out_station_name,out_gate_id,out_evidence,out_time,status) values (?,?,?,?,?,?,?,?,?)"
	return d.sqlxDao.Exec(sql, router.UserId, time.Now(), router.GroupNo, router.OutStationId, router.OutStationName, router.OutGateId, router.OutEvidence, router.OutTime, router.Status)
}

func (d *routerDao) UpdateIn(router *model.RouterInfo) error {
	sql := "update sg_router_info set in_station_id=?,in_station_name=?,in_gate_id=?,in_evidence=?,in_time=?,money=?,paid=?,status=? where id=?"
	return d.sqlxDao.Exec(sql, router.InStationId, router.InStationName, router.InGateId, router.InEvidence, router.InTime, router.Money, router.Paid, router.Status, router.Id)
}

func (d *routerDao) GetCurrentGroupNo(userId string) (groupNo int16, err error) {
	sql := "select group_no from sg_router_info where user_id=? and at_date=? and status = ? order by id desc limit 1"
	err = d.sqlxDao.Query(sql, userId, util.NowDate(),sgconst.ROUTER_STATUS_NORMAL_IN).Scan(&groupNo)
	log4g.ErrorIf(err)
	return
}

func (d *routerDao) FindIn(userId string, routers *[]*model.RouterInfo) error {
	sql := "select * from sg_router_info where status = ? and user_id=? and at_date=? order by id asc"
	return d.sqlxDao.Query(sql, sgconst.ROUTER_STATUS_NORMAL_IN, userId, util.NowDate()).All(routers)
}

func (d *routerDao) FindOut(userId string, routers *[]*model.RouterInfo) error {
	//get last group no
	sql := "select group_no from sg_router_info where user_id=? and at_date=? order by id desc limit 1"
	var groupNo int16
	err := d.sqlxDao.Query(sql, userId, util.NowDate()).Scan(&groupNo)
	if err != nil {
		return err
	}
	sql = "select * from sg_router_info where user_id=? and group_no=? and at_date=? order by id asc"
	return d.sqlxDao.Query(sql, userId, groupNo, util.NowDate()).All(routers)
}

func (d *routerDao) FindByUser(userId string, lastId int64, size int, routers *[]*model.RouterInfo) error {
	var sql bytes.Buffer
	var values []interface{}
	sql.WriteString("select * from sg_router_info where user_id=?")
	values = append(values, userId)
	if lastId > 0 {
		sql.WriteString(" and id < ?")
		values = append(values, lastId)
	}
	sql.WriteString(" order by id desc limit ?")
	values = append(values,  size)
	return d.sqlxDao.Query(sql.String(), values...).All(routers)
}