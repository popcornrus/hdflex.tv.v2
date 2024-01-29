package handler

import (
	"go.uber.org/fx"
)

type Handlers struct {
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func NewHandler() fx.Option {
	return fx.Module(
		"handler",
		fx.Options(),
		fx.Provide(
			NewHandlers,
		),
	)
}
