package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"entql/ent"
	"entql/graph/generated"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.Client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.Client.Noders(ctx, ids)
}

// Friendships is the resolver for the friendships field.
func (r *queryResolver) Friendships(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.FriendshipOrder, where *ent.FriendshipWhereInput) (*ent.FriendshipConnection, error) {
	return r.Client.Friendship.Query().Paginate(ctx, after, first, before, last, ent.WithFriendshipOrder(orderBy), ent.WithFriendshipFilter(where.Filter))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.Client.User.Query().Paginate(ctx, after, first, before, last, ent.WithUserOrder(orderBy), ent.WithUserFilter(where.Filter))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
