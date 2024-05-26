package main

import (
	_ "github.com/balcieren/go-microservice-boilerplate/docs"
	"github.com/balcieren/go-microservice-boilerplate/pkg/infrastructure"
	"go.uber.org/fx"
)

// @title  go-microservice-boilerplate API Documentation
// @version 1.0
// @description This is a boilerplate for a microservice application using Go.
// @host      localhost:8000
// @BasePath  /api
// @schemes http

func main() {
	fx.New(
		infrastructure.CommonModule(),
		infrastructure.SwaggerExternalModule("swagger"),
		fx.Invoke(infrastructure.LaunchHTTPServer),
	).Run()
}
