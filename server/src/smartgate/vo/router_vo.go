package vo

type RouterStatusVO struct {
	Status int16 `json:"status"`
}

type EvidenceVO struct {
	EvidenceKey string `json:"evidence_key"`
	ExpiresAt   int64 `json:"expires_at"`
}