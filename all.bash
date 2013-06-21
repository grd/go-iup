#!/bin/sh

set -eux

export CGO_CFLAGS="-I/usr/include/iup -I/usr/include/cd -I/usr/include/lm"
go get github.com/grd/iup... "$@"
