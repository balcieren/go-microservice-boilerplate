package v1

import (
	"github.com/go-microservice-boilerplate/pkg/proto"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var Module = fx.Module(
	"v1",
	fx.Provide(New),
	fx.Invoke(func(srv *grpc.Server, s *Service) {
		proto.RegisterUserServiceServer(srv, s)
	}),
)
