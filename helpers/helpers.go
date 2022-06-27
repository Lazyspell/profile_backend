package helpers

import "github.com/lazyspell/profile_backend/config"

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}
