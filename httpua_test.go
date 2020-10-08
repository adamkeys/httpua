package httpua_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	ua "github.com/adamkeys/httpua"
)

const userAgent = "Foo/1.0"

func TestNewClient(t *testing.T) {
	client := ua.NewClient(userAgent)
	assertUserAgent(t, client, userAgent)
}

func TestWithClient(t *testing.T) {
	client := ua.WithClient(http.DefaultClient, userAgent)
	assertUserAgent(t, client, userAgent)
}

func assertUserAgent(t *testing.T, client *http.Client, userAgent string) {
	t.Helper()

	var capturedUserAgent string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedUserAgent = r.Header.Get("User-Agent")
	}))
	defer srv.Close()

	resp, err := client.Get(srv.URL)
	if err != nil {
		t.Fatalf("failed to perform request: %v", err)
	}
	resp.Body.Close()

	if capturedUserAgent != userAgent {
		t.Fatalf("expected user agent: %q; got: %q", userAgent, capturedUserAgent)
	}
}
