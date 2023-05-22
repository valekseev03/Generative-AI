package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func openAPIKeyFile() string {
	code, err := os.ReadFile("ChatGPTAPIKey.txt")
	if err != nil {
		panic(err)
	}

	apiKey := string(code)
	apiKey = strings.ReplaceAll(apiKey, "\n", "")
	return apiKey
}

func main() {
	//Get API Key from ChatGPTAPIKey.txt
	key := openAPIKeyFile()
	//fmt.Println("[" + key + "]")

	//Call ChatGPT API
	client := openai.NewClient(key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
