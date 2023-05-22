package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"cadigo-api/graph/modelgraph"
	"context"
)

// Payment is the resolver for the payment field.
func (r *mutationResolver) Payment(ctx context.Context, input modelgraph.PaymentInput) (*modelgraph.Payment, error) {
	return r.PaymentHandler.Payment(ctx, input)
}

// GetPayment is the resolver for the getPayment field.
func (r *queryResolver) GetPayment(ctx context.Context, input modelgraph.GetPaymentInput) (*modelgraph.Payment, error) {
	return r.PaymentHandler.GetPayment(ctx, input)
}
