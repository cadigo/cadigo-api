package caddyInterface

import (
	"cadigo-api/app/modelA"
	"context"
)

type CaddyRepository interface {
	Create(ctx context.Context, caddy *modelA.Caddy) (*modelA.Caddy, error)
}

type CaddyService interface {
	Create(ctx context.Context, caddy *modelA.Caddy) (*modelA.Caddy, error)
}
