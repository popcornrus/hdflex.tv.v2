package root

import (
	"github.com/go-playground/validator/v10"
	"github.com/patrickmn/go-cache"
	"go-boilerplate/external/config"
	"go-boilerplate/external/db"
	"go-boilerplate/internal/root/http/handler"
	"go-boilerplate/internal/root/http/middleware"
	"go-boilerplate/internal/root/repository"
	"go-boilerplate/internal/root/service"
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
