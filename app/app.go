package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/lazyspell/profile_backend/config"
	"github.com/lazyspell/profile_backend/driver"
	"github.com/lazyspell/profile_backend/graph"
	"github.com/lazyspell/profile_backend/graph/generated"
	"github.com/lazyspell/profile_backend/handlers"
	"github.com/lazyspell/profile_backend/helpers"
)

const portNumber = ":8080"

const defaultPort = "8080"

var app config.AppConfig

func Start() {
	fmt.Println("Application has started")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer client.NoSql.Disconnect(ctx)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	fmt.Println("Connected to MongoDB!")
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	// fmt.Println(fmt.Sprintf("Starting application on http://localhost:8080"))

	// srv := &http.Server{
	// 	Addr:    portNumber,
	// 	Handler: r.Routes(&app),
	// }
	// err = srv.ListenAndServe()
	// log.Fatal(err)

}

func run() (*driver.DB, error) {
	log.Println("Connecting to database...")

	db, err := driver.ConnectMongoDB()
	if err != nil {
		log.Fatal("Can not connect to database! Dying...")
	}

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandler(repo)
	helpers.NewHelpers(&app)

	return db, nil

}
