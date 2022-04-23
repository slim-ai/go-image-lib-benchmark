#!/usr/bin/env bash

for t in containerd go-containerregistry containers-image; do
    tdir="cmd/${t}"
    cp ${t}.go $tdir
    go mod tidy
    go build -o ./bin/${t}-copy $tdir
    go build -o -ldflags='-s' ./bin/${t}-copy-nosymbols $tdir
    rm -f ${t}.go
    du -h ./bin/${t}-copy
    du -h ./bin/${t}-copy-nosymbols
done
