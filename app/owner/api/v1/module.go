package v1

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"owner-api-v1",
	fx.Provide(NewHandler),
	fx.Provide(NewRouter),
	fx.Invoke(func(lc fx.Lifecycle, r *Router) {
		r.Setup()
	}),
)
