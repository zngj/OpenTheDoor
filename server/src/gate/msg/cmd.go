package msg

import (
	"github.com/carsonsx/net4g"
)

const (
	C2S_GATE_LOGIN      = 100
	S2C_GATE_LOGIN      = 101
	S2C_NOT_LOGIN       = 102
	C2S_RSA_KEY         = 103
	S2C_RSA_KEY         = 104
	C2S_VERIFY_EVIDENCE = 200
	S2C_VERIFY_EVIDENCE = 201
	C2S_SUBMIT_EVIDENCE = 202
	S2C_SUBMIT_EVIDENCE = 203
)

var Serializer = NewGateSerializer()

func OnInit() {
	InitSerializer(Serializer)
}

func InitSerializer(s net4g.Serializer) {
	net4g.RegisterId(s, new(S2CGateLogin), S2C_GATE_LOGIN)
	net4g.RegisterId(s, new(S2CRsaKey), S2C_RSA_KEY)
	net4g.RegisterId(s, new(C2SVerifyEvidence), C2S_VERIFY_EVIDENCE)
	net4g.RegisterId(s, new(S2CVerifyEvidence), S2C_VERIFY_EVIDENCE)
	net4g.RegisterId(s, new(C2SSubmitEvidence), C2S_SUBMIT_EVIDENCE)
	net4g.RegisterId(s, new(S2CSubmitEvidence), S2C_SUBMIT_EVIDENCE)
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
	EvidenceId string `json:"evidence_id"`
}

type S2CVerifyEvidence struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

type C2SSubmitEvidence struct {
	EvidenceId string `json:"evidence_id"`
	ScanTime    int64  `json:"scan_time"`
}

type S2CSubmitEvidence struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}
