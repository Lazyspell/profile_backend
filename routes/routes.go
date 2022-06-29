package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lazyspell/profile_backend/config"
	"github.com/lazyspell/profile_backend/handlers"
)

func Routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Post("/profiles/new", handlers.Repo.InsertProfiles)

	mux.Get("/profiles/all", handlers.Repo.GetAllProfiles)

	return mux
}
