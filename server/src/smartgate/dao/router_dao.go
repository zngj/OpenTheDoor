package dao

import (
	"common/dbx"
	"common/model"
)

const (
	SQL_EVIDENCE_INSERT        = "insert sg_router_evidence (evidence_id,user_id,direction,create_time,expires_at,status) values (?,?,?,?,?,?)"
	SQL_GET_USERID_BY_EVIDENCE = "select user_id from sg_router_evidence where evidence_id = ? "
)

func NewRouterDao() *routerDao {
	d := new(routerDao)
	d.dao = new(dbx.Dao)
	return d
}

type routerDao struct {
	dao *dbx.Dao
}

func (d *routerDao) InsertEvidence(evidence *model.RouterEvidence) error {
	return d.dao.Exec(SQL_EVIDENCE_INSERT, evidence.EvidenceId, evidence.UserId, evidence.Direction, evidence.CreateTime, evidence.ExpiresAt, evidence.Status)
}

func (d *routerDao) IsValidEvidence(evidenceId string) (bool, error) {
	return true, nil
}

func (d *routerDao) GetUserIdByEvidenceId(evidenceId string) (userId string, err error) {
	err = d.dao.Query(SQL_GET_USERID_BY_EVIDENCE, evidenceId).One(&userId)
	return
}

func GetRouterEvidence(evidenceId string) (evidence *model.RouterEvidence, err error) {
	evidence = new(model.RouterEvidence)
	_dao := dbx.NewDao()
	sql := "select user_id,direction,expires_at from sg_router_evidence where evidence_id = ?"
	err = _dao.Query(sql, evidenceId).One(&evidence.UserId, &evidence.Direction, &evidence.ExpiresAt)
	return
}

func GetOngoingRouterInfo(userId string) (router *model.RouterInfo, err error) {
	router = new(model.RouterInfo)
	sql := "select id,user_id,status from sg_router_info where user_id=? and status <> 2 order by id desc"
	err = dbx.NewDao().Query(sql, userId).One(&router.Id, &router.UserId, &router.Status)
	return
}

func GetExceptionRouterInfo(userId string) (router *model.RouterInfo, err error) {
	router = new(model.RouterInfo)
	sql := "select id,status from sg_router_info where user_id=? and status in (4,5)"
	err = dbx.NewDao().Query(sql, userId).One(&router.Id, &router.Status)
	return
}

func InsertRouteInfoOfIn(router *model.RouterInfo) error {
	sql := "insert sg_router_info (user_id,in_station_id,in_station_name,in_gate_id,in_evidence,in_time,status) values (?,?,?,?,?,?,?)"
	return dbx.NewDao().Exec(sql, router.UserId, router.InStationId, router.InStationName, router.InGateId, router.InEvidence, router.InTime, router.Status)
}

func UpdateRouteInfoOfOut(router *model.RouterInfo) error {
	sql := "update sg_router_info set out_station_id=?,out_station_name=?,out_gate_id=?,out_evidence=?,out_time=?,money=?,status=? where id=?"
	return dbx.NewDao().Exec(sql, router.OutStationId, router.OutStationName, router.OutGateId, router.OutEvidence, router.OutTime, router.Money, router.Status, router.Id)
}

func InsertRouteInfoOfOut(router *model.RouterInfo) error {
	sql := "insert sg_router_info (user_id,out_station_id,out_station_name,out_gate_id,out_evidence,out_time,status) values (?,?,?,?,?,?,?)"
	return dbx.NewDao().Exec(sql, router.UserId, router.OutStationId, router.OutStationName, router.OutGateId, router.OutEvidence, router.OutTime, router.Status)
}

func UpdateRouteInfoOfIn(router *model.RouterInfo) error {
	sql := "update sg_router_info set in_station_id=?,in_station_name=?,in_gate_id=?,in_evidence=?,in_time=?,money=?,status=? where id=?"
	return dbx.NewDao().Exec(sql, router.InStationId, router.InStationName, router.InGateId, router.InEvidence, router.InTime, router.Money, router.Status, router.Id)
}

func GetNotification(userId, category string) (n *model.Notification, err error) {
	n = new(model.Notification)
	sql := "select id,category,content_id from sg_sys_notification where received=0 and category =? and user_id=? order by id asc"
	err = dbx.NewDao().Query(sql, category, userId).One(&n.Id, &n.Category, &n.ContentId)
	return
}

func ConsumeNotification(id uint64) error {
	return dbx.NewDao().Exec("update sg_sys_notification set received=1 where id=?", id)
}
