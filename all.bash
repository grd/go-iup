#!/bin/sh

set -eux

export CGO_CFLAGS="-I/usr/include/iup -I/usr/include/cd -I/usr/include/lm"
go install github.com/grd/iup... "$@"
