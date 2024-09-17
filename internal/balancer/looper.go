package balancer

import (
	"context"
	"go-hdflex/internal/balancer/service/cdnmovies"
	"go.uber.org/fx"
	"log/slog"
)

func RunLooper(
	lc fx.Lifecycle,
	log *slog.Logger,
	lcdnmovies cdnmovies.LooperServiceInterface,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go lcdnmovies.Start(context.WithoutCancel(ctx))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Error("Server stopped")
			return nil
		},
	})
}
