package customerhandler

import (
	"cadigo-api/app/interface/customerinterface"
	"cadigo-api/app/modela"
	"cadigo-api/graph/modelgraph"
	"context"

	"github.com/jinzhu/copier"
)

type Handler struct {
	serv customerinterface.CustomerService
}

func NewHandler(serv customerinterface.CustomerService) *Handler {
	return &Handler{
		serv: serv,
	}
}

// Customer is the resolver for the customer field.
func (r *Handler) Customer(ctx context.Context, input modelgraph.CustomerInput) (*modelgraph.Customer, error) {
	record := modela.Customer{}
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

// GetCustomer is the resolver for the getCustomer field.
func (r *Handler) GetCustomer(ctx context.Context, input modelgraph.GetCustomerInput) (*modelgraph.Customer, error) {
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
