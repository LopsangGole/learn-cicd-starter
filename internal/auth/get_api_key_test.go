package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey some-secret-key")

	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	want := "some-secret-key"
	// if got != want {
		t.Errorf("expected %v, got %v", want, got)
	}
}