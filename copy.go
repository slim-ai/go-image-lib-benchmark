package main

import (
	"context"
	"crypto/tls"
	"net/http"
)

var (
	roundTripper *http.Transport
)

type copyFunc func(context.Context, string, string, int) error

func init() {
	roundTripper = http.DefaultTransport.(*http.Transport).Clone()
	roundTripper.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
}
