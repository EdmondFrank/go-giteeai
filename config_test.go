package giteeai_test

import (
	"testing"

	"github.com/edmondfrank/go-giteeai"
)

func TestClientConfigString(t *testing.T) {
	// String() should always return the constant value
	conf := giteeai.DefaultConfig("dummy-token")
	expected := "<GiteeAI API ClientConfig>"
	got := conf.String()
	if got != expected {
		t.Errorf("ClientConfig.String() = %q; want %q", got, expected)
	}
}
