package service

import (
	"go.uber.org/fx"
)

func NewService() fx.Option {
	return fx.Module(
		"service",
	)
}
