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
}
