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

// GetCaddy is the resolver for the getCaddy field.
func (r *Handler) GetCaddy(ctx context.Context, input modelgraph.GetCaddyInput) (*modelgraph.Caddy, error) {
	panic(fmt.Errorf("not implemented: GetCaddy - getCaddy"))
}

// GetCaddys is the resolver for the getCaddys field.
func (r *Handler) GetCaddys(ctx context.Context, input modelgraph.GetCaddysInput) (data *modelgraph.CaddyData, err error) {
	defaultPagination := new(modelA.Pagination).Init()

	err = copier.Copy(&defaultPagination, &input.Pagination)
	if err != nil {
		return nil, err
	}

	res, total, err := r.servCaddy.GetAll(ctx, defaultPagination)
	if err != nil {
		return nil, err
	}

	caddyData := []*modelgraph.Caddy{}
	for _, v := range res {
		g := v.ToGraph()
		caddyData = append(caddyData, &g)
	}

	data = &modelgraph.CaddyData{
		Data: caddyData,
		Pagination: &modelgraph.PaginationType{
			Page:  int(defaultPagination.Page),
			Limit: int(defaultPagination.Limit),
			Total: int(total),
		},
	}

	return data, nil
}
