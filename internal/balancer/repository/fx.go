package repository

import (
	database "go-hdflex/internal/database/repository"
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
				database.NewTranslationRepository,
				fx.As(new(database.TranslationRepositoryInterface)),
			),
			fx.Annotate(
				database.NewGenreRepository,
				fx.As(new(database.GenreRepositoryInterface)),
			),
			fx.Annotate(
				database.NewCountryRepository,
				fx.As(new(database.CountryRepositoryInterface)),
			),
			fx.Annotate(
				database.NewCreditRepository,
				fx.As(new(database.CreditRepositoryInterface)),
			),
		),
	)
}
