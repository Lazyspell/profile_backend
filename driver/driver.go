package driver

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	NoSql      *mongo.Client
	Connection *options.ClientOptions
}

var dbConn = &DB{}

const maxOpenDBConn = 10
const maxIdleDBCon = 5
const maxDBLifeTime = 5 * time.Minute

func ConnectMongoDB() (*DB, error) {
	// Set client options

	// clientOptions := NewDatabase("mongodb://localhost:27017")
	// connectionString := fmt.Sprintf(app.DB.ConnectionString)
	connectionString := os.Getenv("DIGITAL_OCEAN_MONGO_DB")
	clientOptions := NewDatabase(connectionString)

	clientOptions.SetMaxConnIdleTime(maxIdleDBCon)
	clientOptions.SetMaxConnecting(uint64(maxDBLifeTime))
	clientOptions.SetMaxConnecting(maxOpenDBConn)

	dbConn.Connection = clientOptions

	// Connect to MongoDB

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	dbConn.NoSql = client

	err = testDB(dbConn.NoSql)
	if err != nil {
		return nil, err
	}

	return dbConn, err
}

// Check the connection
func testDB(client *mongo.Client) error {
	err := client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func NewDatabase(dns string) *options.ClientOptions {
	db := options.Client().ApplyURI(dns)

	return db
}
