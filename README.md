## xo-grpc

Create a **gRPC** (and **HTTP/JSON** provided by grpc-gateway reverse proxy) **Server** from the generated code by the [xo project](https://github.com/xo/xo).

### Requirements

- Go 1.16 or superior
- [protoc](https://github.com/protocolbuffers/protobuf/releases)
- xo, protoc-gen-go and protoc-gen-go-grpc

```sh
go install github.com/xo/xo@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
```

### Installation

```sh
go install github.com/walterwanderley/xo-grpc@latest
```

### Example

1. Generate go code to access your database using xo. For more informations about `xo command line parameters`, please visit [the xo Documentation](https://github.com/xo/xo)

```sh
mkdir models
xo schema -o models [Database Connection URL] 
```

2. Execute xo-grpc to generate the gRPC Server boilerplate code:

```sh
xo-grpc models
```

3. Run the generated server:

```sh
go run . -db [Database Connection URL] -dev -grpcui
```

4. Enjoy!

- gRPC UI [http://localhost:5000/grpcui](http://localhost:5000/grpcui)
- Swagger UI [http://localhost:5000/swagger](http://localhost:5000/swagger)

### Similar Projects

- [sqlc-grpc](https://github.com/walterwanderley/sqlc-grpc)
