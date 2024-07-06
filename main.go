package main

import (
	_ "github.com/balcieren/go-microservice-boilerplate/docs"
	"github.com/balcieren/go-microservice-boilerplate/pkg/infrastructure"
	"go.uber.org/fx"

	ownerApiV1 "github.com/balcieren/go-microservice-boilerplate/app/owner/api/v1"
	petApiV1 "github.com/balcieren/go-microservice-boilerplate/app/pet/api/v1"

	ownerGrpcV1 "github.com/balcieren/go-microservice-boilerplate/app/owner/grpc/v1"
	petGrpcV1 "github.com/balcieren/go-microservice-boilerplate/app/pet/grpc/v1"
)

// @title  go-microservice-boilerplate API Documentation
// @version 1.0
// @description This is a boilerplate for microservice using golang, grpc, and http.
// @host      localhost:8000
// @BasePath  /api
// @schemes http

func main() {
	fx.New(
		infrastructure.CommonModule(),
		infrastructure.HTTPModule("http-module"),
		infrastructure.GRPCModule("grpc-module"),
		infrastructure.SwaggerModule(),
		petApiV1.Module,
		ownerGrpcV1.Module,
		ownerApiV1.Module,
		petGrpcV1.Module,
		fx.Invoke(infrastructure.LaunchGRPCServer, infrastructure.LaunchHTTPServer),
	).Run()
}
