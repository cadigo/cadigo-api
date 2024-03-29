package caddyservice

import (
	"cadigo-api/app/interface/caddyinterface"
	"cadigo-api/app/modela"
	"context"
)

type Service struct {
	repo caddyinterface.CaddyRepository
}

func NewService(repo caddyinterface.CaddyRepository) caddyinterface.CaddyService {
	return &Service{
		repo: repo,
	}
}

func (serv *Service) Create(ctx context.Context, caddy *modela.Caddy) (*modela.Caddy, error) {
	return serv.repo.Create(ctx, caddy)
}

func (serv *Service) GetAll(ctx context.Context, pagination modela.Pagination, filter modela.CaddyFilter) (result []*modela.Caddy, total int64, err error) {
	return serv.repo.GetAll(ctx, pagination, filter)
}

func (serv *Service) Update(ctx context.Context, argID string, record *modela.Caddy) (*modela.Caddy, error) {
	return serv.repo.Update(ctx, argID, record)
}

func (serv *Service) GetByID(ctx context.Context, id string) (result *modela.Caddy, err error) {
	return serv.repo.GetByID(ctx, id)
}

func (serv *Service) GetByReference(ctx context.Context, reference string) (result *modela.Caddy, err error) {
	return serv.repo.GetByReference(ctx, reference)
}
