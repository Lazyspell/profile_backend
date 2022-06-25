package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lazyspell/profile_backend/driver"
	"github.com/lazyspell/profile_backend/handlers"
)

const portNumber = ":8080"

func Start() {
	fmt.Println("Application has started")
	_, err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("Starting application on http://localhost:8080"))

	srv := &http.Server{
		Addr: portNumber,
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

	repo := handlers.NewRepo(db)
	handlers.NewHandler(repo)

	return db, nil

}
