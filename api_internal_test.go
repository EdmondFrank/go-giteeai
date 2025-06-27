package giteeai

import (
	"context"
	"testing"
)

func TestGiteeAIFullURL(t *testing.T) {
	cases := []struct {
		Name   string
		Suffix string
		Expect string
	}{
		{
			"ChatCompletionsURL",
			"/chat/completions",
			"https://ai.gitee.com/v1/chat/completions",
		},
		{
			"CompletionsURL",
			"/completions",
			"https://ai.gitee.com/v1/completions",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			az := DefaultConfig("dummy")
			cli := NewClientWithConfig(az)
			actual := cli.fullURL(c.Suffix)
			if actual != c.Expect {
				t.Errorf("Expected %s, got %s", c.Expect, actual)
			}
			t.Logf("Full URL: %s", actual)
		})
	}
}

func TestRequestAuthHeader(t *testing.T) {
	cases := []struct {
		Name      string
		HeaderKey string
		Token     string
		Expect    string
	}{
		{
			"GiteeAIDefault",
			"Authorization",
			"dummy-token-giteeai",
			"Bearer dummy-token-giteeai",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			az := DefaultConfig(c.Token)

			cli := NewClientWithConfig(az)
			req, err := cli.newRequest(context.Background(), "POST", "/chat/completions")
			if err != nil {
				t.Errorf("Failed to create request: %v", err)
			}
			actual := req.Header.Get(c.HeaderKey)
			if actual != c.Expect {
				t.Errorf("Expected %s, got %s", c.Expect, actual)
			}
			t.Logf("%s: %s", c.HeaderKey, actual)
		})
	}
}
