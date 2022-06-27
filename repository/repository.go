package repository

import "github.com/lazyspell/profile_backend/models"

type DatabaseRepo interface {
	AllProfiles() ([]models.Profile, error)
}
