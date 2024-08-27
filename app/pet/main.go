package main

import (
	"github.com/balcieren/go-microservice-boilerplate/pkg/infrastructure"
	"go.uber.org/fx"

	petV1 "github.com/balcieren/go-microservice-boilerplate/app/pet/v1"
)

func main() {
	fx.New(
		infrastructure.CommonModule(),
		infrastructure.HTTPModule("pet-api"),
		infrastructure.GRPCModule("pet-grpc"),
		petV1.GrpcModule,
		petV1.ApiModule,
		fx.Invoke(infrastructure.LaunchGRPCServer, infrastructure.LaunchHTTPServer),
	).Run()
}
