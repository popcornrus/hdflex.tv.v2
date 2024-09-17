package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/patrickmn/go-cache"
	"go-hdflex/external/config"
	"go-hdflex/external/db"
	"go-hdflex/internal/api/http/handler"
	"go-hdflex/internal/api/http/middleware"
	"go-hdflex/internal/api/repository"
	"go-hdflex/internal/api/service"
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Options(
			repository.NewRepository(),
			service.NewService(),
			handler.NewHandler(),
			middleware.NewMiddleware(),
			db.NewDataBase(),
		),
		fx.Provide(
			config.NewConfig,
			validator.New,
			NewCache,
			NewLogger,
			NewRouter,
			NewServer,
		),
		fx.Invoke(RunServer),
	)
}

func NewCache() *cache.Cache {
	return cache.New(cache.NoExpiration, cache.NoExpiration)
}
