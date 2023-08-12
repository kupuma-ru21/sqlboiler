package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"sqlboiler/graph/model"
)

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.CreateLinkInput) (*model.Link, error) {
	return r.Srv.CreateLink(ctx, input)
}

// CreateLinks is the resolver for the createLinks field.
func (r *mutationResolver) CreateLinks(ctx context.Context, input []*model.CreateLinkInput) ([]*model.Link, error) {
	return r.Srv.CreateLinks(ctx, input)
}

// UpdateLink is the resolver for the updateLink field.
func (r *mutationResolver) UpdateLink(ctx context.Context, input model.UpdateLinkInput) (*model.Link, error) {
	return r.Srv.UpdateLink(ctx, input)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented: RefreshToken - refreshToken"))
}

// Links is the resolver for the links field.
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	return r.Srv.GetLinks(ctx)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, name string) (*model.User, error) {
	return r.Srv.GetUserByName(ctx, name)
}

// Link is the resolver for the link field.
func (r *queryResolver) Link(ctx context.Context, title string) (*model.Link, error) {
	return r.Srv.GetLinkByTitle(ctx, title)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
