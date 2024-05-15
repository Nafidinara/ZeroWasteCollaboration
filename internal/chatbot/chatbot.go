package chatbot

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"

	"redoocehub/domains/entities"
	"redoocehub/internal/constant"
)

func SendMessage(config entities.ChatbotConfig, userContent string) (string, error) {

	client := openai.NewClient(config.APIKey)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: constant.ChatbotSystemContent,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: userContent,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	// res, _ := json.MarshalIndent(resp.Choices[0], "", "  ")
	// fmt.Println(resp.Choices[0].Message.Content)
	// fmt.Println(string(res))

	return resp.Choices[0].Message.Content, nil
}
