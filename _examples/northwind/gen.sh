#!/bin/sh
set -u
set -e
set -x

go install github.com/xo/xo@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

go install github.com/walterwanderley/xo-grpc@latest

rm -rf internal/application internal/validation internal/server proto go.mod go.sum main.go
xo-grpc -m github.com/walterwanderley/xo-grpc/_examples/northwind internal/models
