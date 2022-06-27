package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lazyspell/profile_backend/helpers"
)

func (m *Repository) GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := m.DB.AllProfiles()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)

}

func (m *Repository) Working(w http.ResponseWriter, r *http.Request) {

	fmt.Println("working")
}
