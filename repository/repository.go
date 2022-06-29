package repository

import "github.com/lazyspell/profile_backend/models"

type DatabaseRepo interface {
	//profiles
	AllProfilesDB() ([]models.Profile, error)
	InsertProfilesDB(profile models.Profile) (string, error)
}
