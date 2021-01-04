package ingress

import "time"

type Request struct {
	Account     string            `json:"account"`
	Category    string            `json:"category"`
	Metadata    map[string]string `json:"metadata"`
	RequestID   string            `json:"request_id"`
	Principal   string            `json:"principal"`
	Service     string            `json:"service"`
	Size        int64             `json:"size"`
	URL         string            `json:"url"`
	ID          string            `json:"id,omitempty"`
	B64Identity string            `json:"b64_identity"`
	Timestamp   time.Time         `json:"timestamp"`
}

type Response struct {
	Request
	Validation string `json:"validation"`
}
