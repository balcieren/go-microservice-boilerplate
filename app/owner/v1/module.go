package v1

import (
	"github.com/balcieren/go-microservice-boilerplate/pkg/proto"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var GrpcModule = fx.Module(
	"owner-grpc-v1",
	fx.Provide(NewService),
	fx.Invoke(func(srv *grpc.Server, s *Service) {
		proto.RegisterOwnerServiceServer(srv, s)
	}),
)

var ApiModule = fx.Module(
	"owner-api-v1",
	fx.Provide(NewAPI),
	fx.Invoke(func(lc fx.Lifecycle, api *API) {
		api.Setup()
	}),
)
