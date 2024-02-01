#!/usr/bin/env bash

PLATFORM=linux

set -e

ROOT=$(cd "$(dirname "$0")";cd .;pwd)

BIN=${ROOT}/ip-reporter
rm -rf ${BIN}
CGO_ENABLED=0 GOOS=${PLATFORM} GOARCH=amd64 go build -o ${BIN} main.go
chmod +x ${BIN}
