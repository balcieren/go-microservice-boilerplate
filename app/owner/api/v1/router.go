package v1

import (
	"github.com/balcieren/go-microservice-boilerplate/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	handler *Handler
	root    fiber.Router
}

func NewRouter(app *fiber.App, h *Handler, env *config.Env) *Router {
	var path string = "/v1/owners"
	if env.AppEnv == "dev" {
		path = "/api/v1/owners"
	}

	return &Router{
		handler: h,
		root:    app.Group(path),
	}
}

func (r *Router) Setup() {
	r.root.Get("/", r.handler.ListOwners)
	r.root.Get("/:id", r.handler.GetOwner)
	r.root.Post("/", r.handler.CreateOwner)
	r.root.Patch("/:id", r.handler.UpdateOwner)
	r.root.Delete("/:id", r.handler.DeleteOwner)
}
