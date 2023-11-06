package main

import (
	v1 "github.com/go-microservice-boilerplate/app/api/user/v1"
	"github.com/go-microservice-boilerplate/pkg/infrastructure"
	"go.uber.org/fx"
)

func main() {
	fx.
		New(
			infrastructure.HTTPModule("user-api"),
			v1.Module,
			fx.Invoke(infrastructure.LaunchHTTPServer),
		).
		Run()
}
