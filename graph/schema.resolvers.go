package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lazyspell/profile_backend/graph/generated"
	"github.com/lazyspell/profile_backend/graph/model"
)

func (r *mutationResolver) CreateProfile(ctx context.Context, input model.InputProfile) (*model.ProfileQl, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SendEmail(ctx context.Context, input model.InputExternalEmail) (string, error) {
	emailAddress := input.EmailAddress
	emailName := input.Name
	emailSubject := input.EmailSubject
	emailMessage := input.EmailMessage
}

func (r *queryResolver) ProfileID(ctx context.Context, email string) (*model.ProfileQl, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Profile(ctx context.Context) ([]*model.ProfileQl, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
