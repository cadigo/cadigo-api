package caddyservice

import (
	"cadigo-api/app/interface/caddyInterface"
	"cadigo-api/app/modela"
	"context"
)

type Service struct {
	repo caddyInterface.CaddyRepository
}

func NewService(repo caddyInterface.CaddyRepository) caddyInterface.CaddyService {
	return &Service{
		repo: repo,
	}
}

func (serv *Service) Create(ctx context.Context, caddy *modela.Caddy) (*modela.Caddy, error) {
	return serv.repo.Create(ctx, caddy)
}

func (serv *Service) GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.Caddy, total int64, err error) {
	return serv.repo.GetAll(ctx, pagination)
}

func (serv *Service) Update(ctx context.Context, argID string, record *modela.Caddy) (*modela.Caddy, error) {
	return serv.repo.Update(ctx, argID, record)
}
