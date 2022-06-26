package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lazyspell/profile_backend/config"
	"github.com/lazyspell/profile_backend/handlers"
)

func Routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Working)

	mux.Get("/testing", handlers.Repo.GetAllProfiles)

	return mux
}
