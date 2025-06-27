package giteeai

import (
	"net/http"
)

const (
	giteeaiAPIURLv1           = "https://ai.gitee.com/v1"
	defaultEmptyMessagesLimit = 300
)

const defaultAssistantVersion = "v2" 

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	authToken string

	BaseURL          string
	AssistantVersion string
	HTTPClient       HTTPDoer

	EmptyMessagesLimit uint
}

func DefaultConfig(authToken string) ClientConfig {
	return ClientConfig{
		authToken:        authToken,
		BaseURL:          giteeaiAPIURLv1,
		AssistantVersion: defaultAssistantVersion,

		HTTPClient: &http.Client{},

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}

func (ClientConfig) String() string {
	return "<GiteeAI API ClientConfig>"
}
