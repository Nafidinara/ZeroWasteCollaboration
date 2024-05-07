package infra

type SuccessResponse struct {
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
