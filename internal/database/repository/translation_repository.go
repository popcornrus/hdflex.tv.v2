package repository

import (
	"context"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"go-hdflex/internal/database/model"
)

type TranslationRepository struct {
	db *sql.DB
}

type TranslationRepositoryInterface interface {
	Create(context.Context, model.Translation) (int64, error)
	FindFirstByExternalId(context.Context, int64) (model.Translation, error)
}

func NewTranslationRepository(
	db *sql.DB,
) *TranslationRepository {
	return &TranslationRepository{
		db: db,
	}
}

func (r *TranslationRepository) Create(ctx context.Context, translation model.Translation) (int64, error) {
	dialect := goqu.Dialect("mysql")

	sql, _, err := dialect.Insert("translations").
		Rows(translation).
		ToSQL()

	if err != nil {
		return 0, err
	}

	result, err := r.db.ExecContext(ctx, sql)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TranslationRepository) FindFirstByExternalId(ctx context.Context, externalId int64) (model.Translation, error) {
	dialect := goqu.Dialect("mysql")

	sql, _, err := dialect.From("translations").
		Select("id", "cdnmovies_id", "title", "format_title").
		Where(goqu.Ex{
			"cdnmovies_id": externalId,
		}).
		ToSQL()

	if err != nil {
		return model.Translation{}, err
	}

	row := r.db.QueryRowContext(ctx, sql)

	var translation model.Translation

	err = row.Scan(
		&translation.Id,
		&translation.ExternalId,
		&translation.Title,
		&translation.FormatTitle,
	)
	if err != nil {
		return model.Translation{}, err
	}

	return translation, nil
}
