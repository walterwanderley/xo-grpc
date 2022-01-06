## xo-grpc

Create a **gRPC** (and **HTTP/JSON** provided by grpc-gateway reverse proxy) **Server** from the generated code by the [xo project](https://github.com/xo/xo).

### Requirements

- Go 1.16 or superior
- [xo](https://github.com/xo/xo)
- [buf](https://buf.build/)

```sh
go install github.com/xo/xo@latest
go install github.com/bufbuild/buf/cmd/buf@latest
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
