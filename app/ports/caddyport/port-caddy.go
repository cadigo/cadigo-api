package caddyport

import (
	"cadigo-api/app/modelapp"
	"context"
)

type CaddyRepository interface {
	Create(ctx context.Context, caddy *modelapp.Caddy) (*modelapp.Caddy, error)
}

type CaddyService interface {
	Create(ctx context.Context, caddy *modelapp.Caddy) (*modelapp.Caddy, error)
}
