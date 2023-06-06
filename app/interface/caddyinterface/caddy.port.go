package caddyinterface

import (
	"cadigo-api/app/modela"
	"context"
)

type CaddyRepository interface {
	Create(ctx context.Context, caddy *modela.Caddy) (*modela.Caddy, error)
	Update(ctx context.Context, argID string, record *modela.Caddy) (*modela.Caddy, error)
	GetAll(ctx context.Context, pagination modela.Pagination, filter modela.CaddyFilter) (result []*modela.Caddy, total int64, err error)
	GetByID(ctx context.Context, id string) (result *modela.Caddy, err error)
	GetByReference(ctx context.Context, reference string) (result *modela.Caddy, err error)
}

type CaddyService interface {
	Create(ctx context.Context, caddy *modela.Caddy) (*modela.Caddy, error)
	Update(ctx context.Context, argID string, record *modela.Caddy) (*modela.Caddy, error)
	GetAll(ctx context.Context, pagination modela.Pagination, filter modela.CaddyFilter) (result []*modela.Caddy, total int64, err error)
	GetByID(ctx context.Context, id string) (result *modela.Caddy, err error)
	GetByReference(ctx context.Context, reference string) (result *modela.Caddy, err error)
}
