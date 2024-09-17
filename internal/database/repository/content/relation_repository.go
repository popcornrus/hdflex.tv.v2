package repository

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_struct "go-hdflex/internal/database/struct"
)

type RelationContentRepositoryInterface interface {
	GetContentGenres(context.Context, int64) ([]_struct.Genre, error)
	GetContentCountries(context.Context, int64) ([]_struct.Country, error)
	GetContentExternalIds(context.Context, int64) ([]_struct.ExternalId, error)
	GetContentCasts(context.Context, int64) ([]_struct.Cast, error)
	GetContentCrews(context.Context, int64) ([]_struct.Crew, error)
}

func (r *ContentRepository) GetContentGenres(ctx context.Context, contentId int64) ([]_struct.Genre, error) {
	const op = "ContentRepository.GetContentGenres() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.From("content_genres").
		Select(
			"genres.id",
			"genres.title",
			"genres.en_title",
		).
		LeftJoin(goqu.T("genres").As("genres"), goqu.On(goqu.Ex{"content_genres.genre_id": goqu.I("genres.id")})).
		Where(goqu.Ex{"content_id": contentId}).ToSQL()

	if err != nil {
		return nil, fmt.Errorf("%s error building query: %w", op, err)
	}

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, fmt.Errorf("%s error getting content genres: %w", op, err)
	}

	defer rows.Close()

	var items []_struct.Genre
	for rows.Next() {
		var item _struct.Genre
		if err := rows.Scan(
			&item.Id,
			&item.Title,
			&item.EnTitle,
		); err != nil {
			return nil, fmt.Errorf("%s error scanning content genres: %w", op, err)
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *ContentRepository) GetContentCountries(ctx context.Context, contentId int64) ([]_struct.Country, error) {
	const op = "ContentRepository.GetContentCountries() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.From("content_countries").
		Select(
			"countries.id",
			"countries.title",
			"countries.en_title",
		).
		LeftJoin(goqu.T("countries").As("countries"), goqu.On(goqu.Ex{"content_countries.country_id": goqu.I("countries.id")})).
		Where(goqu.Ex{"content_id": contentId}).ToSQL()

	if err != nil {
		return nil, fmt.Errorf("%s error building query: %w", op, err)
	}

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, fmt.Errorf("%s error getting content countries: %w", op, err)
	}

	defer rows.Close()

	var items []_struct.Country
	for rows.Next() {
		var item _struct.Country
		if err := rows.Scan(
			&item.Id,
			&item.Title,
			&item.EnTitle,
		); err != nil {
			return nil, fmt.Errorf("%s error scanning content countries: %w", op, err)
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *ContentRepository) GetContentExternalIds(ctx context.Context, contentId int64) ([]_struct.ExternalId, error) {
	const op = "ContentRepository.GetContentExternalIds() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.From("content_external_ids").
		Select(
			"content_external_ids.id",
			"content_external_ids.external_id",
			"content_external_ids.external_type",
			"content_external_ids.external_rating",
			"content_external_ids.external_rating_votes",
		).
		Where(goqu.Ex{"content_id": contentId}).ToSQL()

	if err != nil {
		return nil, fmt.Errorf("%s error building query: %w", op, err)
	}

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, fmt.Errorf("%s error getting content external ids: %w", op, err)
	}

	defer rows.Close()

	var items []_struct.ExternalId
	for rows.Next() {
		var item _struct.ExternalId
		if err := rows.Scan(
			&item.Id,
			&item.ExternalId,
			&item.ExternalTypeId,
			&item.Rating,
			&item.Votes,
		); err != nil {
			return nil, fmt.Errorf("%s error scanning content external ids: %w", op, err)
		}

		item.ExternalType = item.ExternalTypeId.String()

		items = append(items, item)
	}

	return items, nil
}

func (r *ContentRepository) GetContentCasts(ctx context.Context, contentId int64) ([]_struct.Cast, error) {
	const op = "ContentRepository.GetContentCasts() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.From("content_casts").
		Select(
			"content_casts.id",
			"content_casts.department",
			"content_casts.character",
			"credits.popularity",
			"credits.name",
			"credits.orig_name",
			"credits.image",
		).
		LeftJoin(goqu.T("credits").As("credits"), goqu.On(goqu.Ex{"content_casts.credit_id": goqu.I("credits.id")})).
		Where(goqu.Ex{"content_id": contentId}).
		Order(goqu.I("content_casts.sort").Asc()).
		ToSQL()

	if err != nil {
		return nil, fmt.Errorf("%s error building query: %w", op, err)
	}

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, fmt.Errorf("%s error getting content casts: %w", op, err)
	}

	defer rows.Close()

	var items []_struct.Cast
	for rows.Next() {
		var item _struct.Cast
		if err := rows.Scan(
			&item.Id,
			&item.Department,
			&item.Character,
			&item.Person.Popularity,
			&item.Person.Name,
			&item.Person.OrigName,
			&item.Person.Image,
		); err != nil {
			return nil, fmt.Errorf("%s error scanning content casts: %w", op, err)
		}

		items = append(items, item)
	}

	return items, nil
}

func (r *ContentRepository) GetContentCrews(ctx context.Context, contentId int64) ([]_struct.Crew, error) {
	const op = "ContentRepository.GetContentCrews() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.From("content_crews").
		Select(
			"content_crews.id",
			"content_crews.department",
			"content_crews.job",
			"credits.popularity",
			"credits.name",
			"credits.orig_name",
			"credits.image",
		).
		LeftJoin(goqu.T("credits").As("credits"), goqu.On(goqu.Ex{"content_crews.credit_id": goqu.I("credits.id")})).
		Where(goqu.Ex{"content_id": contentId}).
		ToSQL()

	if err != nil {
		return nil, fmt.Errorf("%s error building query: %w", op, err)
	}

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, fmt.Errorf("%s error getting content crews: %w", op, err)
	}

	defer rows.Close()

	var items []_struct.Crew
	for rows.Next() {
		var item _struct.Crew
		if err := rows.Scan(
			&item.Id,
			&item.Department,
			&item.Job,
			&item.Person.Popularity,
			&item.Person.Name,
			&item.Person.OrigName,
			&item.Person.Image,
		); err != nil {
			return nil, fmt.Errorf("%s error scanning content crews: %w", op, err)
		}

		items = append(items, item)
	}

	return items, nil
}
