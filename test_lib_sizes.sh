#!/usr/bin/env bash

export CGO_ENABLED=0

mkdir -p bin
for t in containers-image go-containerregistry containerd; do
    tdir="cmd/${t}"
    cp ${t}.go $tdir
    pushd $tdir >/dev/null
    go mod tidy
    go build -o ../../bin/${t}-copy -tags containers_image_openpgp .
    go build -o ../../bin/${t}-copy-nosymbols -tags containers_image_openpgp -ldflags='-s' .
    popd >/dev/null
    rm -f "${tdir}/${t}.go"
    rm -f "${tdir}/go.sum"
    du -h ./bin/${t}-copy
    du -h ./bin/${t}-copy-nosymbols
done
