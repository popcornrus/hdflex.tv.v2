package middleware

import (
	"go.uber.org/fx"
)

type Middleware struct {
}

func NewMiddlewares() *Middleware {
	return &Middleware{}
}

func NewMiddleware() fx.Option {
	return fx.Module(
		"middleware",
		fx.Provide(
			NewMiddlewares,
		),
	)
}
