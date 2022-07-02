package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/lazyspell/profile_backend/graph/generated"
	"github.com/lazyspell/profile_backend/graph/model"
	"github.com/lazyspell/profile_backend/handlers"
)

func (r *mutationResolver) CreateProfile(ctx context.Context, input model.NewProfile) (*model.ProfileQl, error) {

	ProfileID := *input.FirstName + strconv.Itoa(rand.Int())

	profiles := &model.ProfileQl{
		ID:          ProfileID,
		Name:        &model.Name{FirstName: *input.FirstName, LastName: *input.LastName},
		Location:    &model.Location{State: *input.State, City: *input.City, ZipCode: *input.ZipCode},
		Email:       *input.Email,
		Skills:      []*model.Skill{},
		Description: *input.Description,
	}
	handlers.Repo.InsertProfilesQL(profiles)
	return profiles, nil
}

func (r *queryResolver) Profile(ctx context.Context) ([]*model.ProfileQl, error) {
	profiles, err := handlers.Repo.FindAll()
	if err != nil {
		return profiles, err
	}
	fmt.Println(profiles)
	return profiles, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
