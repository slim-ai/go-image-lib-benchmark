package main

import (
	"context"
	"testing"
)

func BenchmarkCopyContainersImage(b *testing.B) {
	benchmarkCopy(b, CopyContainersImage, "docker://")
}

func BenchmarkCopyGoContainerregistry(b *testing.B) {
	benchmarkCopy(b, CopyGoContainerregistry, "")
}

func BenchmarkCopyContainerd(b *testing.B) {
	cf, close, err := NewCopyContainerd()
	checkErrB(b, err)
	b.Cleanup(func() { _ = close() })
	benchmarkCopy(b, cf, "")
}

func benchmarkCopy(b *testing.B, f copyFunc, prefix string) {
	b.Helper()

	b.StopTimer()
	srcRegSrv, dstRegSrv, srcImgs, dstImgs := getRegistries(b, prefix)
	defer srcRegSrv.Close()
	defer dstRegSrv.Close()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ctx := context.Background()
		checkErrB(b, f(ctx, srcImgs[i], dstImgs[i], parallelism))
	}
}
