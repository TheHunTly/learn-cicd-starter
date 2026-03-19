package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	expectedApiKey := "test123"
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey "+expectedApiKey)

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatal(err.Error())
	}

	if apiKey != expectedApiKey {
		t.Fatalf("Expected apiKey != apiKey: '%s' != '%s'\n",
			expectedApiKey, apiKey)
	}
}
