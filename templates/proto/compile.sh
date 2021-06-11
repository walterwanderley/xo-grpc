#!/bin/sh
set -u
set -e

Compile () {
    rm -rf $1
    mkdir -p $1
    echo "Compiling $1.proto..."
    protoc -I. -Ivendor --go_out $1 --go_opt paths=source_relative --go-grpc_out $1 --go-grpc_opt paths=source_relative $1.proto
    echo "Generating reverse proxy (grpc-gateway) $1.proto..."
    protoc -I. -Ivendor --grpc-gateway_out $1 --grpc-gateway_opt logtostderr=true,paths=source_relative,allow_repeated_fields_in_body=true,generate_unbound_methods=true $1.proto
    echo "Generating OpenAPIv2 specs $1.proto"
    protoc -I. -Ivendor --openapiv2_out $1 --openapiv2_opt logtostderr=true,allow_repeated_fields_in_body=true,generate_unbound_methods=true $1.proto
}

for i in *.proto; do
    pkg=$(echo "$i" | cut -f 1 -d '.')
    Compile $pkg
done    

echo "Finished!"
