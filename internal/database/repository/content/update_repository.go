package repository

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
)

type UpdateContentActionInterface interface {
	UpdateLastWatchAt(context.Context, int64) error
}

func (r *ContentRepository) UpdateLastWatchAt(ctx context.Context, id int64) error {
	const op = "ContentRepository.UpdateLastWatchAt() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.Update("contents").
		Set(goqu.Record{"last_watch_at": goqu.L("NOW()")}).
		Where(goqu.Ex{"id": id}).ToSQL()

	if err != nil {
		return fmt.Errorf("%s error building query: %w", op, err)
	}

	_, err = r.db.ExecContext(ctx, query, params...)
	if err != nil {
		return fmt.Errorf("%s error updating last watch at: %w", op, err)
	}

	return nil
}
