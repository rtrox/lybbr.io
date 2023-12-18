package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/graph/generated"

	"entgo.io/contrib/entgql"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id ksuid.ID) (ent.Noder, error) {
	return r.client.Noder(ctx, id, ent.WithNodeType(ent.IDToType))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []ksuid.ID) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids, ent.WithNodeType(ent.IDToType))
}

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context, after *entgql.Cursor[ksuid.ID], first *int, before *entgql.Cursor[ksuid.ID], last *int, orderBy []*ent.AuthorOrder, where *ent.AuthorWhereInput) (*ent.AuthorConnection, error) {
	return r.client.Author.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithAuthorOrder(orderBy),
		)
}

// Books is the resolver for the books field.
func (r *queryResolver) Books(ctx context.Context, after *entgql.Cursor[ksuid.ID], first *int, before *entgql.Cursor[ksuid.ID], last *int, orderBy []*ent.BookOrder, where *ent.BookWhereInput) (*ent.BookConnection, error) {
	return r.client.Book.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithBookOrder(orderBy),
		)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
