package repository

import (
	"go.uber.org/fx"
)

func NewRepository() fx.Option {
	return fx.Module(
		"repository",
		fx.Provide(
			fx.Annotate(
				NewUserRepository,
				fx.As(new(UserRepositoryInterface)),
			),
		),
	)
}
