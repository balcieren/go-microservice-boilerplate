package infrastructure

import (
	docs "github.com/balcieren/go-microservice-boilerplate/docs"
	"github.com/balcieren/go-microservice-boilerplate/pkg/config"
	"github.com/balcieren/go-microservice-boilerplate/pkg/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"go.uber.org/fx"
)

func SwaggerExternalModule(name string) fx.Option {
	return fx.Module(
		"swagger",
		fx.Provide(func(env *config.Env) *fiber.App {
			docs.SwaggerInfo.Host = env.Swagger.Host
			docs.SwaggerInfo.BasePath = env.Swagger.BasePath

			app := fiber.New(fiber.Config{
				AppName:      name,
				ErrorHandler: helper.ErrorHandler,
			})

			app.Use(logger.New())
			app.Use(cors.New(
				cors.Config{
					AllowOrigins: "*",
				},
			))

			app.Get("/*", basicauth.New(basicauth.Config{
				Authorizer: func(user, pass string) bool {
					return user == env.Swagger.User && pass == env.Swagger.Password
				},
			}), swagger.HandlerDefault)

			return app
		}),
	)
}
