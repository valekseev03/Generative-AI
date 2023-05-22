package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func getKeyFromFile(fileName string) string {
	code, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	apiKey := string(code)
	apiKey = strings.ReplaceAll(apiKey, "\n", "")
	return apiKey
}

func callChatGPT(key string, message string) string {
	client := openai.NewClient(key)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		return err.Error()
	}

	return resp.Choices[0].Message.Content
}

func main() {
	//Get API Key from ChatGPTAPIKey.txt
	key := getKeyFromFile("ChatGPTAPIKey.txt")
	//fmt.Println("[" + key + "]")

	//Call ChatGPT API
	outputMessage := callChatGPT(key, "Write a paragraph describing how ChatGPT works.")
	fmt.Println(outputMessage)
}
