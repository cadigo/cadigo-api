package caddyhandler

import (
	"cadigo-api/app/interface/caddyinterface"
	"cadigo-api/app/interface/coursegolfinterface"
	"cadigo-api/app/modela"
	"cadigo-api/graph/modelgraph"
	"context"

	"github.com/jinzhu/copier"
)

type Handler struct {
	servCaddy      caddyinterface.CaddyService
	servCourseGolf coursegolfinterface.CourseGolfService
}

func NewHandler(servCaddy caddyinterface.CaddyService, servCourseGolf coursegolfinterface.CourseGolfService) *Handler {
	return &Handler{
		servCaddy:      servCaddy,
		servCourseGolf: servCourseGolf,
	}
}

// Caddy is the resolver for the Caddy field.
func (r *Handler) Caddy(ctx context.Context, input modelgraph.CaddyInput) (*modelgraph.Caddy, error) {
	record := modela.Caddy{}
	err := copier.CopyWithOption(&record, &input, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return nil, err
	}

	// Insert
	if input.ID == nil {
		res, err := r.servCaddy.Create(ctx, &record)
		if err != nil {
			return nil, err
		}
		c := res.ToGraph()

		return &c, nil
	} else {
		// Update
		res, err := r.servCaddy.Update(ctx, *input.ID, &record)
		if err != nil {
			return nil, err
		}
		c := res.ToGraph()

		return &c, nil
	}
}

// GetCaddy is the resolver for the getCaddy field.
func (r *Handler) GetCaddy(ctx context.Context, input modelgraph.GetCaddyInput) (*modelgraph.Caddy, error) {
	if input.ID != nil {
		d, err := r.servCaddy.GetByID(ctx, *input.ID)
		if err != nil {
			return nil, nil
		}

		g := d.ToGraph()

		return &g, nil
	}

	return nil, nil
}

// GetCaddys is the resolver for the getCaddys field.
func (r *Handler) GetCaddys(ctx context.Context, input modelgraph.GetCaddysInput) (data *modelgraph.CaddyData, err error) {
	defaultPagination := new(modela.Pagination).Init()

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

func (r *Handler) CourseGolf(ctx context.Context, obj *modelgraph.Caddy) ([]*modelgraph.CourseGolf, error) {
	r.servCourseGolf.GetByIDs(ctx, (obj).CourseGolfIDs)

	return nil, nil
}
