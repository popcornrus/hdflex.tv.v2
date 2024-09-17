package service

import (
	"go-hdflex/internal/balancer/service/cdnmovies"
	"go-hdflex/internal/balancer/service/helpers"
	"go.uber.org/fx"
)

func NewService() fx.Option {
	return fx.Module(
		"service",
		fx.Options(
			cdnmovies.NewCdnMovies(),
		),
		fx.Provide(
			fx.Annotate(
				helpers.NewImageService,
				fx.As(new(helpers.ImageServiceInterface)),
			),
		),
	)
}
