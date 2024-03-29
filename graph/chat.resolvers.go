package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"cadigo-api/graph/modelgraph"
	"context"
)

// PostMessage is the resolver for the postMessage field.
func (r *mutationResolver) PostMessage(ctx context.Context, input *modelgraph.PostMessageInput) (*modelgraph.Message, error) {
	return r.ChatHandler.PostMessage(ctx, input)
}

// GetMessages is the resolver for the getMessages field.
func (r *queryResolver) GetMessages(ctx context.Context, input modelgraph.GetMessagesInput) (*modelgraph.GetMessagesType, error) {
	return r.ChatHandler.GetMessages(ctx, input)
}

// GetOnline is the resolver for the getOnline field.
func (r *queryResolver) GetOnline(ctx context.Context, input modelgraph.GetOnlineInput) ([]*modelgraph.Online, error) {
	return r.ChatHandler.GetOnline(ctx, input)
}

// Chat is the resolver for the chat field.
func (r *subscriptionResolver) Chat(ctx context.Context, input modelgraph.ChatInput) (<-chan *modelgraph.Message, error) {
	return r.ChatHandler.SubscriptionChat(ctx, input)
}

// Online is the resolver for the online field.
func (r *subscriptionResolver) Online(ctx context.Context, input modelgraph.OnlineInput) (<-chan string, error) {
	return r.ChatHandler.SubscriptionOnline(ctx, input)
}

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
