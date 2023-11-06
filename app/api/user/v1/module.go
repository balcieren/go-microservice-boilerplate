package v1

import "go.uber.org/fx"

var Module = fx.Module(
	"v1",
	fx.Provide(NewService),
	fx.Provide(NewHandler),
	fx.Provide(NewRouter),
	fx.Invoke(func(rv1 *Router) {
		rv1.Setup()
	}),
)
