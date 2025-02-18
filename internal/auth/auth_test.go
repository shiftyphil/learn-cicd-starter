package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	want := "THISISANAPIKEY"

	header := http.Header{}
	header.Set("Authorization", fmt.Sprintf("ApiKey %s", want))

	got, _ := GetAPIKey(header)
	if want != got {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
