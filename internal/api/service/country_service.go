package service

import (
	"context"
	"fmt"
	"go-hdflex/internal/api/http/response"
	"go-hdflex/internal/database/repository"
)

type CountryService struct {
	r repository.CountryRepositoryInterface
}

type CountryServiceInterface interface {
	Get(context.Context) (response.CountryGetResponse, error)
}

func NewCountryService(
	r repository.CountryRepositoryInterface,
) *CountryService {
	return &CountryService{
		r: r,
	}
}

func (s *CountryService) Get(ctx context.Context) (response.CountryGetResponse, error) {
	const op = "CountryService.Get() ->"

	countries, err := s.r.Get(ctx)
	if err != nil {
		return response.CountryGetResponse{}, fmt.Errorf("%s %w", op, err)
	}

	return response.CountryGetResponse{
		Items: countries,
	}, nil
}
