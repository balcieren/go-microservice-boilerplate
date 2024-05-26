package v1

import "github.com/gofiber/fiber/v2"

type Router struct {
	handler *Handler
	root    fiber.Router
}

func NewRouter(app *fiber.App, h *Handler) *Router {
	return &Router{
		handler: h,
		root:    app.Group("/v1/owners"),
	}
}

func (r *Router) Setup() {
	r.root.Get("/", r.handler.ListOwners)
	r.root.Get("/:id", r.handler.GetOwner)
	r.root.Post("/", r.handler.CreateOwner)
	r.root.Patch("/:id", r.handler.UpdateOwner)
	r.root.Delete("/:id", r.handler.DeleteOwner)
}
