package infrastructure

import (
	"context"

	"net"

	"github.com/balcieren/go-microservice-boilerplate/pkg/config"
	"github.com/balcieren/go-microservice-boilerplate/pkg/database"
	"github.com/balcieren/go-microservice-boilerplate/pkg/log"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func GRPCModule(name string) fx.Option {
	return fx.Module(
		"grpc",
		fx.Provide(database.NewPostgreSQL),
		fx.Provide(grpc.NewServer),
	)
}

func LaunchGRPCServer(lc fx.Lifecycle, srv *grpc.Server, cmn *config.Common, l *log.Logger) {
	lis, err := net.Listen("tcp", net.JoinHostPort("", cmn.GRPC.PORT))
	if err != nil {
		panic(err)
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				l.Infof("Starting %s server on port: %s ", "UNKNOW", cmn.GRPC.PORT)
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
