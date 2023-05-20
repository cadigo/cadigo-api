package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	modelgraph1 "cadigo-api/graph/modelgraph"
	"context"
	"fmt"
)

// Booking is the resolver for the booking field.
func (r *mutationResolver) Booking(ctx context.Context, input modelgraph1.BookingInput) (*modelgraph1.Booking, error) {
	panic(fmt.Errorf("not implemented: Booking - booking"))
}

// GetBooking is the resolver for the getBooking field.
func (r *queryResolver) GetBooking(ctx context.Context, input modelgraph1.GetBookingInput) (*modelgraph1.Booking, error) {
	panic(fmt.Errorf("not implemented: GetBooking - getBooking"))
}

// GetBookingsHistory is the resolver for the getBookingsHistory field.
func (r *queryResolver) GetBookingsHistory(ctx context.Context, input modelgraph1.BookingsHistoryInput) (*modelgraph1.BookingData, error) {
	panic(fmt.Errorf("not implemented: GetBookingsHistory - getBookingsHistory"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
