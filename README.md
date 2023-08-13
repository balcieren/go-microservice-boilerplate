
# Go Microservice Pattern

The pattern which used with dependency injection, makes it easy to build expandable applications.

![microservice](https://i.ibb.co/rxRpDnL/microservice.png)


## Packages

- [Fiber](https://github.com/gofiber/fiber)
- [Fiber Swagger](https://github.com/gofiber/swagger)
- [Uber FX](https://github.com/uber-go/fx)
- [Uber Zap](https://github.com/uber-go/zap)
- [Ent](https://github.com/ent/ent)
- [gRPC](https://github.com/grpc/grpc-go)


## Deployment

#### Development Mode  (hot reload)
```bash
make dev
```
#### Production Mode
```bash
make build
```    
## Additional Makefile Commands

Generate all proto files

```bash
make proto
```

Genere swagger

```bash
make swagger
```

## Swagger Path

localhost:8000/api/swagger


## Missing Features

- Load balancer in proxy

