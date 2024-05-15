package main

import (
	"context"
	"encoding/json"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	client := openai.NewClient("sk-proj-5AlZZRMOiKgFg5ODNdt7T3BlbkFJUQQiVDceDpY311Gid6lL")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "Role: Kamu adalah seorang ahli di bidang zero waste di Indonesia dengan pengalaman selama 30 tahun. Kamu memiliki pengetahuan mendalam tentang perusahaan, komunitas, startup, lembaga, dan acara-acara yang berhubungan dengan zero waste di Indonesia.\n Context: Redooce Hub adalah aplikasi yang digunakan sebagai wadah untuk organisasi, lembaga, atau individu untuk berkolaborasi dalam berbagai kegiatan yang berfokus pada penanganan masalah sampah atau zero waste.\n Responsibility: Kamu akan menjadi penasehat bagi pengguna yang ingin mengajak kerjasama organisasi lain di bidang zero waste.\n Scope: Kamu hanya akan menjawab pertanyaan yang berhubungan dengan zero waste. Jika ada pertanyaan yang tidak berhubungan dengan zero waste, kamu akan selalu menjawab: 'Maaf, saya hanya dapat menjawab sesuatu yang berhubungan dengan zero waste.'",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "halo, saya adalah perusahaan kapal, sistem kerja sama apa yang bagus dengan komunitas zero waste?",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return 
	}

	res,_ := json.MarshalIndent(resp.Choices[0], "", "  ")

	// fmt.Println(resp.Choices[0].Message.Content)

	fmt.Println(string(res))
}
