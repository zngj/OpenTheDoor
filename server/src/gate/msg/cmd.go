package msg

const (
	GATE_LOGIN      = 100
	NOT_LOGIN       = 101
	RSA_KEY         = 102
	VERIFY_EVIDENCE = 103
	SUBMIT_EVIDENCE = 104
)

var Serializer = NewGateSerializer()

func OnInit() {
	Serializer.SerializeId(new(S2CGateLogin), GATE_LOGIN)
	Serializer.SerializeId(new(S2CRsaKey), RSA_KEY)
	Serializer.DeserializeId(new(C2SVerifyEvidence), VERIFY_EVIDENCE)
	Serializer.SerializeId(new(S2CVerifyEvidence), VERIFY_EVIDENCE)
	Serializer.DeserializeId(new(C2SSubmitEvidence), SUBMIT_EVIDENCE)
	Serializer.SerializeId(new(S2CSubmitEvidence), SUBMIT_EVIDENCE)
}

type S2CGateLogin struct {
	GateId        string `json:"gate_id,omitempty"`
	GateDirection int8   `json:"gate_direction,omitempty"`
	StationName   string `json:"station_name,omitempty"`
	CityName      string `json:"city_name,omitempty"`
	ErrCode       int    `json:"errcode,omitempty"` //3100
	ErrMsg        string `json:"errmsg,omitempty"`
}

type S2CRsaKey struct {
	Key     string `json:"key"`
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

type C2SVerifyEvidence struct {
	EvidenceKey string `json:"evidence_key"`
}

type S2CVerifyEvidence struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

type C2SSubmitEvidence struct {
	EvidenceKey string `json:"evidence_key"`
	ScanTime    int64  `json:"scan_time"`
}

type S2CSubmitEvidence struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}
