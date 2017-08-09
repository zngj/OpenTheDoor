package dao

import (
	"common/model"
	"common/dbx"
)

func GetGateInfo(id string) (gate *model.GateInfo, err error) {
	gate = new(model.GateInfo)
	sql := "select id,direction,station_id,station_name,city_id,city_name from sg_gate_info where id = ?"
	err = dbx.NewDao().Query(sql, id).One(&gate.Id, &gate.Direction, &gate.StationId, &gate.StationName, &gate.CityId, &gate.CityName)
	return
}
