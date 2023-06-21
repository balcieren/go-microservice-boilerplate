package main

import (
	"context"
	"net"

	v1 "github.com/balcieren/go-microservice/app/api/user/v1"
	"github.com/balcieren/go-microservice/pkg/config"
	"github.com/balcieren/go-microservice/pkg/helper"
	"github.com/balcieren/go-microservice/pkg/log"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func main() {
	fx.
		New(
			fx.Provide(config.New),
			fx.Provide(log.New),
			fx.Provide(func(c *config.Config) *fiber.App {
				return fiber.New(fiber.Config{
					AppName:      "User API",
					ErrorHandler: helper.ErrorHandler,
				})
			}),
			fx.Module("v1", v1.Module),
			fx.Invoke(func(lc fx.Lifecycle, app *fiber.App, c *config.Config, rv1 *v1.Router) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						rv1.Setup()
						go app.Listen(net.JoinHostPort(c.USER_API_HOST_NAME, c.API_PORT))
						return nil
					},
					OnStop: func(ctx context.Context) error {
						return app.Shutdown()
					},
				})
			}),
		).
		Run()
}
