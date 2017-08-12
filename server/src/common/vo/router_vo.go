package vo

type RouterStatusVO struct {
	Status int8 `json:"status"`
}

type EvidenceVO struct {
	EvidenceId  string `json:"evidence_id"`
	EvidenceKey string `json:"evidence_key"`
	ExpiresAt   int64  `json:"expires_at"`
}

type RouterNotificationVo struct {
	NotificationId uint64  `json:"notification_id,omitempty"`
	Direction      int8    `json:"direction"`
	InGateId       string  `json:"in_gate_id,omitempty"`
	InStationId    string  `json:"in_station_id,omitempty"`
	InStationName  string  `json:"in_station_name,omitempty"`
	InTime         int64   `json:"in_time,omitempty"`
	OutGateId      string  `json:"out_gate_id,omitempty"`
	OutStationId   string  `json:"out_station_id,omitempty"`
	OutStationName string  `json:"out_station_name,omitempty"`
	OutTime        int64   `json:"out_time,omitempty"`
	Money          float32 `json:"money,omitempty"`
}