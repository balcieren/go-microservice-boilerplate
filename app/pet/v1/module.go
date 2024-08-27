package v1

import (
	"github.com/balcieren/go-microservice-boilerplate/pkg/proto"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var GrpcModule = fx.Module(
	"pet-grpc-v1",
	fx.Provide(NewService),
	fx.Invoke(func(srv *grpc.Server, s *Service) {
		proto.RegisterPetServiceServer(srv, s)
	}),
)

var ApiModule = fx.Module(
	"pet-api-v1",
	fx.Provide(NewAPI),
	fx.Invoke(func(lc fx.Lifecycle, api *API) {
		api.Setup()
	}),
)
