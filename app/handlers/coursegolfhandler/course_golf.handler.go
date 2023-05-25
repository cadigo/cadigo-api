package coursegolfhandler

import (
	"cadigo-api/app/interface/coursegolfinterface"
	"cadigo-api/app/modela"
	"cadigo-api/graph/modelgraph"
	"context"
	"fmt"

	"github.com/jinzhu/copier"
)

type Handler struct {
	serv coursegolfinterface.CourseGolfService
}

func NewHandler(serv coursegolfinterface.CourseGolfService) *Handler {
	return &Handler{
		serv: serv,
	}
}

// CourseGolf is the resolver for the courseGolf field.
func (r *Handler) CourseGolf(ctx context.Context, input modelgraph.CourseGolfInput) (*modelgraph.CourseGolf, error) {
	record := modela.CourseGolf{}
	err := copier.CopyWithOption(&record, &input, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return nil, err
	}

	// Insert
	if input.ID == nil {
		res, err := r.serv.Create(ctx, &record)
		if err != nil {
			return nil, err
		}
		c := res.ToGraph()

		return &c, nil
	} else {
		// Update
		res, err := r.serv.Update(ctx, *input.ID, &record)
		if err != nil {
			return nil, err
		}
		c := res.ToGraph()

		return &c, nil
	}
}

// DeleteCourseGolf is the resolver for the deleteCourseGolf field.
func (r *Handler) DeleteCourseGolf(ctx context.Context, id string) (*modelgraph.CourseGolf, error) {
	panic(fmt.Errorf("not implemented: DeleteCourseGolf - deleteCourseGolf"))
}

// GetCourseGolf is the resolver for the getCourseGolf field.
func (r *Handler) GetCourseGolf(ctx context.Context, input modelgraph.GetCourseGolfInput) (*modelgraph.CourseGolf, error) {
	if input.ID != nil {
		d, err := r.serv.GetByID(ctx, *input.ID)
		if err != nil {
			return nil, nil
		}

		g := d.ToGraph()

		return &g, nil
	}

	return nil, nil
}

// GetCourseGolfs is the resolver for the getCourseGolfs field.
func (r *Handler) GetCourseGolfs(ctx context.Context, input modelgraph.GetCourseGolfsInput) (data *modelgraph.CourseGolfData, err error) {
	defaultPagination := new(modela.Pagination).Init()

	err = copier.Copy(&defaultPagination, &input.Pagination)
	if err != nil {
		return nil, err
	}

	res, total, err := r.serv.GetAll(ctx, defaultPagination)
	if err != nil {
		return nil, err
	}

	courseGolfData := []*modelgraph.CourseGolf{}
	for _, v := range res {
		g := v.ToGraph()
		courseGolfData = append(courseGolfData, &g)
	}

	data = &modelgraph.CourseGolfData{
		Data: courseGolfData,
		Pagination: &modelgraph.PaginationType{
			Page:  int(defaultPagination.Page),
			Limit: int(defaultPagination.Limit),
			Total: int(total),
		},
	}

	return data, nil
}
