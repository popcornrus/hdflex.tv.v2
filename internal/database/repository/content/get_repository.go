package repository

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	f "go-hdflex/external/filter"
	"go-hdflex/external/helpers"
	"go-hdflex/internal/database/enum"
	"go-hdflex/internal/database/model"
	_struct "go-hdflex/internal/database/struct"
)

type GetContentRepositoryInterface interface {
	GetBySlug(context.Context, string) (_struct.ShowContent, error)
	GetByCdnMoviesId(context.Context, string) (model.Content, error)
	GetByTmdbId(context.Context, int64) (model.Content, error)
	GetByContext(context.Context, *f.GetFilter) ([]_struct.GetContent, error)
}

func (r *ContentRepository) GetByContext(ctx context.Context, filters *f.GetFilter) ([]_struct.GetContent, error) {
	const op = "ContentRepository.GetByContext() ->"

	dialect := goqu.Dialect("mysql")

	builder := dialect.From("contents").
		Select(
			"contents.id",
			"contents.ru_title",
			"contents.en_title",
			"contents.orig_title",
			"contents.url",
			"contents.content_type",
			"contents.year",
			"contents.poster",
			"contents.rating_age",
			"contents.last_season",
			"contents.last_episode",
			"contents.lgbt",
			"contents.created_at",
			"contents.updated_at",
		)

	builder = helpers.ApplyParams(builder, filters)

	query, params, err := builder.ToSQL()
	if err != nil {
		return nil, fmt.Errorf("%s error building query: %w", op, err)
	}

	rows, err := r.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, fmt.Errorf("%s error getting content: %w", op, err)
	}

	defer rows.Close()

	var items []_struct.GetContent
	for rows.Next() {
		var item _struct.GetContent
		if err := rows.Scan(
			&item.Id,
			&item.RuTitle,
			&item.EnTitle,
			&item.OrigTitle,
			&item.Url,
			&item.ContentType,
			&item.Year,
			&item.Poster,
			&item.RatingAge,
			&item.LastSeason,
			&item.LastEpisode,
			&item.Lgbt,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("%s error scanning content: %w", op, err)
		}

		item.Genres, _ = r.GetContentGenres(ctx, item.Id)
		item.Countries, _ = r.GetContentCountries(ctx, item.Id)
		item.ExternalIds, _ = r.GetContentExternalIds(ctx, item.Id)

		items = append(items, item)
	}

	return items, nil
}

func (r *ContentRepository) GetByCdnMoviesId(ctx context.Context, cdnMoviesId string) (model.Content, error) {
	const op = "ContentRepository.GetByCdnMoviesId() ->"

	dialect := goqu.Dialect("mysql")

	query, _, err := dialect.From("contents").
		Select(
			"id",
			"cdnmovies_id",
			"ru_title",
			"en_title",
			"orig_title",
			"url",
			"slogan",
			"description",
			"duration",
			"iframe_url",
			"content_type",
			"year",
			"poster",
			"backdrop",
			"rating_age",
			"rating_mpaa",
			"world_premiere",
			"ru_premiere",
			"last_season",
			"last_episode",
			"lgbt",
			"created_at",
			"updated_at",
		).
		Where(goqu.Ex{"cdnmovies_id": cdnMoviesId}).ToSQL()
	if err != nil {
		return model.Content{}, fmt.Errorf("%s error finding content by cdn movies id: %w", op, err)
	}

	row := r.db.QueryRowContext(ctx, query)

	var content model.Content
	if err := row.Scan(
		&content.Id,
		&content.CdnMoviesId,
		&content.RuTitle,
		&content.EnTitle,
		&content.OrigTitle,
		&content.Url,
		&content.Slogan,
		&content.Description,
		&content.Duration,
		&content.IframeUrl,
		&content.ContentType,
		&content.Year,
		&content.Poster,
		&content.Backdrop,
		&content.RatingAge,
		&content.RatingMpaa,
		&content.WorldPremiere,
		&content.RuPremiere,
		&content.LastSeason,
		&content.LastEpisode,
		&content.Lgbt,
		&content.CreatedAt,
		&content.UpdatedAt,
	); err != nil {
		return model.Content{}, fmt.Errorf("%s error finding content by cdn movies id: %w", op, err)
	}

	return content, nil
}

func (r *ContentRepository) GetByTmdbId(ctx context.Context, tmdbId int64) (model.Content, error) {
	const op = "ContentRepository.GetByTmdbId() ->"

	dialect := goqu.Dialect("mysql")

	query, _, err := dialect.From("contents").
		Select(
			"id",
			"cdnmovies_id",
			"ru_title",
			"en_title",
			"orig_title",
			"url",
			"slogan",
			"description",
			"duration",
			"iframe_url",
			"content_type",
			"year",
			"poster",
			"backdrop",
			"rating_age",
			"rating_mpaa",
			"world_premiere",
			"ru_premiere",
			"last_season",
			"last_episode",
			"lgbt",
			"created_at",
			"updated_at",
		).
		LeftJoin(
			goqu.T("content_external_ids").As("cei"),
			goqu.On(goqu.Ex{"contents.id": goqu.I("cei.content_id")},
				goqu.Ex{"cei.external_id_type": enum.TmdbExternalId},
			)).
		Where(goqu.Ex{"cei.external_id": tmdbId}).ToSQL()
	if err != nil {
		return model.Content{}, fmt.Errorf("%s error finding content by tmdb id: %w", op, err)
	}

	row := r.db.QueryRowContext(ctx, query)

	var content model.Content
	if err := row.Scan(
		&content.Id,
		&content.CdnMoviesId,
		&content.RuTitle,
		&content.EnTitle,
		&content.OrigTitle,
		&content.Url,
		&content.Slogan,
		&content.Description,
		&content.Duration,
		&content.IframeUrl,
		&content.ContentType,
		&content.Year,
		&content.Poster,
		&content.Backdrop,
		&content.RatingAge,
		&content.RatingMpaa,
		&content.WorldPremiere,
		&content.RuPremiere,
		&content.LastSeason,
		&content.LastEpisode,
		&content.Lgbt,
		&content.CreatedAt,
		&content.UpdatedAt,
	); err != nil {
		return model.Content{}, fmt.Errorf("%s error finding content by tmdb id: %w", op, err)
	}

	return content, nil
}

func (r *ContentRepository) GetBySlug(ctx context.Context, slug string) (_struct.ShowContent, error) {
	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.From("contents").
		Select(
			"id",
			"ru_title",
			"en_title",
			"orig_title",
			"slogan",
			"description",
			"iframe_url",
			"content_type",
			"duration",
			"year",
			"poster",
			"backdrop",
			"rating_age",
			"rating_mpaa",
			"last_season",
			"last_episode",
			"lgbt",
			"world_premiere",
			"created_at",
			"updated_at",
		).Where(goqu.Ex{"url": slug}).ToSQL()

	if err != nil {
		return _struct.ShowContent{}, err
	}

	var item _struct.ShowContent
	err = r.db.QueryRowContext(ctx, query, params...).Scan(
		&item.Id,
		&item.RuTitle,
		&item.EnTitle,
		&item.OrigTitle,
		&item.Slogan,
		&item.Description,
		&item.IframeUrl,
		&item.ContentType,
		&item.Duration,
		&item.Year,
		&item.Poster,
		&item.Backdrop,
		&item.RatingAge,
		&item.RatingMpaa,
		&item.LastSeason,
		&item.LastEpisode,
		&item.Lgbt,
		&item.WorldPremiere,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err != nil {
		return _struct.ShowContent{}, err
	}

	item.Genres, _ = r.GetContentGenres(ctx, item.Id)
	item.Countries, _ = r.GetContentCountries(ctx, item.Id)
	item.ExternalIds, _ = r.GetContentExternalIds(ctx, item.Id)
	item.Casts, _ = r.GetContentCasts(ctx, item.Id)
	item.Crew, _ = r.GetContentCrews(ctx, item.Id)

	return item, nil
}
