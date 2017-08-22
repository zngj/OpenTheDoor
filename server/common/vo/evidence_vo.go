package vo

type EvidenceVO struct {
	EvidenceId  string `json:"evidence_id"`
	EvidenceKey string `json:"evidence_key"`
	ExpiresAt   int64  `json:"expires_at"`
}

