package main

import (
	"github.com/balcieren/go-microservice-boilerplate/pkg/infrastructure"
	"go.uber.org/fx"

	ownerV1 "github.com/balcieren/go-microservice-boilerplate/app/owner/v1"
)

func main() {
	fx.New(
		infrastructure.CommonModule(),
		infrastructure.HTTPModule("owner-api"),
		infrastructure.GRPCModule("owner-grpc"),
		ownerV1.ApiModule,
		ownerV1.GrpcModule,
		fx.Invoke(infrastructure.LaunchGRPCServer, infrastructure.LaunchHTTPServer),
	).Run()
}
