# Go GiteeAI

This library provides Go clients for [GiteeAI API](https://ai.gitee.com/). We support: 

* Qwen2
* DeepSeek
* flux-1-schnell

## Installation

```
go get github.com/edmondfrank/go-giteeai
```
Currently, go-giteeai requires Go version 1.18 or greater.


## Usage

### Chat example usage:

```go
package main

import (
	"context"
	"fmt"
	giteeai "github.com/edmondfrank/go-giteeai"
)

func main() {
	client := giteeai.NewClient("your token")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		giteeai.ChatCompletionRequest{
			Model: giteeai.Qwen2_7B_Instruct,
			Messages: []giteeai.ChatCompletionMessage{
				{
					Role:    giteeai.ChatMessageRoleUser,
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

```

### Getting a GiteeAI API Key:

1. Visit the GiteeAI website at [https://ai.gitee.com/](https://ai.gitee.com/).
2. If you don't have an account, click on "Sign Up" to create one. If you do, click "Log In".
3. Once logged in, navigate to your API key management page.
4. Click on "Create new secret key".
5. Enter a name for your new key, then click "Create secret key".
6. Your new API key will be displayed. Use this key to interact with the GiteeAI API.

**Note:** Your API key is sensitive information. Do not share it with anyone.

### Other examples:

<details>
<summary>Chat streaming completion</summary>

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	giteeai "github.com/edmondfrank/go-giteeai"
)

func main() {
	c := giteeai.NewClient("your token")
	ctx := context.Background()

	req := giteeai.ChatCompletionRequest{
		Model:     giteeai.Qwen2_7B_Instruct,
		MaxTokens: 20,
		Messages: []giteeai.ChatCompletionMessage{
			{
				Role:    giteeai.ChatMessageRoleUser,
				Content: "Lorem ipsum",
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	fmt.Printf("Stream response: ")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("\nStream finished")
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}


		fmt.Printf(response.Choices[0].Delta.Content)
	}
}
```
</details>

<details>
<summary>Image generation</summary>

```go
package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	giteeai "github.com/edmondfrank/go-giteeai"
	"image/png"
	"os"
)

func main() {
	c := giteeai.NewClient("your token")
	ctx := context.Background()

	// Sample image by link
	reqUrl := giteeai.ImageRequest{
		Prompt:         "Parrot on a skateboard performs a trick, cartoon style, natural light, high detail",
		Size:           giteeai.CreateImageSize256x256,
		ResponseFormat: giteeai.CreateImageResponseFormatURL,
		N:              1,
	}

	respUrl, err := c.CreateImage(ctx, reqUrl)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return
	}
	fmt.Println(respUrl.Data[0].URL)

	// Example image as base64
	reqBase64 := giteeai.ImageRequest{
		Prompt:         "Portrait of a humanoid parrot in a classic costume, high detail, realistic light, unreal engine",
		Size:           giteeai.CreateImageSize256x256,
		ResponseFormat: giteeai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	respBase64, err := c.CreateImage(ctx, reqBase64)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return
	}

	imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	if err != nil {
		fmt.Printf("Base64 decode error: %v\n", err)
		return
	}

	r := bytes.NewReader(imgBytes)
	imgData, err := png.Decode(r)
	if err != nil {
		fmt.Printf("PNG decode error: %v\n", err)
		return
	}

	file, err := os.Create("example.png")
	if err != nil {
		fmt.Printf("File creation error: %v\n", err)
		return
	}
	defer file.Close()

	if err := png.Encode(file, imgData); err != nil {
		fmt.Printf("PNG encode error: %v\n", err)
		return
	}

	fmt.Println("The image was saved as example.png")
}

```
</details>

<details>
<summary>Configuring proxy</summary>

```go
config := giteeai.DefaultConfig("token")
proxyUrl, err := url.Parse("http://localhost:{port}")
if err != nil {
	panic(err)
}
transport := &http.Transport{
	Proxy: http.ProxyURL(proxyUrl),
}
config.HTTPClient = &http.Client{
	Transport: transport,
}

c := giteeai.NewClientWithConfig(config)
```

See also: https://pkg.go.dev/badge/github.com/edmondfrank/go-giteeai.svg
</details>

<details>
<summary>Error handling</summary>

example:
```
e := &giteeai.APIError{}
if errors.As(err, &e) {
  switch e.HTTPStatusCode {
    case 401:
      // invalid auth or key (do not retry)
    case 429:
      // rate limiting or engine overload (wait and retry) 
    case 500:
      // giteeai server error (retry)
    default:
      // unhandled
  }
}

```
</details>

See the `examples/` folder for more.

## Contributing

By following [Contributing Guidelines](https://github.com/edmondfrank/go-giteeai/blob/master/CONTRIBUTING.md), we hope to ensure that your contributions are made smoothly and efficiently.
