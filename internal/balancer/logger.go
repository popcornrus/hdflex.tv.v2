package balancer

import (
	"go-hdflex/external/config"
	"go-hdflex/external/logger/handler/slogpretty"
	"log/slog"
	"os"
)

func NewLogger(cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	switch cfg.Env {
	case cfg.EnvState.Local:
		log = setupPrettySlogLogger()
	case cfg.EnvState.Dev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case cfg.EnvState.Prod:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func setupPrettySlogLogger() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	newPrettyHandler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(newPrettyHandler)
}
