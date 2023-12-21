package middleware

import (
	"go.uber.org/fx"
)

type Middleware struct {
	AuthMiddleware *AuthMiddleware
}

func NewMiddlewares(
	auth *AuthMiddleware,
) *Middleware {
	return &Middleware{
		AuthMiddleware: auth,
	}
}

func NewMiddleware() fx.Option {
	return fx.Module(
		"middleware",
		fx.Provide(
			NewAuthMiddleware,
			NewMiddlewares,
		),
	)
}
