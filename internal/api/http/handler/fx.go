package handler

import (
	"go.uber.org/fx"
)

type Handlers struct {
	Content *ContentHandler
	Storage *StorageHandler
	Genre   *GenreHandler
	Country *CountryHandler
}

func NewHandlers(
	ch *ContentHandler,
	sh *StorageHandler,
	gh *GenreHandler,
	cth *CountryHandler,
) *Handlers {
	return &Handlers{
		Content: ch,
		Storage: sh,
		Genre:   gh,
		Country: cth,
	}
}

func NewHandler() fx.Option {
	return fx.Module(
		"handler",
		fx.Provide(
			NewContentHandler,
			NewStorageHandler,
			NewGenreHandler,
			NewCountryHandler,
			NewHandlers,
		),
	)
}
