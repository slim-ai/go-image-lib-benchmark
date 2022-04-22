package main

import (
	"context"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

// CopyCrane copies srcRef to dstRef using go-containerregistry libs.
func CopyCrane(ctx context.Context, srcRef, dstRef string, parallelism int) error {
	opts := []crane.Option{
		crane.WithContext(ctx),
		crane.WithTransport(roundTripper),
		crane.Insecure,
		func(o *crane.Options) {
			o.Remote = append(o.Remote, remote.WithJobs(parallelism))
		},
	}
	return crane.Copy(srcRef, dstRef, opts...)
}
