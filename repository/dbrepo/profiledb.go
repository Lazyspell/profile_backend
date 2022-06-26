package dbrepo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *mongoDBRepo) AllProfiles() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	coll := m.DB.Database("profiletest").Collection("profile")

	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
	}

	var results []bson.D
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	for _, result := range results {
		fmt.Println(result)
	}

	fmt.Println(results)

}
