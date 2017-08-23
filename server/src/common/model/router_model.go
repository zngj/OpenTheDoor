package model

import (
	"common/sqlx"
	"time"
)

type GateInfo struct {
	Id          string `db:"id"`
	Direction   int8   `db:"direction"`
	StationId   string `db:"station_id"`
	StationName string `db:"station_name"`
	CityId      string `db:"city_id"`
	CityName    string `db:"city_name"`
}

type RouterEvidence struct {
	EvidenceId string     `db:"evidence_id"`
	UserId     string     `db:"user_id"`
	Direction  int8       `db:"direction"`
	CreateTime time.Time  `db:"create_time"`
	ExpiresAt  time.Time  `db:"expires_at"`
	Status     int8       `db:"status"`
	UpdateTime *time.Time `db:"update_time"`
}

type RouterInfo struct {
	Id             int64            `db:"id"`
	UserId         string           `db:"user_id"`
	AtDate         time.Time        `db:"at_date"`
	GroupNo        int16            `db:"group_no"`
	InStationId    sqlx.NullString  `db:"in_station_id"`
	InStationName  sqlx.NullString  `db:"in_station_name"`
	InGateId       sqlx.NullString  `db:"in_gate_id"`
	InEvidence     sqlx.NullString  `db:"in_evidence"`
	InTime         *time.Time       `db:"in_time"`
	OutStationId   sqlx.NullString  `db:"out_station_id"`
	OutStationName sqlx.NullString  `db:"out_station_name"`
	OutGateId      sqlx.NullString  `db:"out_gate_id"`
	OutEvidence    sqlx.NullString  `db:"out_evidence"`
	Money          sqlx.NullFloat64 `db:"money"`
	OutTime        *time.Time       `db:"out_time"`
	OutGroup       sqlx.NullInt64   `db:"out_group"`
	Status         sqlx.NullInt64   `db:"status"`
	Paid           bool             `db:"paid"`
	ExceptionTime  *time.Time       `db:"exception_time"`
}

type Notification struct {
	Id         uint64     `db:"id"`
	UserId     string     `db:"user_id"`
	Type       int8       `db:"type"`
	Received   bool       `db:"received"`
	InsertTime *time.Time `db:"insert_time"`
}
