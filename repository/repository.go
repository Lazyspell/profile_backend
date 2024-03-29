package repository

import (
	"github.com/lazyspell/profile_backend/graph/model"
	"github.com/lazyspell/profile_backend/models"
)

type DatabaseRepo interface {
	//profiles REST
	AllProfilesDB() ([]models.Profile, error)
	InsertProfilesDB(profile models.Profile) (string, error)

	//profiles Graphql
	AllProfilesQL() ([]*model.ProfileQl, error)
	InsertProfilesQL(profile model.ProfileQl) (string, error)
	FindProfileById(profileId string) (model.ProfileQl, error)
	UpdateSkillsQL(technologies model.Technologies, email string, category string) error
	UpdateProjectsQL(application model.Application, email string) error
	UpdateJobQL(job model.Job, email string) error
	UpdateQuotes(quotes model.Quote, email string) error
}
