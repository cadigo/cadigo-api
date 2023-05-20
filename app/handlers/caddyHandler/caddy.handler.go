package caddyHandler

import (
	"cadigo-api/app/interface/caddyInterface"
	"cadigo-api/app/modelA"
	"cadigo-api/graph/modelgraph"
	"context"
	"encoding/json"
	"fmt"

	"github.com/jinzhu/copier"
)

type Handler struct {
	servCaddy caddyInterface.CaddyService
}

func NewHandler(servCaddy caddyInterface.CaddyService) *Handler {
	return &Handler{
		servCaddy: servCaddy,
	}
}

// Caddy is the resolver for the Caddy field.
func (r *Handler) Caddy(ctx context.Context, input modelgraph.CaddyInput) (*modelgraph.Caddy, error) {
	record := modelA.Caddy{}
	err := copier.CopyWithOption(&record, &input, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return nil, err
	}
	{
		c, _ := json.Marshal(record)
		fmt.Println(string(c))
	}

	res, err := r.servCaddy.Create(ctx, &record)
	if err != nil {
		return nil, err
	}
	c := res.ToGraph()

	return &c, nil
}
