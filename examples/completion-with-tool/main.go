package main

import (
	"context"
	"fmt"
	"os"

	"github.com/edmondfrank/go-giteeai"
	"github.com/edmondfrank/go-giteeai/jsonschema"
)

func main() {
	ctx := context.Background()
	client := giteeai.NewClient(os.Getenv("GITEEAI_API_KEY"))

	// describe the function & its inputs
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"location": {
				Type:        jsonschema.String,
				Description: "The city and state, e.g. San Francisco, CA",
			},
			"unit": {
				Type: jsonschema.String,
				Enum: []string{"celsius", "fahrenheit"},
			},
		},
		Required: []string{"location"},
	}
	f := giteeai.FunctionDefinition{
		Name:        "get_current_weather",
		Description: "Get the current weather in a given location",
		Parameters:  params,
	}
	t := giteeai.Tool{
		Type:     giteeai.ToolTypeFunction,
		Function: &f,
	}

	// simulate user asking a question that requires the function
	dialogue := []giteeai.ChatCompletionMessage{
		{Role: giteeai.ChatMessageRoleUser, Content: "What is the weather in Boston today?"},
	}
	fmt.Printf("Asking GiteeAI '%v' and providing it a '%v()' function...\n",
		dialogue[0].Content, f.Name)
	resp, err := client.CreateChatCompletion(ctx,
		giteeai.ChatCompletionRequest{
			Model:    giteeai.Qwen2_7B_Instruct,
			Messages: dialogue,
			Tools:    []giteeai.Tool{t},
		},
	)
	if err != nil || len(resp.Choices) != 1 {
		fmt.Printf("Completion error: err:%v len(choices):%v\n", err,
			len(resp.Choices))
		return
	}
	msg := resp.Choices[0].Message
	if len(msg.ToolCalls) != 1 {
		fmt.Printf("Completion error: len(toolcalls): %v\n", len(msg.ToolCalls))
		return
	}

	// simulate calling the function & responding to GiteeAI
	dialogue = append(dialogue, msg)
	fmt.Printf("GiteeAI called us back wanting to invoke our function '%v' with params '%v'\n",
		msg.ToolCalls[0].Function.Name, msg.ToolCalls[0].Function.Arguments)
	dialogue = append(dialogue, giteeai.ChatCompletionMessage{
		Role:       giteeai.ChatMessageRoleTool,
		Content:    "Sunny and 80 degrees.",
		Name:       msg.ToolCalls[0].Function.Name,
		ToolCallID: msg.ToolCalls[0].ID,
	})
	fmt.Printf("Sending GiteeAI our '%v()' function's response and requesting the reply to the original question...\n",
		f.Name)
	resp, err = client.CreateChatCompletion(ctx,
		giteeai.ChatCompletionRequest{
			Model:    giteeai.Qwen2_7B_Instruct,
			Messages: dialogue,
			Tools:    []giteeai.Tool{t},
		},
	)
	if err != nil || len(resp.Choices) != 1 {
		fmt.Printf("2nd completion error: err:%v len(choices):%v\n", err,
			len(resp.Choices))
		return
	}

	// display GiteeAI's response to the original question utilizing our function
	msg = resp.Choices[0].Message
	fmt.Printf("GiteeAI answered the original request with: %v\n",
		msg.Content)
}
