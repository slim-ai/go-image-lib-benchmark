package main

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

var (
	gcrRoundTripper *http.Transport
)

func init() {
	gcrRoundTripper = http.DefaultTransport.(*http.Transport).Clone()
	gcrRoundTripper.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}
}

// CopyGoContainerregistry copies srcRef to dstRef using go-containerregistry libs.
func CopyGoContainerregistry(ctx context.Context, srcRef, dstRef string, parallelism int) error {
	opts := []crane.Option{
		crane.WithContext(ctx),
		crane.WithTransport(gcrRoundTripper),
		crane.Insecure,
		func(o *crane.Options) {
			o.Remote = append(o.Remote, remote.WithJobs(parallelism))
		},
	}
	return crane.Copy(srcRef, dstRef, opts...)
}
