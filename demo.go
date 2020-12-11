// Package plugindemo a demo plugin.
package traefik-x-request-start

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// Config the plugin configuration.
type Config struct {

}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
	}
}

// Demo a Demo plugin.
type XRequestStart struct {
	next     http.Handler
	name     string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &XRequestStart{
		next:     next,
		name:     name,
	}, nil
}

func (a *XRequestStart) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	
	req.Header.Set("X-Request-Start", time.Now().Unix())

	a.next.ServeHTTP(rw, req)
}
