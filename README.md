# Go Image Library Benchmarking

This repository contains a test harness and wrapper functions for benchmarking Go image library functionality.

The current list of libraries:
- [`containers/image`](https://github.com/containers/image), called "skopeo" here
 since [`skopeo`](https://github.com/containers/skopeo) is the most well-known entrypoint for this library.
- [`google/go-containerregistry`](https://github.com/google/go-containerregistry),
specifically the [`crane` package](https://github.com/google/go-containerregistry/tree/main/pkg/crane).
- [`containerd/containerd`](https://github.com/containerd/containerd), an industry standard container runtime lib.

## Benchmarks

You can run benchmarks on your machine by running `make`.

Hardware info:

```
goos: linux
goarch: amd64
pkg: github.com/slim-ai/go-image-lib-benchmark
cpu: 11th Gen Intel(R) Core(TM) i7-1195G7 @ 2.90GHz
```

Each test is run with max parallelism of 8. Libraries differ in their usage of parallelism;
in most cases it equates to number of goroutine workers.

Current stats are for 1, 4, and 8 CPUs.

### Copy

This benchmark tests copying from one networked registry to another.

Notes and caveats:
- The `containerd` library does not offer a copy method, and also does not purport itself to be tuned for such a workload,
so it does not fit well in this benchmark. Nonetheless it is worth analysing for breadth.

```
BenchmarkCopySkopeo                  100          14055206 ns/op         1045418 B/op       6487 allocs/op
BenchmarkCopySkopeo-4                100          10218015 ns/op         1099314 B/op       6615 allocs/op
BenchmarkCopySkopeo-8                100          10393303 ns/op         1241455 B/op       6673 allocs/op
BenchmarkCopyCrane                   100           2160080 ns/op          413823 B/op       2987 allocs/op
BenchmarkCopyCrane-4                 100           1356623 ns/op          407963 B/op       2936 allocs/op
BenchmarkCopyCrane-8                 100           1477929 ns/op          476178 B/op       2960 allocs/op
BenchmarkCopyContainerd              100          24969750 ns/op          605232 B/op       7355 allocs/op
BenchmarkCopyContainerd-4            100          24587615 ns/op          753558 B/op       7449 allocs/op
BenchmarkCopyContainerd-8            100          24783194 ns/op          990529 B/op       7448 allocs/op
```
