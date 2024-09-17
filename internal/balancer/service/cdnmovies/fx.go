package cdnmovies

import (
	"go.uber.org/fx"
)

func NewCdnMovies() fx.Option {
	return fx.Module(
		"cdnmovies-service",
		fx.Provide(
			fx.Annotate(
				NewBalancerService,
				fx.As(new(BalancerServiceInterface)),
			),
			fx.Annotate(
				NewLooperService,
				fx.As(new(LooperServiceInterface)),
			),
		),
	)
}
