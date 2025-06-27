package giteeai_test

import (
	"testing"

	giteeai "github.com/edmondfrank/go-giteeai/internal"
	"github.com/edmondfrank/go-giteeai/internal/test/checks"
)

func TestJSONMarshaller_Normal(t *testing.T) {
	jm := &giteeai.JSONMarshaller{}
	data := map[string]string{"key": "value"}

	b, err := jm.Marshal(data)
	checks.NoError(t, err)
	if len(b) == 0 {
		t.Fatal("should return non-empty bytes")
	}
}

func TestJSONMarshaller_InvalidInput(t *testing.T) {
	jm := &giteeai.JSONMarshaller{}
	_, err := jm.Marshal(make(chan int))
	checks.HasError(t, err, "should return error for unsupported type")
}

func TestJSONMarshaller_EmptyValue(t *testing.T) {
	jm := &giteeai.JSONMarshaller{}
	b, err := jm.Marshal(nil)
	checks.NoError(t, err)
	if string(b) != "null" {
		t.Fatalf("unexpected marshaled value: %s", string(b))
	}
}
