package handlers

import (
	"github.com/lazyspell/profile_backend/driver"
	"github.com/lazyspell/profile_backend/repository"
	"github.com/lazyspell/profile_backend/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	DB repository.DatabaseRepo
}

func NewRepo(db *driver.DB) *Repository {
	return &Repository{
		DB: dbrepo.NewMongoRepo(db.NoSql),
	}
}

func NewHandler(r *Repository) {
	Repo = r
}
