package customerservice

import (
	"cadigo-api/app/interface/customerinterface"
	"cadigo-api/app/modela"
	"context"
)

type Service struct {
	repo customerinterface.CustomerRepository
}

func NewService(repo customerinterface.CustomerRepository) customerinterface.CustomerService {
	return &Service{
		repo: repo,
	}
}

func (serv *Service) Create(ctx context.Context, caddy *modela.Customer) (*modela.Customer, error) {
	return serv.repo.Create(ctx, caddy)
}

func (serv *Service) GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.Customer, total int64, err error) {
	return serv.repo.GetAll(ctx, pagination)
}

func (serv *Service) Update(ctx context.Context, argID string, record *modela.Customer) (*modela.Customer, error) {
	return serv.repo.Update(ctx, argID, record)
}

func (serv *Service) GetByID(ctx context.Context, id string) (result *modela.Customer, err error) {
	return serv.repo.GetByID(ctx, id)
}
