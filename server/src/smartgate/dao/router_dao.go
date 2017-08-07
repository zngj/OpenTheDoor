package dao

import (
	"common/mysqlx"
	"smartgate/model"
)

const (

	SQL_EVIDENCE_INSERT = "insert sg_router_evidence (evidence_id,user_id,type,create_time,expires_at,status) values (?,?,?,?,?,?)"
	SQL_GET_USERID_BY_EVIDENCEID = "select user_id from sg_router_evidence where evidence_id = ? "

)

func NewRouterDao() *routerDao {
	d := new(routerDao)
	d.dao = new(mysqlx.Dao)
	return d
}

type routerDao struct {
	dao *mysqlx.Dao
}

func (d *routerDao) InsertEvidence(evidence *model.RouterEvidence) error {
	return d.dao.Exec(SQL_EVIDENCE_INSERT, evidence.EvidenceId, evidence.UserId, evidence.Type, evidence.CreateTime, evidence.ExpiresAt, evidence.Status)
}

func (d *routerDao) IsValidEvidence(evidenceId string) (bool, error) {
	return true, nil
}

func (d *routerDao) GetUserIdByEvidenceId(evidenceId string) (userId string, err error) {
	err = d.dao.Query(SQL_GET_USERID_BY_EVIDENCEID, evidenceId).Result(&userId)
	return
}
