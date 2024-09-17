package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"go-hdflex/internal/database/model"
)

type GenreRepository struct {
	db *sql.DB
}

type GenreRepositoryInterface interface {
	Get(context.Context) ([]model.Genre, error)
	FindFirstByTitle(context.Context, string) (model.Genre, error)
}

func NewGenreRepository(
	db *sql.DB,
) *GenreRepository {
	return &GenreRepository{
		db: db,
	}
}

func (r *GenreRepository) Get(ctx context.Context) ([]model.Genre, error) {
	const op = "GenreRepository.Get() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.From("genres").
		Select("id", "title", "en_title").
		Order(goqu.I("title").Asc()).
		ToSQL()

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, fmt.Errorf("%s failed to query genres: %w", op, err)
	}
	defer rows.Close()

	var genres []model.Genre

	for rows.Next() {
		var genre model.Genre

		err = rows.Scan(
			&genre.Id,
			&genre.Title,
			&genre.EnTitle,
		)
		if err != nil {
			return nil, fmt.Errorf("%s failed to scan genre: %w", op, err)
		}

		genres = append(genres, genre)
	}

	return genres, nil
}

func (r *GenreRepository) FindFirstByTitle(ctx context.Context, title string) (model.Genre, error) {
	const op = "GenreRepository.FindFirstByTitle() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.From("genres").
		Select("id", "title", "en_title").
		Where(goqu.Ex{
			"title": title,
		}).
		ToSQL()

	row := r.db.QueryRowContext(ctx, query, params...)

	var genre model.Genre

	err = row.Scan(
		&genre.Id,
		&genre.Title,
		&genre.EnTitle,
	)
	if err != nil {
		return model.Genre{}, fmt.Errorf("%s failed to scan genre: %w", op, err)
	}

	return genre, nil
}
