package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/smtp"
	"os"

	"github.com/lazyspell/profile_backend/graph/generated"
	"github.com/lazyspell/profile_backend/graph/model"
)

func (r *mutationResolver) CreateProfile(ctx context.Context, input model.InputProfile) (*model.ProfileQl, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SendEmail(ctx context.Context, input model.InputExternalEmail) (string, error) {
	host := "smtp.mail.yahoo.com"
	port := "587"

	emailAddress := input.EmailAddress
	toList := []string{emailAddress}
	// emailName := input.Name
	emailSubject := "Subject: " + input.EmailSubject + "\n"
	emailMessage := "\n" + input.EmailName + "\n" + emailAddress + "\n" + input.EmailMessage
	fromEmail := os.Getenv("YAHOO_EMAIL_ADDRESS")
	passwordEmail := os.Getenv("YAHOO_PASSWORD")

	// from := "jeremy2975@yahoo.com"
	// password := "frrthgtadmasjcof"

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
