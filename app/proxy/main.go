package main

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/balcieren/go-microservice/pkg/config"
	"github.com/balcieren/go-microservice/pkg/helper"
	"github.com/balcieren/go-microservice/pkg/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(config.New),
		fx.Provide(log.New),
		fx.Provide(func() *fiber.App {
			return fiber.New(fiber.Config{
				AppName:      "Proxy",
				ErrorHandler: helper.ErrorHandler,
			})
		}),
		fx.Invoke(func(lc fx.Lifecycle, app *fiber.App, cfg *config.Config, l *log.Logger) {
			app.Use(logger.New())

			app.All("/api/v1*", func(c *fiber.Ctx) error {
				var fullPath, parentPath string = "", ""
				if len(strings.Split(string(c.Request().URI().RequestURI()), "/api/v1")) >= 2 {
					fullPath = strings.Split(string(c.Request().URI().RequestURI()), "/api/v1")[1]
				} else {
					return fiber.ErrNotFound
				}

				if len(strings.Split(fullPath, "/")) >= 2 {
					parentPath = strings.Split(fullPath, "/")[1]
				} else {
					return fiber.ErrNotFound
				}

				switch parentPath {
				case "users":
					url := fmt.Sprintf("http://%s:%s/v1%s", cfg.USER_API_HOST_NAME, cfg.API_PORT, fullPath)

					if err := proxy.Do(c, url); err != nil {
						return fiber.NewError(fiber.StatusInternalServerError, err.Error())
					}
				}

				c.Response().Header.Del(fiber.HeaderServer)
				return nil
			})

			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go app.Listen(net.JoinHostPort("", cfg.PROXY_PORT))
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
