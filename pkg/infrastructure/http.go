package infrastructure

import (
	"context"
	"net"

	"github.com/go-microservice-boilerplate/pkg/config"
	"github.com/go-microservice-boilerplate/pkg/helper"
	"github.com/go-microservice-boilerplate/pkg/log"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func HTTPModule(name AppName) fx.Option {
	return fx.Module(
		"http",
		fx.Provide(func() *AppName {
			return &name
		}),
		fx.Provide(config.New),
		fx.Provide(log.New),
		fx.Provide(func(c *config.Config) (*grpc.ClientConn, error) {
			conn, err := grpc.Dial(
				net.JoinHostPort(c.GRPC.Host, c.GRPC.Port),
				grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				return nil, err
			}
			return conn, nil
		}),
		fx.Provide(func(c *config.Config, an *AppName) *fiber.App {
			return fiber.New(fiber.Config{
				AppName:      an.String(),
				ErrorHandler: helper.ErrorHandler,
			})
		}),
	)
}

func LaunchHTTPServer(lc fx.Lifecycle, app *fiber.App, c *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Listen(net.JoinHostPort("", c.API.Port))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}
