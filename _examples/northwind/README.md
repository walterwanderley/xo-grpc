# About

Northwind data taken from [Yugabyte database's][yugabyte] Git repository [data samples][yugabyte-git].

[yugabyte]: https://www.yugabyte.com
[yugabyte-git]: https://github.com/yugabyte/yugabyte-db/tree/master/sample
[xo]: https://github.com/xo/xo

## Running

```sh
docker-compose up
```

### Exploring

- gRPC UI [http://localhost:8080/grpcui](http://localhost:8080/grpcui)
- Swagger UI [http://localhost:8080/swagger](http://localhost:8080/swagger)
- Grafana [http://localhost:3000](http://localhost:3000/d/7_VGtoLma/go-grpc1?orgId=1&refresh=10s&from=now-5m&to=now) *user/pass*: **admin/admin**
- Jaeger [http://localhost:16686](http://localhost:16686)
