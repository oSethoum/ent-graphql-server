package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"entql/ent"
	"entql/graph/generated"
)

// CreateFriendship is the resolver for the createFriendship field.
func (r *mutationResolver) CreateFriendship(ctx context.Context, input ent.CreateFriendshipInput) (*ent.Friendship, error) {
	return ent.FromContext(ctx).Friendship.Create().SetInput(input).Save(ctx)
}

// UpdateFriendship is the resolver for the updateFriendship field.
func (r *mutationResolver) UpdateFriendship(ctx context.Context, id int, input ent.UpdateFriendshipInput) (*ent.Friendship, error) {
	return ent.FromContext(ctx).Friendship.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteFriendship is the resolver for the deleteFriendship field.
func (r *mutationResolver) DeleteFriendship(ctx context.Context, id int) (*ent.Friendship, error) {
	friendship, _ := r.Client.Friendship.Get(ctx, id)
	return friendship, r.Client.Friendship.DeleteOneID(id).Exec(ctx)
}

// Friendship is the resolver for the Friendship field.
func (r *queryResolver) Friendship(ctx context.Context, id int) (*ent.Friendship, error) {
	return r.Client.Friendship.Get(ctx, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
