package caddyhdl

import (
	"cadigo-api/app/ports/caddyport"
	modelgraph "cadigo-api/graph/model"
	"context"
	"fmt"
)

type Handler struct {
	servCaddy *caddyport.CaddyService
}

func NewHandler(servCaddy caddyport.CaddyService) *Handler {
	return &Handler{
		servCaddy: &servCaddy,
	}
}

// Caddy is the resolver for the Caddy field.
func (r *Handler) Caddy(ctx context.Context, input modelgraph.CaddyInput) (*modelgraph.Caddy, error) {
	panic(fmt.Errorf("not implemented: Caddy - Caddy"))
}
