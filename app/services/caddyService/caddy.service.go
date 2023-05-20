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
