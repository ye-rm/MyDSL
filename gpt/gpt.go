// Package gpt provides a wrapper for OpenAI GPT-3 API
// You are supposed to set OPENAI_FOR_DSL to enable openai
// Set OPENAI_API_KEY to your API key if your wanner enable openai
package gpt

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"os"
)

const prompt = "You are a customer assistant Bot, answer shorter than 60 chars, query: "

var OPENAI_API_KEY = os.Getenv("OPENAI_API_KEY")

// Init check OPENAI_FOR_DSL environment variable
// if OPENAI_FOR_DSL is not set, set it to false
// if OPENAI_FOR_DSL is set to true, check OPENAI_API_KEY
// if OPENAI_API_KEY is not set, set OPENAI_FOR_DSL to false
func Init() {
	env := os.Getenv("OPENAI_FOR_DSL")
	switch env {
	case "true":
		println("OPENAI_FOR_DSL is set to true")
		if OPENAI_API_KEY == "" {
			fmt.Println("OPENAI_API_KEY environment variable is not set, automatically set OPENAI_FOR_DSL to false")
			os.Setenv("OPENAI_FOR_DSL", "false")
		}
	case "false":
		println("OPENAI_FOR_DSL is set to false")
	default:
		os.Setenv("OPENAI_FOR_DSL", "false")
		println("OPENAI_FOR_DSL is not set, default: false")
	}
}

// GPT use openai to generate response, accept query string, return response string
func GPT(query string) string {
	if OPENAI_API_KEY == "" {
		fmt.Println("OPENAI_API_KEY environment variable is not set")
		os.Exit(1)
	}
	client := openai.NewClient(OPENAI_API_KEY)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt + query,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}

	return resp.Choices[0].Message.Content
}
