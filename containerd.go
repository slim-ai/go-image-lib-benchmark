package main

import (
	"context"
	"fmt"

	"github.com/containerd/containerd"
)

// NewCopyContainerd returns a function that copies srcRef to dstRef using containerd libs.
func NewCopyContainerd() (func(ctx context.Context, srcRef, dstRef string, parallelism int) error, func() error, error) {
	client, err := containerd.New("/run/containerd/containerd.sock", containerd.WithDefaultNamespace("sai-bench"))
	if err != nil {
		return nil, nil, err
	}

	return func(ctx context.Context, srcRef, dstRef string, parallelism int) error {
		pullOpts := []containerd.RemoteOpt{
			containerd.WithMaxConcurrentDownloads(parallelism),
		}
		img, err := client.Pull(ctx, srcRef, pullOpts...)
		if err != nil {
			return fmt.Errorf("pull %s: %v", srcRef, err)
		}

		pushOpts := []containerd.RemoteOpt{
			containerd.WithMaxConcurrentUploadedLayers(parallelism),
		}
		if err = client.Push(ctx, dstRef, img.Target(), pushOpts...); err != nil {
			return fmt.Errorf("push %s: %v", srcRef, err)
		}

		return nil
	}, client.Close, nil
}
