package caddyInterface

import (
	"cadigo-api/app/modelA"
	"context"
)

type CaddyRepository interface {
	Create(ctx context.Context, caddy *modelA.Caddy) (*modelA.Caddy, error)
	GetAll(ctx context.Context, pagination modelA.Pagination) (result []*modelA.Caddy, total int64, err error)
}

type CaddyService interface {
	Create(ctx context.Context, caddy *modelA.Caddy) (*modelA.Caddy, error)
	GetAll(ctx context.Context, pagination modelA.Pagination) (result []*modelA.Caddy, total int64, err error)
}
