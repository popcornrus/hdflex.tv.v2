package repository

import (
	"database/sql"
	"go.uber.org/fx"
	"log/slog"

	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type ContentRepository struct {
	db  *sql.DB
	log *slog.Logger
}

type ContentRepositoryInterface interface {
	GetContentRepositoryInterface
	RelationContentRepositoryInterface
	CountContentRepositoryInterface
	CreateContentActionInterface
	UpdateContentActionInterface
}

func NewContentRepository(
	db *sql.DB,
	log *slog.Logger,
) *ContentRepository {
	return &ContentRepository{
		db:  db,
		log: log,
	}
}

func NewContentRepositoryOption() fx.Option {
	return fx.Module(
		"contents-repository",
		fx.Provide(
			fx.Annotate(
				NewContentRepository,
				fx.As(new(GetContentRepositoryInterface)),
				fx.As(new(RelationContentRepositoryInterface)),
				fx.As(new(CountContentRepositoryInterface)),
				fx.As(new(CreateContentActionInterface)),
				fx.As(new(UpdateContentActionInterface)),
				fx.As(new(ContentRepositoryInterface)),
			),
		),
	)
}
