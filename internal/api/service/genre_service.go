package service

import (
	"context"
	"fmt"
	"go-hdflex/internal/api/http/response"
	"go-hdflex/internal/database/repository"
)

type GenreService struct {
	r repository.GenreRepositoryInterface
}

type GenreServiceInterface interface {
	Get(context.Context) (response.GenreGetResponse, error)
}

func NewGenreService(
	r repository.GenreRepositoryInterface,
) *GenreService {
	return &GenreService{
		r: r,
	}
}

func (s *GenreService) Get(ctx context.Context) (response.GenreGetResponse, error) {
	const op = "GenreService.Get() ->"

	genres, err := s.r.Get(ctx)
	if err != nil {
		return response.GenreGetResponse{}, fmt.Errorf("%s %w", op, err)
	}

	return response.GenreGetResponse{
		Items: genres,
	}, nil
}
