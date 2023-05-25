package customerinterface

import (
	"cadigo-api/app/modela"
	"context"
)

type CustomerRepository interface {
	Create(ctx context.Context, caddy *modela.Customer) (*modela.Customer, error)
	Update(ctx context.Context, argID string, record *modela.Customer) (*modela.Customer, error)
	GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.Customer, total int64, err error)
	GetByID(ctx context.Context, id string) (result *modela.Customer, err error)
}

type CustomerService interface {
	Create(ctx context.Context, caddy *modela.Customer) (*modela.Customer, error)
	Update(ctx context.Context, argID string, record *modela.Customer) (*modela.Customer, error)
	GetAll(ctx context.Context, pagination modela.Pagination) (result []*modela.Customer, total int64, err error)
	GetByID(ctx context.Context, id string) (result *modela.Customer, err error)
}
