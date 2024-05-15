package entities

import "redoocehub/domains/dto"

type Chatbot struct {
	Message string
}

type ChatbotUsecase interface {
	SendMessage(chatbotConfig ChatbotConfig, content dto.ChatbotRequest) (string, error)
}

type ChatbotConfig struct {
	APIKey string
}
