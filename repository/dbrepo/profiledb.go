package dbrepo

import (
	"context"
	"fmt"
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

	coll := m.DB.Database("profile_db").Collection("profile_collection")

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

	coll := m.DB.Database("profile_db").Collection("profile_collection")

	_, err := coll.InsertOne(context.TODO(), profile)
	if err != nil {
		log.Println(err)
		return "failed", err
	}

	return "success", nil

}

func (m *mongoDBRepo) FindProfileById(profileId string) (model.ProfileQl, error) {
	profileResults := model.ProfileQl{}
	call := m.DB.Database("profile_db").Collection("profile_collection")

	cursor := call.FindOne(context.TODO(), bson.M{"contact.email": profileId})

	cursor.Decode(&profileResults)

	return profileResults, nil
}

func (m *mongoDBRepo) UpdateSkillsQL(technologies model.Technologies, email string, category string) error {
	call := m.DB.Database("profile_db").Collection("profile_collection")

	query := bson.M{"contact.email": email}
	location := fmt.Sprint("skills.", category)

	update := bson.M{"$addToSet": bson.M{location: technologies}}

	_, err := call.UpdateOne(context.TODO(), query, update)
	if err != nil {
		fmt.Println("failed")
		fmt.Println(err)
		return err
	}
	return nil
}

func (m *mongoDBRepo) UpdateQuotes(quotes model.Quote, email string) error {
	call := m.DB.Database("profile_db").Collection("profile_collection")

	query := bson.M{"contact.email": email}

	update := bson.M{"$addToSet": bson.M{"quotes": quotes}}

	_, err := call.UpdateOne(context.TODO(), query, update)
	if err != nil {
		fmt.Println("failed")
		fmt.Println(err)
		return err
	}
	return nil
}

func (m *mongoDBRepo) UpdateProjectsQL(application model.Application, email string) error {
	call := m.DB.Database("profile_db").Collection("profile_collection")

	query := bson.M{"contact.email": email}
	update := bson.M{"$push": bson.M{"projects": application}}

	_, err := call.UpdateOne(context.TODO(), query, update)
	if err != nil {
		fmt.Println("failed")
		fmt.Println(err)
		return err
	}
	return nil
}

func (m *mongoDBRepo) UpdateJobQL(job model.Job, email string) error {
	call := m.DB.Database("profile_db").Collection("profile_collection")

	query := bson.M{"contact.email": email}
	update := bson.M{"$push": bson.M{"experience": job}}

	_, err := call.UpdateOne(context.TODO(), query, update)
	if err != nil {
		fmt.Println("failed")
		fmt.Println(err)
		return err
	}
	return nil
}
