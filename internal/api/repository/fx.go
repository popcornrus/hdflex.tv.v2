package repository

import (
	"go.uber.org/fx"
)

func NewRepository() fx.Option {
	return fx.Module(
		"repository",
		fx.Provide(),
	)
}
