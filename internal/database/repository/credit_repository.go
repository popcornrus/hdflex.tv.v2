package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
	"go-hdflex/internal/database/model"
)

type CreditRepository struct {
	db *sql.DB
}

type CreditRepositoryInterface interface {
	Create(context.Context, model.Credit) (int64, error)
	GetCreditByTmdbId(context.Context, int64) (model.Credit, error)
}

func NewCreditRepository(
	db *sql.DB,
) *CreditRepository {
	return &CreditRepository{
		db: db,
	}
}

func (r *CreditRepository) Create(ctx context.Context, credit model.Credit) (int64, error) {
	const op = "CreditRepository.Create() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.Insert("credits").
		Rows(
			goqu.Record{
				"external_id": credit.ExternalId,
				"name":        credit.Name,
				"orig_name":   credit.OrigName,
				"popularity":  credit.Popularity,
				"image":       credit.Image,
			},
		).
		ToSQL()
	if err != nil {
		return 0, fmt.Errorf("%s failed to build query: %w", op, err)
	}

	result, err := r.db.ExecContext(ctx, query, params...)
	if err != nil {
		return 0, fmt.Errorf("%s failed to execute query: %w", op, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s failed to get last insert id: %w", op, err)
	}

	return id, nil
}

func (r *CreditRepository) GetCreditByTmdbId(ctx context.Context, tmdbId int64) (model.Credit, error) {
	const op = "CreditRepository.GetCreditByTmdbId() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.From("credits").
		Select(
			"id",
			"external_id",
			"name",
			"orig_name",
			"popularity",
			"image",
			"created_at",
			"updated_at",
		).
		Where(goqu.Ex{
			"external_id": tmdbId,
		}).
		ToSQL()
	if err != nil {
		return model.Credit{}, fmt.Errorf("%s failed to build query: %w", op, err)
	}

	var credit model.Credit
	if err := r.db.QueryRowContext(ctx, query, params...).Scan(
		&credit.Id,
		&credit.ExternalId,
		&credit.Name,
		&credit.OrigName,
		&credit.Popularity,
		&credit.Image,
	); err != nil {
		return model.Credit{}, fmt.Errorf("%s failed to scan credit: %w", op, err)
	}

	return credit, nil
}
