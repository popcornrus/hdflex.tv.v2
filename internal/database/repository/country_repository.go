package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"go-hdflex/internal/database/model"
)

type CountryRepository struct {
	db *sql.DB
}

type CountryRepositoryInterface interface {
	Get(context.Context) ([]model.Country, error)
	FindCountryByTitle(context.Context, string) (model.Country, error)
}

func NewCountryRepository(
	db *sql.DB,
) *CountryRepository {
	return &CountryRepository{
		db: db,
	}
}

func (r *CountryRepository) Get(ctx context.Context) ([]model.Country, error) {
	const op = "CountryRepository.Get() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.From("countries").
		Select("id", "title").
		ToSQL()

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, fmt.Errorf("%s failed to get countries: %w", op, err)
	}

	defer rows.Close()

	var countries []model.Country

	for rows.Next() {
		var country model.Country

		err := rows.Scan(
			&country.Id,
			&country.Title,
		)

		if err != nil {
			return nil, fmt.Errorf("%s failed to scan country: %w", op, err)
		}

		countries = append(countries, country)
	}

	return countries, nil
}

func (r *CountryRepository) FindCountryByTitle(ctx context.Context, title string) (model.Country, error) {
	const op = "CountryRepository.FindCountryByTitle() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.From("countries").
		Select("id", "title").
		Where(goqu.Ex{
			"title": title,
		}).
		ToSQL()

	row := r.db.QueryRowContext(ctx, query, params...)

	var country model.Country

	err = row.Scan(
		&country.Id,
		&country.Title,
	)

	if err != nil {
		return model.Country{}, fmt.Errorf("%s failed to scan country: %w", op, err)
	}

	return country, nil
}
