package service

import (
	"common/model"
	"common/dbx"
	"gate/msg"
	"common/errcode"
)

func GetGateInfo(id string) (gate *model.GateInfo, err error) {
	dao := dbx.NewDao()
	gate = new(model.GateInfo)
	sql := "select id,direction,station_name,city_name from sg_gate_info where id = ?"
	err = dao.Query(sql, id).One(&gate.Id, &gate.Direction, &gate.StationName, &gate.CityName)
	return
}

// 3201 凭证不存在
// 3202 凭证已过期
// 3203 凭证与机闸不匹配
// 3204 用户不符合付费标准
func VerifyEvidence(evidenceId, gateId string) (code int, err error) {
	dao := dbx.NewDao()
	var evidence model.RouterEvidence
	sql := "select user_id,type,expires_at from sg_router_evidence where evidence_id = ?"
	err = dao.Query(sql, evidenceId).One(&evidence.UserId, evidence.Type, &evidence.ExpiresAt)
	if err == dbx.ErrNotFound {
		code = errcode.CODE_GATE_INVALID_GATE
		err = nil
		return
	} else  if err != nil {
		return
	}

	var gate *model.GateInfo
	gate, err = GetGateInfo(gateId)
	if err != nil {
		return
	}

	if evidence.Type != 3 && gate.Direction != evidence.Type {
		code = errcode.CODE_GATE_NOT_MATCH_EVIDENCE
		return
	}

	//Check user pay status


	return
}

func SubmitEvidence(evidence *msg.C2SSubmitEvidence) (err error) {



	return
}