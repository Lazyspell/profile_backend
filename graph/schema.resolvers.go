package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/lazyspell/profile_backend/graph/generated"
	"github.com/lazyspell/profile_backend/graph/model"
	"github.com/lazyspell/profile_backend/handlers"
	"github.com/lazyspell/profile_backend/helpers"
)

func (r *mutationResolver) CreateProfile(ctx context.Context, input model.InputProfile) (*model.ProfileQl, error) {
	newCategories := &model.Categories{
		Frontend:        []*model.Technologies{},
		Backend:         []*model.Technologies{},
		MachineLearning: []*model.Technologies{},
		CloudService:    []*model.Technologies{},
	}

	newProfile := &model.ProfileQl{
		FirstName:  *input.FirstName,
		LastName:   *input.LastName,
		Dob:        (*model.DateOfBirth)(input.Dob),
		Location:   (*model.Location)(input.Location),
		Contact:    (*model.ContactInfo)(input.Contact),
		Skills:     newCategories,
		Projects:   []*model.Application{},
		Experience: []*model.Job{},
	}

	err := handlers.Repo.InsertProfilesQL(newProfile)
	if err != nil {
		log.Fatal(err)
	}
	return newProfile, nil
}

func (r *mutationResolver) SendEmail(ctx context.Context, input model.InputExternalEmail) (string, error) {
	host := "smtp.mail.yahoo.com"
	port := "587"

	toList := []string{os.Getenv("RECEIVING_EMAIL_ADDRESS")}

	emailAddress := input.EmailAddress
	if !helpers.ValidMailAddress(emailAddress) {
		return "Failed", nil
	}

	emailSubject := "Subject: " + input.EmailSubject + "\n"
	emailMessage := "\n" + input.EmailName + "\n" + emailAddress + "\n" + input.EmailMessage
	fromEmail := os.Getenv("YAHOO_EMAIL_ADDRESS")
	passwordEmail := os.Getenv("YAHOO_PASSWORD")

	message := []byte(emailSubject + emailMessage)

	auth := smtp.PlainAuth("", fromEmail, passwordEmail, host)
	err := smtp.SendMail(host+":"+port, auth, fromEmail, toList, message)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
		return "Failed", err
	}

	return "Success", err
}

func (r *mutationResolver) UpdateSkills(ctx context.Context, input model.InputTechnologies, category string, email string) (*model.ProfileQl, error) {
	inputSkills := &model.Technologies{
		TechName:          input.TechName,
		TechLink:          input.TechLink,
		ImageURL:          input.ImageLink,
		YearsOfExperience: input.YearsOfExperience,
		TechDescription:   input.TechDescription,
		Category:          input.Category,
	}

	results, err := handlers.Repo.UpdateSkills(*inputSkills, email, category)
	if err != nil {
		return &results, err
	}
	return &results, err
}

func (r *mutationResolver) UpdateProjects(ctx context.Context, input model.InputApplication, email string) (*model.ProfileQl, error) {
	inputProjects := &model.Application{
		ProjectName:         input.ProjectName,
		FrontendLink:        input.FrontendLink,
		FrontendDescription: input.FrontendDescription,
		BackendLink:         input.BackendLink,
		BackendDescription:  input.BackendDescription,
	}

	results, err := handlers.Repo.UpdateProjects(*inputProjects, email)
	if err != nil {
		return &results, err
	}
	return &results, err
}

func (r *mutationResolver) UpdateJob(ctx context.Context, input model.InputJob, email string) (*model.ProfileQl, error) {
	inputJobs := &model.Job{
		CompanyName:     input.CompanyName,
		WorkDescription: input.WorkDescription,
		YearsWorked:     input.YearsWorked,
		TechUsed:        input.TechUsed,
	}

	results, err := handlers.Repo.UpdateJob(*inputJobs, email)
	if err != nil {
		return &results, err
	}
	return &results, err
}

func (r *queryResolver) ProfileID(ctx context.Context, email string) (*model.ProfileQl, error) {
	profile, err := handlers.Repo.GetProfileByIdQL(email)
	if err != nil {
		log.Fatal(err)
	}

	return &profile, err
}

func (r *queryResolver) Profile(ctx context.Context) ([]*model.ProfileQl, error) {
	profiles, err := handlers.Repo.GetAllProfilesQL()
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
