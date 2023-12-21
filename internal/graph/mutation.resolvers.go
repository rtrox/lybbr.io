package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"lybbrio/internal/ent"
	"lybbrio/internal/ent/schema/ksuid"
	"lybbrio/internal/graph/generated"
)

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, input ent.CreateBookInput) (*ent.Book, error) {
	client := ent.FromContext(ctx)
	return client.Book.Create().SetInput(input).Save(ctx)
}

// UpdateBook is the resolver for the updateBook field.
func (r *mutationResolver) UpdateBook(ctx context.Context, id ksuid.ID, input ent.UpdateBookInput) (*ent.Book, error) {
	client := ent.FromContext(ctx)
	return client.Book.UpdateOneID(id).SetInput(input).Save(ctx)
}

// CreateAuthor is the resolver for the createAuthor field.
func (r *mutationResolver) CreateAuthor(ctx context.Context, input ent.CreateAuthorInput) (*ent.Author, error) {
	client := ent.FromContext(ctx)
	return client.Author.Create().SetInput(input).Save(ctx)
}

// UpdateAuthor is the resolver for the updateAuthor field.
func (r *mutationResolver) UpdateAuthor(ctx context.Context, id ksuid.ID, input ent.UpdateAuthorInput) (*ent.Author, error) {
	client := ent.FromContext(ctx)
	return client.Author.UpdateOneID(id).SetInput(input).Save(ctx)
}

// CreateShelf is the resolver for the createShelf field.
func (r *mutationResolver) CreateShelf(ctx context.Context, input ent.CreateShelfInput) (*ent.Shelf, error) {
	client := ent.FromContext(ctx)
	return client.Shelf.Create().SetInput(input).Save(ctx)
}

// UpdateShelf is the resolver for the updateShelf field.
func (r *mutationResolver) UpdateShelf(ctx context.Context, id ksuid.ID, input ent.UpdateShelfInput) (*ent.Shelf, error) {
	client := ent.FromContext(ctx)
	return client.Shelf.UpdateOneID(id).SetInput(input).Save(ctx)
}

// CreateTag is the resolver for the createTag field.
func (r *mutationResolver) CreateTag(ctx context.Context, input ent.CreateTagInput) (*ent.Tag, error) {
	client := ent.FromContext(ctx)
	return client.Tag.Create().SetInput(input).Save(ctx)
}

// UpdateTag is the resolver for the updateTag field.
func (r *mutationResolver) UpdateTag(ctx context.Context, id ksuid.ID, input ent.UpdateTagInput) (*ent.Tag, error) {
	client := ent.FromContext(ctx)
	return client.Tag.UpdateOneID(id).SetInput(input).Save(ctx)
}

// CreatePublisher is the resolver for the createPublisher field.
func (r *mutationResolver) CreatePublisher(ctx context.Context, input ent.CreatePublisherInput) (*ent.Publisher, error) {
	client := ent.FromContext(ctx)
	return client.Publisher.Create().SetInput(input).Save(ctx)
}

// UpdatePublisher is the resolver for the updatePublisher field.
func (r *mutationResolver) UpdatePublisher(ctx context.Context, id ksuid.ID, input ent.UpdatePublisherInput) (*ent.Publisher, error) {
	client := ent.FromContext(ctx)
	return client.Publisher.UpdateOneID(id).SetInput(input).Save(ctx)
}

// CreateLanguage is the resolver for the createLanguage field.
func (r *mutationResolver) CreateLanguage(ctx context.Context, input ent.CreateLanguageInput) (*ent.Language, error) {
	client := ent.FromContext(ctx)
	return client.Language.Create().SetInput(input).Save(ctx)
}

// UpdateLanguage is the resolver for the updateLanguage field.
func (r *mutationResolver) UpdateLanguage(ctx context.Context, id ksuid.ID, input ent.UpdateLanguageInput) (*ent.Language, error) {
	client := ent.FromContext(ctx)
	return client.Language.UpdateOneID(id).SetInput(input).Save(ctx)
}

// CreateSeries is the resolver for the createSeries field.
func (r *mutationResolver) CreateSeries(ctx context.Context, input ent.CreateSeriesInput) (*ent.Series, error) {
	client := ent.FromContext(ctx)
	return client.Series.Create().SetInput(input).Save(ctx)
}

// UpdateSeries is the resolver for the updateSeries field.
func (r *mutationResolver) UpdateSeries(ctx context.Context, id ksuid.ID, input ent.UpdateSeriesInput) (*ent.Series, error) {
	client := ent.FromContext(ctx)
	return client.Series.UpdateOneID(id).SetInput(input).Save(ctx)
}

// CreateIdentifier is the resolver for the createIdentifier field.
func (r *mutationResolver) CreateIdentifier(ctx context.Context, input ent.CreateIdentifierInput) (*ent.Identifier, error) {
	client := ent.FromContext(ctx)
	return client.Identifier.Create().SetInput(input).Save(ctx)
}

// UpdateIdentifier is the resolver for the updateIdentifier field.
func (r *mutationResolver) UpdateIdentifier(ctx context.Context, id ksuid.ID, input ent.UpdateIdentifierInput) (*ent.Identifier, error) {
	client := ent.FromContext(ctx)
	return client.Identifier.UpdateOneID(id).SetInput(input).Save(ctx)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	client := ent.FromContext(ctx)
	return client.User.Create().SetInput(input).Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id ksuid.ID, input ent.UpdateUserInput) (*ent.User, error) {
	client := ent.FromContext(ctx)
	return client.User.UpdateOneID(id).SetInput(input).Save(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
