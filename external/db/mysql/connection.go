package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"go-hdflex/external/config"
	"go.uber.org/fx"
	"log/slog"
)

func NewMysqlDatabase(lc fx.Lifecycle, cfg *config.Config) (*sql.DB, error) {
	const op = "db.NewMysqlDatabase() ->"

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local",
		cfg.MySql.User,
		cfg.MySql.Password,
		cfg.MySql.Host,
		cfg.MySql.Port,
		cfg.MySql.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	db.SetMaxOpenConns(cfg.MySql.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.MySql.MaxIdleConnections)
	db.SetConnMaxLifetime(cfg.MySql.MaxConnectionLifetime)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			slog.Error("Closing database connection", fmt.Errorf("%s", op))
			return db.Close()
		},
	})

	return db, nil
}
