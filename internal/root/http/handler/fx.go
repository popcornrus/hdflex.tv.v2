package handler

import (
	"go.uber.org/fx"
)

type Handlers struct {
	User *UserHandler
}

func NewHandlers(
	uh *UserHandler,
) *Handlers {
	return &Handlers{
		User: uh,
	}
}

func NewHandler() fx.Option {
	return fx.Module(
		"handler",
		fx.Options(),
		fx.Provide(
			NewUserHandler,
			NewHandlers,
		),
	)
}
