package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"cadigo-api/graph/modelgraph"
	"context"
)

// Customer is the resolver for the customer field.
func (r *mutationResolver) Customer(ctx context.Context, input modelgraph.CustomerInput) (*modelgraph.Customer, error) {
	return r.CustomerHandler.Customer(ctx, input)
}

// GetCustomer is the resolver for the getCustomer field.
func (r *queryResolver) GetCustomer(ctx context.Context, input modelgraph.GetCustomerInput) (*modelgraph.Customer, error) {
	return r.CustomerHandler.GetCustomer(ctx, input)
}
