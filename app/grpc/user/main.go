package main

import (
	"context"
	"net"

	"github.com/balcieren/go-microservice/app/grpc/user/database"
	"github.com/balcieren/go-microservice/app/grpc/user/service"
	"github.com/balcieren/go-microservice/pkg/config"
	"github.com/balcieren/go-microservice/pkg/log"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func main() {
	fx.
		New(
			fx.Provide(config.New),
			fx.Provide(log.New),
			fx.Provide(grpc.NewServer),
			fx.Provide(database.NewPostgreSQL),
			fx.Provide(service.New),
			fx.Invoke(func(lc fx.Lifecycle, srv *grpc.Server, c *config.Config, l *log.Logger) {
				lis, err := net.Listen("tcp", net.JoinHostPort(c.USER_GRPC_HOST_NAME, c.GRPC_PORT))
				if err != nil {
					panic(err)
				}

				lc.Append(
					fx.Hook{
						OnStart: func(context.Context) error {
							l.Info("Starting GRPC server on port", c.GRPC_PORT)
							return srv.Serve(lis)
						},
						OnStop: func(ctx context.Context) error {
							srv.GracefulStop()
							return nil
						},
					},
				)

			}),
		).
		Run()
}
