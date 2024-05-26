package infrastructure

import (
	"github.com/balcieren/go-microservice-boilerplate/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/swagger"
	"go.uber.org/fx"
)

func SwaggerModule() fx.Option {
	return fx.Module(
		"swagger",
		fx.Invoke(func(app *fiber.App, env *config.Env) {
			app.Get("/api/swagger/*", basicauth.New(basicauth.Config{
				Authorizer: func(user, pass string) bool {
					return user == env.Swagger.User && pass == env.Swagger.Password
				},
			}), swagger.HandlerDefault)
		}),
	)
}
