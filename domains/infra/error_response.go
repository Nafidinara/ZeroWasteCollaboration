package infra

type ErrorResponse struct {
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
