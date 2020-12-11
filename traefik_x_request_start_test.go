package traefikxrequeststart_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EasySolutionsIO/traefikxrequeststart"
)

func TestXRequestStart(t *testing.T) {
	cfg := traefikxrequeststart.CreateConfig()

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := traefikxrequeststart.New(ctx, next, cfg, "traefik-x-request-plugin")
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
