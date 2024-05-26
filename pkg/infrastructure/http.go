package infrastructure

import (
	"context"

	"net"

	"github.com/balcieren/go-microservice-boilerplate/pkg/config"
	"github.com/balcieren/go-microservice-boilerplate/pkg/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func HTTPModule(name string) fx.Option {
	return fx.Module(
		"http",
		fx.Provide(func(c *config.Common) (*grpc.ClientConn, error) {
			return grpc.NewClient(
				net.JoinHostPort(c.GRPC.HOST, c.GRPC.PORT),
				grpc.WithTransportCredentials(insecure.NewCredentials()),
			)
		}),
		fx.Provide(func(env *config.Env) *fiber.App {
			app := fiber.New(fiber.Config{
				AppName:      name,
				ErrorHandler: helper.ErrorHandler,
			})

			app.Use(logger.New())

			return app
		}),
	)
}

func LaunchHTTPServer(lc fx.Lifecycle, app *fiber.App, cmn *config.Common) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Listen(net.JoinHostPort("", cmn.API.PORT))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})
}
