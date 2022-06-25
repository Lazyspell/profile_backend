package dbrepo

import (
	"github.com/lazyspell/profile_backend/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDBRepo struct {
	DB *mongo.Client
}

func NewMongoRepo(conn *mongo.Client) repository.DatabaseRepo {
	return &mongoDBRepo{
		DB: conn,
	}
}
