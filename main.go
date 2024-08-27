package main

import (
	_ "github.com/balcieren/go-microservice-boilerplate/docs"
	"github.com/balcieren/go-microservice-boilerplate/pkg/infrastructure"
	"go.uber.org/fx"

	ownerV1 "github.com/balcieren/go-microservice-boilerplate/app/owner/v1"
	petV1 "github.com/balcieren/go-microservice-boilerplate/app/pet/v1"
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
		ownerV1.GrpcModule,
		ownerV1.ApiModule,
		petV1.GrpcModule,
		petV1.ApiModule,
		fx.Invoke(infrastructure.LaunchGRPCServer, infrastructure.LaunchHTTPServer),
	).Run()
}
