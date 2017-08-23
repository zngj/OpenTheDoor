package dao

import (
	"common/model"
	"common/sqlx"
	"time"
)

func NewEvidenceDao(sqlxDao ...*sqlx.Dao) *evidenceDao {
	d := new(evidenceDao)
	if len(sqlxDao) > 0 {
		d.sqlxDao = sqlxDao[0]
	} else {
		d.sqlxDao = new(sqlx.Dao)
	}
	return d
}

type evidenceDao struct {
	sqlxDao *sqlx.Dao
}

func (d *evidenceDao) Insert(evidence *model.RouterEvidence) error {
	sql := "insert sg_router_evidence (evidence_id,user_id,direction,create_time,expires_at,status) values (?,?,?,?,?,?)"
	return d.sqlxDao.Exec(sql, evidence.EvidenceId, evidence.UserId, evidence.Direction, evidence.CreateTime, evidence.ExpiresAt, evidence.Status)
}

func (d *evidenceDao) GetUserId(evidenceId string, userId *string) error {
	sql := "insert sg_router_evidence (evidence_id,user_id,direction,create_time,expires_at,status) values (?,?,?,?,?,?)"
	return d.sqlxDao.Query(sql, evidenceId).Scan(&userId)
}

func (d *evidenceDao) Get(evidenceId string, evidence *model.RouterEvidence) error {
	sql := "select evidence_id,user_id,direction,expires_at,status from sg_router_evidence where evidence_id = ?"
	return d.sqlxDao.Query(sql, evidenceId).One(evidence)
}

func (d *evidenceDao) Consume(evidenceId string) error {
	sql := "update sg_router_evidence set status=2,update_time=? where evidence_id=?"
	return d.sqlxDao.Exec(sql, time.Now(), evidenceId)
}

func (d *evidenceDao) Discard(evidenceId string) error {
	sql := "update sg_router_evidence set status=4,update_time=? where evidence_id=?"
	return d.sqlxDao.Exec(sql, time.Now(), evidenceId)
}

func (d *evidenceDao) IsValidEvidence(evidenceId string) (bool, error) {
	return true, nil
}
