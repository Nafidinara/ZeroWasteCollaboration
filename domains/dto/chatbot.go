package dto

type ChatbotRequest struct {
	Message string `json:"message" validate:"required"`
}

type ChatbotResponse struct {
	Message string `json:"message"`
	Response string `json:"response"`
}
