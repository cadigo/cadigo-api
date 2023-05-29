package coursegolfinterface

import (
	"cadigo-api/app/modela"
	"context"
)

type CourseGolfRepository interface {
	Create(ctx context.Context, caddy *modela.CourseGolf) (*modela.CourseGolf, error)
	Update(ctx context.Context, argID string, record *modela.CourseGolf) (*modela.CourseGolf, error)
	GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.CourseGolf, total int64, err error)
	GetByID(ctx context.Context, id string) (result *modela.CourseGolf, err error)
	GetByIDs(ctx context.Context, ids []string) (result *[]modela.CourseGolf, err error)
}

type CourseGolfService interface {
	Create(ctx context.Context, caddy *modela.CourseGolf) (*modela.CourseGolf, error)
	Update(ctx context.Context, argID string, record *modela.CourseGolf) (*modela.CourseGolf, error)
	GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.CourseGolf, total int64, err error)
	GetByID(ctx context.Context, id string) (result *modela.CourseGolf, err error)
	GetByIDs(ctx context.Context, ids []string) (result *[]modela.CourseGolf, err error)
}
