package repository

import (
	"go-hdflex/internal/database/repository"
	cr "go-hdflex/internal/database/repository/content"
	"go.uber.org/fx"
)

func NewRepository() fx.Option {
	return fx.Module(
		"repository",
		fx.Options(
			cr.NewContentRepositoryOption(),
		),
		fx.Provide(
			fx.Annotate(
				repository.NewGenreRepository,
				fx.As(new(repository.GenreRepositoryInterface)),
			),
			fx.Annotate(
				repository.NewCountryRepository,
				fx.As(new(repository.CountryRepositoryInterface)),
			),
		),
	)
}
