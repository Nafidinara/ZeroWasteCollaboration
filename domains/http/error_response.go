package domains

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}