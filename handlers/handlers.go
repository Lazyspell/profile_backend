package handlers

import (
	"github.com/lazyspell/profile_backend/config"
	"github.com/lazyspell/profile_backend/driver"
	"github.com/lazyspell/profile_backend/repository"
	"github.com/lazyspell/profile_backend/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewMongoRepo(db.NoSql),
	}
}

func NewHandler(r *Repository) {
	Repo = r
}
