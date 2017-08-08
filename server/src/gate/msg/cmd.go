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
	C2S_USER_EVIDENCE   = 202
	S2C_USER_EVIDENCE   = 203
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
	net4g.RegisterId(s, new(C2SUserEvidence), C2S_USER_EVIDENCE)
	net4g.RegisterId(s, new(S2CUserEvidence), S2C_USER_EVIDENCE)
}

type S2CGateLogin struct {
	Code          int   `json:"code"` // 0-登录成功;1-GateId不存在
	GateDirection int8 `json:"gate_direction,omitempty"`
	StationName   string `json:"station_name,omitempty"`
	CityName      string `json:"city_name,omitempty"`
}

type S2CRsaKey struct {
	Key string `json:"key"`
}

type C2SVerifyEvidence struct {
	EvidenceKey string `json:"evidence_key"`
}

type S2CVerifyEvidence struct {
	Code int `json:"code"` //0-通过;1-凭证不存在;2-凭证已过期;3-凭证与机闸不匹配;4-用户不符合付费标准
}

type C2SUserEvidence struct {
	EvidenceKey string `json:"evidence_key"`
	ScanTime    int64    `json:"scan_time"`
}

type S2CUserEvidence struct {
	Success bool `json:"success"`
}
