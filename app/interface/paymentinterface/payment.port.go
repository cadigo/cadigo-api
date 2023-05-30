package paymentinterface

import (
	"cadigo-api/app/modela"
	"context"
)

type PaymentRepository interface {
	Create(ctx context.Context, record *modela.Payment) (*modela.Payment, error)
	Update(ctx context.Context, argID string, record *modela.Payment) (*modela.Payment, error)
	GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.Payment, total int64, err error)
	GetByID(ctx context.Context, id string) (result *modela.Payment, err error)
}

type PaymentService interface {
	Create(ctx context.Context, record *modela.PaymentRequest) (*modela.Payment, error)
	Update(ctx context.Context, argID string, record *modela.Payment) (*modela.Payment, error)
	GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.Payment, total int64, err error)
	GetByID(ctx context.Context, id string) (result *modela.Payment, err error)
}
