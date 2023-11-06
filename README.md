
# Go Microservice Boilerplate

The pattern which used with dependency injection, makes it easy to build expandable applications.

![microservice](https://i.ibb.co/rxRpDnL/microservice.png)

## Packages

- [Fiber](https://github.com/gofiber/fiber)
- [Fiber Swagger](https://github.com/gofiber/swagger)
- [Uber FX](https://github.com/uber-go/fx)
- [Uber Zap](https://github.com/uber-go/zap)
- [Ent](https://github.com/ent/ent)
- [gRPC](https://github.com/grpc/grpc-go)

## Databases

- [PostgreSQL](https://www.postgresql.org.pl/)
- [Meilisearch](https://www.meilisearch.com/)
- [Redis](https://redis.io/)

## Storages

- [MinIO](https://min.io/docs/minio/linux/developers/go/minio-go.html)

## Worker

- [asynq](https://github.com/hibiken/asynq)

## Deployment

### Development Mode (hot reload)

```bash
make dev
```

### Production Mode

```bash
make build
```

## Additional Makefile Commands

Generate all proto files

```bash
make proto
```

Generate swagger

```bash
make swagger
```

Generate ent

```bash
make ent
```

## Swagger Path

localhost:8000/api/swagger

## Missing Features

- Load balancer in proxy