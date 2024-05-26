package main

import (
	"github.com/balcieren/go-microservice-boilerplate/pkg/infrastructure"
	"go.uber.org/fx"

	petApiV1 "github.com/balcieren/go-microservice-boilerplate/app/pet/api/v1"
	petGrpcV1 "github.com/balcieren/go-microservice-boilerplate/app/pet/grpc/v1"
)

func main() {
	fx.New(
		infrastructure.CommonModule(),
		infrastructure.HTTPModule("pet-api"),
		infrastructure.GRPCModule("pet-grpc"),
		petApiV1.Module,
		petGrpcV1.Module,
		fx.Invoke(infrastructure.LaunchGRPCServer, infrastructure.LaunchHTTPServer),
	).Run()
}
