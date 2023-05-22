package bookingservice

import (
	"cadigo-api/app/interface/bookinginterface"
	"cadigo-api/app/modela"
	"context"
)

type Service struct {
	repo bookinginterface.BookingRepository
}

func NewService(repo bookinginterface.BookingRepository) bookinginterface.BookingService {
	return &Service{
		repo: repo,
	}
}

func (serv *Service) Create(ctx context.Context, caddy *modela.Booking) (*modela.Booking, error) {
	return serv.repo.Create(ctx, caddy)
}

func (serv *Service) GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.Booking, total int64, err error) {
	return serv.repo.GetAll(ctx, pagination)
}

func (serv *Service) Update(ctx context.Context, argID string, record *modela.Booking) (*modela.Booking, error) {
	return serv.repo.Update(ctx, argID, record)
}

func (serv *Service) GetByID(ctx context.Context, id string) (result *modela.Booking, err error) {
	return serv.repo.GetByID(ctx, id)
}
