package user_service

import (
	"go.uber.org/fx"
)

func NewUser() fx.Option {
	return fx.Module(
		"user-service",
		fx.Provide(
			fx.Annotate(
				NewUserService,
				fx.As(new(UserServiceInterface)),
			),
		),
	)
}
