package service

import (
	"context"
	"fmt"
	f "go-hdflex/external/filter"
	"go-hdflex/internal/api/http/response"
	cr "go-hdflex/internal/database/repository/content"
)

type ContentService struct {
	r cr.ContentRepositoryInterface
}

type ContentServiceInterface interface {
	Get(context.Context) (response.ContentGetResponse, error)
	GetFirstBySlug(context.Context, string) (response.ContentDetailResponse, error)
	UpdateLastWatchAt(context.Context, int64) error
	Count(context.Context) (response.ContentCountResponse, error)
}

func NewContentService(
	r cr.ContentRepositoryInterface,
) *ContentService {
	return &ContentService{
		r: r,
	}
}

func (s *ContentService) Get(ctx context.Context) (response.ContentGetResponse, error) {
	const op = "ContentService.Get() ->"

	filters := ctx.Value("filters").(*f.GetFilter)

	var r response.ContentGetResponse
	var err error

	r.Items, err = s.r.GetByContext(ctx, filters)
	if err != nil {
		return r, fmt.Errorf("%s %w", op, err)
	}

	return r, nil
}

func (s *ContentService) GetFirstBySlug(ctx context.Context, slug string) (response.ContentDetailResponse, error) {
	const op = "ContentService.GetFirstBySlug() ->"

	var r response.ContentDetailResponse
	var err error

	r.Item, err = s.r.GetBySlug(ctx, slug)
	if err != nil {
		return r, fmt.Errorf("%s %w", op, err)
	}

	if err := s.UpdateLastWatchAt(ctx, r.Item.Id); err != nil {
		return r, fmt.Errorf("%s %w", op, err)
	}

	return r, nil
}

func (s *ContentService) Count(ctx context.Context) (response.ContentCountResponse, error) {
	const op = "ContentService.Count() ->"

	filters := ctx.Value("filters").(*f.GetFilter)

	var r response.ContentCountResponse
	var err error

	r.Count, err = s.r.CountByContext(ctx, filters)
	if err != nil {
		return r, fmt.Errorf("%s %w", op, err)
	}

	return r, nil
}

func (s *ContentService) UpdateLastWatchAt(ctx context.Context, id int64) error {
	const op = "ContentService.UpdateLastWatchAt() ->"

	err := s.r.UpdateLastWatchAt(ctx, id)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}
