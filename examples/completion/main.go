package main

import (
	"context"
	"fmt"
	"os"

	"github.com/edmondfrank/go-giteeai"
)

func main() {
	client := giteeai.NewClient(os.Getenv("GITEEAI_API_KEY"))
	resp, err := client.CreateCompletion(
		context.Background(),
		giteeai.CompletionRequest{
			Model:     giteeai.Qwen2_7B_Instruct,
			MaxTokens: 5,
			Prompt:    "Lorem ipsum",
		},
	)
	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}
	fmt.Println(resp.Choices[0].Text)
}
