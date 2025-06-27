package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/edmondfrank/go-giteeai"
)

func main() {
	client := giteeai.NewClient(os.Getenv("GITEEAI_API_KEY"))

	req := giteeai.ChatCompletionRequest{
		Model: giteeai.Qwen2_7B_Instruct,
		Messages: []giteeai.ChatCompletionMessage{
			{
				Role:    giteeai.ChatMessageRoleSystem,
				Content: "you are a helpful chatbot",
			},
		},

	}
	fmt.Println("Conversation")
	fmt.Println("---------------------")
	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		req.Messages = append(req.Messages, giteeai.ChatCompletionMessage{
			Role:    giteeai.ChatMessageRoleUser,
			Content: s.Text(),
		})
		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		fmt.Printf("%s\n\n", resp.Choices[0].Message.Content)
		req.Messages = append(req.Messages, resp.Choices[0].Message)
		fmt.Print("> ")
	}
}
