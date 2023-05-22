package caddyinterface

import (
	"cadigo-api/app/modela"
	"context"
)

type CaddyRepository interface {
	Create(ctx context.Context, caddy *modela.Caddy) (*modela.Caddy, error)
	Update(ctx context.Context, argID string, record *modela.Caddy) (*modela.Caddy, error)
	GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.Caddy, total int64, err error)
}

type CaddyService interface {
	Create(ctx context.Context, caddy *modela.Caddy) (*modela.Caddy, error)
	Update(ctx context.Context, argID string, record *modela.Caddy) (*modela.Caddy, error)
	GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.Caddy, total int64, err error)
}
