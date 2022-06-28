package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crypto/rand"
	"fmt"

	"github.com/lazyspell/profile_backend/graph/generated"
	"github.com/lazyspell/profile_backend/graph/model"
)

func (r *mutationResolver) CreateProfile(ctx context.Context, input model.NewProfile) (*model.Profile, error) {

	var names *model.Name
	names.FirstName = *input.FirstName
	names.LastName = *input.LastName

	var location *model.Location
	location.State = *input.State
	location.City = *input.City
	location.ZipCode = *input.ZipCode

	profiles := model.Profile{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Name:        names,
		Location:    location,
		Email:       *input.Email,
		Skills:      []*model.Skill{},
		Description: *input.Description,
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
