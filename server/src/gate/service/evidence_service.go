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
