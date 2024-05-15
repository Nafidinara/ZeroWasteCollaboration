package usecases

import (
	"time"

	"redoocehub/domains/dto"
	"redoocehub/domains/entities"
	"redoocehub/internal/chatbot"
)

type chatbotUsecase struct {
	contextTimeout time.Duration
}

func NewChatbotUsecase(timeout time.Duration) entities.ChatbotUsecase {
	return &chatbotUsecase{
		contextTimeout: timeout,
	}
}

func (c *chatbotUsecase) SendMessage(chatbotConfig entities.ChatbotConfig, content dto.ChatbotRequest) (string, error) {

	message, err := chatbot.SendMessage(chatbotConfig, content.Message)

	if err != nil {
		return "", err
	}

	return message, nil
}
