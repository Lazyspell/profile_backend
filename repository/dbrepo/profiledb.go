package dbrepo

import (
	"context"
	"log"

	"github.com/lazyspell/profile_backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *mongoDBRepo) AllProfiles() ([]models.Profile, error) {

	var results []models.Profile

	coll := m.DB.Database("profiletest").Collection("profile")

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println(err)
		return results, err

	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, nil

}
