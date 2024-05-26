package v1

import (
	"github.com/balcieren/go-microservice-boilerplate/pkg/proto"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var Module = fx.Module(
	"owner-grpc-v1",
	fx.Provide(NewService),
	fx.Invoke(func(srv *grpc.Server, s *Service) {
		proto.RegisterOwnerServiceServer(srv, s)
	}),
)
