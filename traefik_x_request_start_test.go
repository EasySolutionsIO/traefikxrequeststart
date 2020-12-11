package traefik_x_request_start_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EasySolutionsIO/traefik_x_request_start"
)

func TestXRequestStart(t *testing.T) {
	cfg := traefik_x_request_start.CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := traefik_x_request_start.New(ctx, next, cfg, "traefik-x-request-plugin")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

}

func assertHeader(t *testing.T, req *http.Request, key, expected string) {
	t.Helper()

	if req.Header.Get(key) != expected {
		t.Errorf("invalid header value: %s", req.Header.Get(key))
	}
}
