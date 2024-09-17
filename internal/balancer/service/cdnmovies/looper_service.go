package cdnmovies

import (
	"context"
	"log/slog"
)

type LooperService struct {
	log *slog.Logger
	bls BalancerServiceInterface
}

type LooperServiceInterface interface {
	Start(context.Context)
}

func NewLooperService(
	log *slog.Logger,
	bls BalancerServiceInterface,
) *LooperService {
	return &LooperService{
		log: log,
		bls: bls,
	}
}

func (p *LooperService) Start(ctx context.Context) {
	p.bls.Parse(ctx)
	/*for {
		go

		select {
		case <-ctx.Done():
			return
		default:
		}

		time.Sleep(5 * time.Second)
	}*/
}
