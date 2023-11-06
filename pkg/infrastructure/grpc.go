package infrastructure

import (
	"context"
	"net"

	"github.com/go-microservice-boilerplate/pkg/config"
	"github.com/go-microservice-boilerplate/pkg/database"
	"github.com/go-microservice-boilerplate/pkg/log"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func GRPCModule(name AppName) fx.Option {
	return fx.Module(
		"grpc",
		fx.Provide(func() *AppName {
			return &name
		}),
		fx.Provide(config.New),
		fx.Provide(config.NewCommonConfig),
		fx.Provide(log.New),
		fx.Provide(grpc.NewServer),
		fx.Provide(database.NewPostgreSQL),
	)
}

func LaunchGRPCServer(lc fx.Lifecycle, srv *grpc.Server, c *config.Config, cmn *config.CommonConfig, l *log.Logger, an *AppName) {
	lis, err := net.Listen("tcp", net.JoinHostPort("", c.GRPC.Port))
	if err != nil {
		panic(err)
	}

	lc.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				l.Infof("Starting %s server on port: %s ", an.String(), c.GRPC.Port)
				go srv.Serve(lis)
				return nil
			},
			OnStop: func(ctx context.Context) error {
				srv.GracefulStop()
				return nil
			},
		},
	)

}
