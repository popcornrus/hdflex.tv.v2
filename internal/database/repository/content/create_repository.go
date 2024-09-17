package repository

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"go-hdflex/internal/database/model"
	"log/slog"
)

type CreateContentActionInterface interface {
	Create(context.Context, model.Content) (int64, error)
	CreateContentExternalId(context.Context, model.ContentExternalId) (int64, error)
	CreateContentTranslation(context.Context, model.ContentTranslation) (int64, error)
	CreateContentGenre(context.Context, model.ContentGenre) (int64, error)
	CreateContentCountry(context.Context, model.ContentCountry) (int64, error)
	CreateContentCast(context.Context, model.ContentCast) (int64, error)
	CreateContentCrew(context.Context, model.ContentCrew) (int64, error)
	CreateContentSimilar(context.Context, model.ContentSimilar) (int64, error)
}

func (r *ContentRepository) Create(ctx context.Context, content model.Content) (int64, error) {
	const op = "ContentRepository.Create() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.Insert("contents").Rows(
		goqu.Record{
			"cdnmovies_id":   content.CdnMoviesId,
			"ru_title":       content.RuTitle,
			"orig_title":     content.OrigTitle,
			"url":            content.Url,
			"popularity":     content.Popularity,
			"en_title":       content.EnTitle,
			"slogan":         content.Slogan,
			"description":    content.Description,
			"duration":       content.Duration,
			"iframe_url":     content.IframeUrl,
			"content_type":   content.ContentType,
			"year":           content.Year,
			"poster":         content.Poster,
			"backdrop":       content.Backdrop,
			"rating_age":     content.RatingAge,
			"rating_mpaa":    content.RatingMpaa,
			"world_premiere": content.WorldPremiere,
			"ru_premiere":    content.RuPremiere,
			"last_season":    content.LastSeason,
			"last_episode":   content.LastEpisode,
			"lgbt":           content.Lgbt,
		},
	).ToSQL()
	if err != nil {
		slog.Error("Error creating content sql", err)
		return 0, fmt.Errorf("%s error building query: %w", op, err)
	}

	result, err := r.db.ExecContext(ctx, query, params...)
	if err != nil {
		return 0, fmt.Errorf("%s error creating content: %w", op, err)
	}

	return result.LastInsertId()
}

func (r *ContentRepository) CreateContentExternalId(ctx context.Context, externalId model.ContentExternalId) (int64, error) {
	const op = "ContentRepository.CreateContentExternalId() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.Insert("content_external_ids").Rows(
		goqu.Record{
			"content_id":            externalId.ContentId,
			"external_id":           externalId.ExternalId,
			"external_type":         externalId.ExternalType,
			"external_rating":       externalId.ExternalRating,
			"external_rating_votes": externalId.ExternalRatingVotes,
		},
	).ToSQL()
	if err != nil {
		slog.Error("Error creating content external id sql", err)
		return 0, err
	}

	result, err := r.db.ExecContext(ctx, query, params...)
	if err != nil {
		return 0, fmt.Errorf("%s error creating content external id: %w", op, err)
	}

	return result.LastInsertId()
}

func (r *ContentRepository) CreateContentTranslation(ctx context.Context, translation model.ContentTranslation) (int64, error) {
	const op = "ContentRepository.CreateContentTranslation() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.Insert("content_translations").Rows(
		translation,
	).ToSQL()
	if err != nil {
		return 0, fmt.Errorf("%s error building query: %w", op, err)
	}

	result, err := r.db.ExecContext(ctx, query, params...)
	if err != nil {
		slog.Error("Error creating translation", err)
		return 0, fmt.Errorf("%s error creating translation: %w", op, err)
	}

	return result.LastInsertId()
}

func (r *ContentRepository) CreateContentGenre(ctx context.Context, contentGenre model.ContentGenre) (int64, error) {
	const op = "ContentRepository.CreateContentGenre() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.Insert("content_genres").Rows(
		contentGenre,
	).ToSQL()
	if err != nil {
		slog.Error("Error creating content genre sql", err)
		return 0, fmt.Errorf("%s error building query: %w", op, err)
	}

	result, err := r.db.ExecContext(ctx, query, params...)
	if err != nil {
		return 0, fmt.Errorf("%s error creating content genre: %w", op, err)
	}

	return result.LastInsertId()
}

func (r *ContentRepository) CreateContentCountry(ctx context.Context, contentCountry model.ContentCountry) (int64, error) {
	const op = "ContentRepository.CreateContentCountry() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.Insert("content_countries").Rows(
		contentCountry,
	).ToSQL()
	if err != nil {
		return 0, fmt.Errorf("%s error building query: %w", op, err)
	}

	result, err := r.db.ExecContext(ctx, query, params...)
	if err != nil {
		slog.Error("Error creating content country", err)
		return 0, fmt.Errorf("%s error creating content country: %w", op, err)
	}

	return result.LastInsertId()
}

func (r *ContentRepository) CreateContentCast(ctx context.Context, contentCast model.ContentCast) (int64, error) {
	const op = "ContentRepository.CreateContentCast() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.Insert("content_casts").Rows(
		contentCast,
	).ToSQL()
	if err != nil {
		return 0, fmt.Errorf("%s error building query: %w", op, err)
	}

	result, err := r.db.ExecContext(ctx, query, params...)
	if err != nil {
		return 0, fmt.Errorf("%s error creating content cast: %w", op, err)
	}

	return result.LastInsertId()
}

func (r *ContentRepository) CreateContentCrew(ctx context.Context, contentCrew model.ContentCrew) (int64, error) {
	const op = "ContentRepository.CreateContentCrew() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.Insert("content_crews").Rows(
		contentCrew,
	).ToSQL()
	if err != nil {
		return 0, fmt.Errorf("%s error building query: %w", op, err)
	}

	result, err := r.db.ExecContext(ctx, query, params...)
	if err != nil {
		slog.Error("Error creating content crew", err)
		return 0, fmt.Errorf("%s error creating content crew: %w", op, err)
	}

	return result.LastInsertId()
}

func (r *ContentRepository) CreateContentSimilar(ctx context.Context, contentSimilar model.ContentSimilar) (int64, error) {
	const op = "ContentRepository.CreateContentSimilar() ->"

	dialect := goqu.Dialect("mysql")

	query, params, err := dialect.Insert("content_similars").Rows(
		contentSimilar,
	).ToSQL()
	if err != nil {
		return 0, fmt.Errorf("%s error building query: %w", op, err)
	}

	result, err := r.db.ExecContext(ctx, query, params...)
	if err != nil {
		return 0, fmt.Errorf("%s error creating content similar: %w", op, err)
	}

	return result.LastInsertId()
}
