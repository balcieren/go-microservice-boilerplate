package main

import (
	"context"
	"fmt"
	"net"

	_ "github.com/go-microservice-boilerplate/app/proxy/docs"
	"github.com/go-microservice-boilerplate/pkg/config"
	"github.com/go-microservice-boilerplate/pkg/helper"
	"github.com/go-microservice-boilerplate/pkg/log"
	"github.com/go-microservice-boilerplate/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/swagger"
	"go.uber.org/fx"
)

// @title  go-microservice-boilerplate API Documentation
// @version 1.0
// @description This is a sample server for go-microservice-boilerplate API Documentation.

// @host      localhost:8000
// @BasePath  /api
// @schemes http

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

			app.Get("/api/swagger/*", swagger.HandlerDefault)

			app.All("/api/v1*", func(c *fiber.Ctx) error {
				path, err := utils.SplitProxyPath(c, "/api/v1")
				if err != nil {
					return fiber.ErrNotFound
				}

				if err := proxy.Do(c, utils.GenerateURL(path, utils.OutgoingPath{
					Version:   "v1",
					HostSplit: "-api",
					Port:      cfg.API.Port,
				})); err != nil {
					return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Cannot %s %s", c.Method(), c.Path()))
				}

				c.Response().Header.Del(fiber.HeaderServer)
				return nil
			})

			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go app.Listen(net.JoinHostPort("", cfg.Proxy.Port))
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
