#!/bin/sh
set -u
set -e
set -x

go install github.com/xo/xo@latest

rm -rf internal api proto tools go.mod go.sum main.go registry.go buf*
mkdir -p internal/models

xo schema sq:northwind.db -o internal/models
# patching xo's bug
sed -i 's/sql.NullInt64(id)/sql.NullInt64{Valid: true, Int64: id}/g' internal/models/*

xo-grpc -m northwind -db-driver sqlite3 -db-module github.com/mattn/go-sqlite3 internal/models


