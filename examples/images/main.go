package main

import (
	"context"
	"fmt"
	"os"

	"github.com/edmondfrank/go-giteeai"
)

func main() {
	client := giteeai.NewClient(os.Getenv("OPENAI_API_KEY"))

	respUrl, err := client.CreateImage(
		context.Background(),
		giteeai.ImageRequest{
			Prompt:         "Parrot on a skateboard performs a trick, cartoon style, natural light, high detail",
			Size:           giteeai.CreateImageSize256x256,
			ResponseFormat: giteeai.CreateImageResponseFormatURL,
			N:              1,
		},
	)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return
	}
	fmt.Println(respUrl.Data[0].URL)
}
