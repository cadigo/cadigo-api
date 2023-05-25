package coursegolfservice

import (
	"cadigo-api/app/interface/coursegolfinterface"
	"cadigo-api/app/modela"
	"context"
)

type Service struct {
	repo coursegolfinterface.CourseGolfRepository
}

func NewService(repo coursegolfinterface.CourseGolfRepository) coursegolfinterface.CourseGolfService {
	return &Service{
		repo: repo,
	}
}

func (serv *Service) Create(ctx context.Context, caddy *modela.CourseGolf) (*modela.CourseGolf, error) {
	return serv.repo.Create(ctx, caddy)
}

func (serv *Service) GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.CourseGolf, total int64, err error) {
	return serv.repo.GetAll(ctx, pagination)
}

func (serv *Service) Update(ctx context.Context, argID string, record *modela.CourseGolf) (*modela.CourseGolf, error) {
	return serv.repo.Update(ctx, argID, record)
}

func (serv *Service) GetByID(ctx context.Context, id string) (result *modela.CourseGolf, err error) {
	return serv.repo.GetByID(ctx, id)
}
