package main

import (
	v1 "github.com/go-microservice-boilerplate/app/grpc/user/v1"
	"github.com/go-microservice-boilerplate/pkg/infrastructure"
	"go.uber.org/fx"
)

func main() {
	fx.
		New(
			infrastructure.GRPCModule("user-grpc"),
			v1.Module,
			fx.Invoke(infrastructure.LaunchGRPCServer),
		).
		Run()
}
