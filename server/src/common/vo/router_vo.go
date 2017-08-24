package vo

type RouterStatusVO struct {
	Status int8 `json:"status"`
}

type RouterInfoVO struct {
	Id             int64   `json:"id"`
	UserId         string  `json:"-"`
	AtDate         int64   `json:"at_date"`
	InStationId    string  `json:"-"`
	InStationName  string  `json:"in_station_name,omitempty"`
	InGateId       string  `json:"-"`
	InTime         int64   `json:"in_time,omitempty"`
	OutStationId   string  `json:"-"`
	OutStationName string  `json:"out_station_name,omitempty"`
	OutGateId      string  `json:"-"`
	OutTime        int64   `json:"out_time,omitempty"`
	Status         int8    `json:"status"`
	StatusName     string  `json:"statusName"`
	Money          float32 `json:"money,omitempty"`
	Pay            bool    `json:"pay"`
}

type NotificationVo struct {
	Id   uint64 `json:"id"`
	Type int8   `json:"type"`
}

type TestRouterVo struct {
	UserId      string `json:"user_id"`
	EvidenceKey string `json:"evidence_key"`
	ScanTime    int64  `json:"scan_time"`
}
