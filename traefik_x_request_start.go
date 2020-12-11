// Package traefikxrequeststart a plugin for traefik which adds X-Request-Start header.
package traefikxrequeststart

import (
	"context"
	"net/http"
	"time"
)

// Config the plugin configuration.
type Config struct {
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// XRequestStart a traefik plugin.
type XRequestStart struct {
	next http.Handler
	name string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &XRequestStart{
		next: next,
		name: name,
	}, nil
}

func (a *XRequestStart) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	req.Header.Set("X-Request-Start", string(time.Now().Unix()))

	a.next.ServeHTTP(rw, req)
}
