package app

type Response struct {
	Code int
	Body interface{}
}

// Problem defines the Problem JSON type defined by RFC 7807 - media type
// application/problem+json.
// It should be the expected error response for all APIs.
type Problem struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance"`
}
