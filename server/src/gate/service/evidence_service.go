package service

import (
	"common/model"
	"common/dbx"
)

func GetGateInfo(id string) (gate *model.GateInfo, err error) {
	dao := dbx.NewDao()
	gate = new(model.GateInfo)
	err = dao.Query("select id,direction,station_name,city_name from sg_gate_info where id = ?", id).
		One(&gate.Id, &gate.Direction, &gate.StationName, &gate.CityName)
	return
}

// 3201 凭证不存在
// 3202 凭证已过期
// 3203 凭证与机闸不匹配
// 3204 用户不符合付费标准
func VerifyEvidence(evidence string) (code int, err error) {



	return
}