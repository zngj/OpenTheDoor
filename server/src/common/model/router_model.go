package model

import "time"

type GateInfo struct {
	Id          string
	Direction   int8
	StationCode string
	StationName string
	CityCode    string
	CityName    string
}

type RouterEvidence struct {
	EvidenceId string
	UserId string
	Type int8
	CreateTime time.Time
	ExpiresAt time.Time
	Status int16
	UsedTime *time.Time
}

