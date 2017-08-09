package model

import "time"

type GateInfo struct {
	Id          string
	Direction   int8
	StationId   string
	StationName string
	CityId      string
	CityName    string
}

type RouterEvidence struct {
	EvidenceId string
	UserId     string
	Direction  int8
	CreateTime time.Time
	ExpiresAt  time.Time
	Status     int16
	UsedTime   time.Time
}

type RouterInfo struct {
	Id             int64
	UserId         string
	InStationId    string
	InStationName  string
	InGateId       string
	InEvidence     string
	InTime         time.Time
	OutStationId   string
	OutStationName string
	OutGateId      string
	OutEvidence    string
	OutTime        time.Time
	Status         int8
	Money          float32
	ExceptionTime  time.Time
}

type Notification struct {
	Id        uint64
	UserId    string
	Category  string
	ContentId string
	Received  bool
}
