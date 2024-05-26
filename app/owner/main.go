package main

import (
	"github.com/balcieren/go-microservice-boilerplate/pkg/infrastructure"
	"go.uber.org/fx"

	ownerApiV1 "github.com/balcieren/go-microservice-boilerplate/app/owner/api/v1"
	ownerGrpcV1 "github.com/balcieren/go-microservice-boilerplate/app/owner/grpc/v1"
)

func main() {
	fx.New(
		infrastructure.CommonModule(),
		infrastructure.HTTPModule("owner-api"),
		infrastructure.GRPCModule("owner-grpc"),
		ownerApiV1.Module,
		ownerGrpcV1.Module,
		fx.Invoke(infrastructure.LaunchGRPCServer, infrastructure.LaunchHTTPServer),
	).Run()
}
