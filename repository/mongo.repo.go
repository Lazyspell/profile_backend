package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/lazyspell/profile_backend/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProfileRepository interface {
	Save(profile *model.ProfileQl)
	FindAll() []*model.ProfileQl
}

type database struct {
	client *mongo.Client
}

func New() ProfileRepository {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	clientOptions = clientOptions.SetMaxPoolSize(50)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return &database{
		client: dbClient,
	}
}

func (db *database) Save(profile *model.ProfileQl) {
	collection := db.client.Database("profiletest").Collection("profiletest")
	_, err := collection.InsertOne(context.TODO(), profile)
	if err != nil {
		log.Fatal(err)
	}

}

func (db *database) FindAll() []*model.ProfileQl {
	var results []*model.ProfileQl

	coll := db.client.Database("profiletest").Collection("profiletest")

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println(err)
		return results

	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results

}
