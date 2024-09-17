package service

import (
	"go.uber.org/fx"
)

func NewService() fx.Option {
	return fx.Module(
		"service",
		fx.Provide(
			fx.Annotate(
				NewContentService,
				fx.As(new(ContentServiceInterface)),
			),
			fx.Annotate(
				NewStorageService,
				fx.As(new(StorageServiceInterface)),
			),
			fx.Annotate(
				NewGenreService,
				fx.As(new(GenreServiceInterface)),
			),
			fx.Annotate(
				NewCountryService,
				fx.As(new(CountryServiceInterface)),
			),
		),
	)
}
