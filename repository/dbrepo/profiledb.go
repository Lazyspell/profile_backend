package dbrepo

import (
	"context"
	"log"

	"github.com/lazyspell/profile_backend/graph/model"
	"github.com/lazyspell/profile_backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *mongoDBRepo) AllProfilesDB() ([]models.Profile, error) {

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

func (m *mongoDBRepo) InsertProfilesDB(profile models.Profile) (string, error) {

	coll := m.DB.Database("profiletest").Collection("profile")

	_, err := coll.InsertOne(context.TODO(), profile)
	if err != nil {
		log.Println(err)
		return "failed", err
	}

	return "success", nil

}

func (m *mongoDBRepo) AllProfilesQL() ([]*model.ProfileQl, error) {

	var results []*model.ProfileQl

	coll := m.DB.Database("profiletest").Collection("profiletest")

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

func (m *mongoDBRepo) InsertProfilesQL(profile model.ProfileQl) (string, error) {

	coll := m.DB.Database("profiletest").Collection("profiletest")

	_, err := coll.InsertOne(context.TODO(), profile)
	if err != nil {
		log.Println(err)
		return "failed", err
	}

	return "success", nil

}
