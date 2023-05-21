package caddyService

import (
	"cadigo-api/app/interface/caddyInterface"
	"cadigo-api/app/modelA"
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

func (serv *Service) Create(ctx context.Context, caddy *modelA.Caddy) (*modelA.Caddy, error) {
	return serv.repo.Create(ctx, caddy)
}

func (serv *Service) GetAll(ctx context.Context, pagination modelA.Pagination) (result []*modelA.Caddy, total int64, err error) {
	return serv.repo.GetAll(ctx, pagination)
}
