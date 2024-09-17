package balancer

import (
	"github.com/go-playground/validator/v10"
	"github.com/patrickmn/go-cache"
	"go-hdflex/external/config"
	"go-hdflex/external/db"
	"go-hdflex/internal/balancer/repository"
	"go-hdflex/internal/balancer/service"
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Options(
			repository.NewRepository(),
			service.NewService(),
			db.NewDataBase(),
		),
		fx.Provide(
			config.NewConfig,
			validator.New,
			NewCache,
			NewLogger,
			NewTmdbClient,
		),
		fx.Invoke(RunLooper),
	)
}

func NewCache() *cache.Cache {
	return cache.New(cache.NoExpiration, cache.NoExpiration)
}
