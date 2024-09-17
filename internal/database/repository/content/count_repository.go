package repository

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	f "go-hdflex/external/filter"
	"go-hdflex/external/helpers"
)

type CountContentRepositoryInterface interface {
	CountByContext(context.Context, *f.GetFilter) (int, error)
}

func (r *ContentRepository) CountByContext(ctx context.Context, filters *f.GetFilter) (int, error) {
	const op = "ContentRepository.CountByContext() ->"

	dialect := goqu.Dialect("mysql")

	builder := dialect.From("contents").
		Select(goqu.COUNT("id"))

	builder = helpers.ApplyParams(builder, filters)

	query, params, err := builder.ToSQL()
	if err != nil {
		return 0, fmt.Errorf("%s error building query: %w", op, err)
	}

	var count int
	err = r.db.QueryRowContext(ctx, query, params...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("%s error getting count: %w", op, err)
	}

	return count, nil
}
