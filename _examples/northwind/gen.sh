#!/bin/sh
set -u
set -e
set -x

go install github.com/xo/xo@latest

rm -rf internal/application internal/validation internal/server api proto go.mod go.sum main.go registry.go
xo-grpc -m northwind internal/models
