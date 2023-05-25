package paymentservice

import (
	"cadigo-api/app/interface/paymentinterface"
	"cadigo-api/app/modela"
	"context"
)

type Service struct {
	repo paymentinterface.PaymentRepository
}

func NewService(repo paymentinterface.PaymentRepository) paymentinterface.PaymentService {
	return &Service{
		repo: repo,
	}
}

func (serv *Service) Create(ctx context.Context, caddy *modela.Payment) (*modela.Payment, error) {
	return serv.repo.Create(ctx, caddy)
}

func (serv *Service) GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.Payment, total int64, err error) {
	return serv.repo.GetAll(ctx, pagination)
}

func (serv *Service) Update(ctx context.Context, argID string, record *modela.Payment) (*modela.Payment, error) {
	return serv.repo.Update(ctx, argID, record)
}

func (serv *Service) GetByID(ctx context.Context, id string) (result *modela.Payment, err error) {
	return serv.repo.GetByID(ctx, id)
}
