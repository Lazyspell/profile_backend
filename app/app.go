package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lazyspell/profile_backend/config"
	"github.com/lazyspell/profile_backend/driver"
	"github.com/lazyspell/profile_backend/handlers"
	r "github.com/lazyspell/profile_backend/routes"
)

const portNumber = ":8080"

var app config.AppConfig

func Start() {
	fmt.Println("Application has started")
	_, err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Starting application on http://localhost:8080"))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: r.Routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}

func run() (*driver.DB, error) {
	log.Println("Connecting to database...")

	db, err := driver.ConnectMongoDB()
	if err != nil {
		log.Fatal("Can not connect to database! Dying...")
	}

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandler(repo)

	return db, nil

}
