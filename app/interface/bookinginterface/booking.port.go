package bookinginterface

import (
	"cadigo-api/app/modela"
	"context"
)

type BookingRepository interface {
	Create(ctx context.Context, caddy *modela.Booking) (*modela.Booking, error)
	Update(ctx context.Context, argID string, record *modela.Booking) (*modela.Booking, error)
	GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.Booking, total int64, err error)
	GetByID(ctx context.Context, id string) (result *modela.Booking, err error)
}

type BookingService interface {
	Create(ctx context.Context, caddy *modela.Booking) (*modela.Booking, error)
	Update(ctx context.Context, argID string, record *modela.Booking) (*modela.Booking, error)
	GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.Booking, total int64, err error)
	GetByID(ctx context.Context, id string) (result *modela.Booking, err error)
}
