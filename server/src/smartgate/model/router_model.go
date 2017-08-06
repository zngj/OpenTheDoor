package model

import "time"

type RouterEvidence struct {
	EvidenceId string
	UserId string
	CreateTime time.Time
	ExpiresAt time.Time
	Status int16
	UsedTime *time.Time
}
