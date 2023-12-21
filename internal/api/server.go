package api

import (
	"context"
	"github.com/go-chi/chi/v5"
	"go-boilerplate/external/config"
	"go-boilerplate/external/logger/sl"
	"go.uber.org/fx"
	"log/slog"
	"net/http"
)

func NewServer(cfg *config.Config, r *chi.Mux) *http.Server {
	return &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
}

func RunServer(
	lc fx.Lifecycle,
	log *slog.Logger,
	server *http.Server,
) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil {
					log.Error("Server failed", sl.Err(err))
				} else {
					log.Info("Server started")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Error("Server stopped")
			return server.Shutdown(ctx)
		},
	})
}
