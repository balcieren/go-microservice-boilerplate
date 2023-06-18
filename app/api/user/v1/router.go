package v1

import "github.com/gofiber/fiber/v2"

type Router struct {
	handler *Handler
	root    fiber.Router
}

func NewRouter(app *fiber.App, h *Handler) *Router {
	return &Router{
		handler: h,
		root:    app.Group("/api/v1"),
	}
}

func (r *Router) Setup() {
	// r.root.Get("/users", r.handler.ListUsers)
	// r.root.Get("/users/:id", r.handler.GetUser)
	// r.root.Post("/users", r.handler.CreateUser)
	// r.root.Put("/users/:id", r.handler.UpdateUser)
	// r.root.Delete("/users/:id", r.handler.DeleteUser)
}
