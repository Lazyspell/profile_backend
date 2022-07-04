package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lazyspell/profile_backend/graph/model"
	"github.com/lazyspell/profile_backend/helpers"
	"github.com/lazyspell/profile_backend/models"
)

func (m *Repository) GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := m.DB.AllProfilesDB()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)

}

func (m *Repository) InsertProfiles(w http.ResponseWriter, r *http.Request) {
	var payload models.Profile
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		helpers.BadRequest400(w, "invalid type please check request body")
		return
	}

	_, err = m.DB.InsertProfilesDB(payload)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

}

func (m *Repository) GetAllProfilesQL() ([]*model.ProfileQl, error) {
	var results []*model.ProfileQl

	results, err := m.DB.AllProfilesQL()
	if err != nil {
		log.Fatal(err)
	}

	return results, nil
}

func (m *Repository) InsertProfilesQL(newProfile *model.ProfileQl) {
	_, err := m.DB.InsertProfilesQL(*newProfile)
	if err != nil {
		log.Fatal(err)
	}

}

func (m *Repository) GetProfileByIdQL(profileId string) (model.ProfileQl, error) {

	results, err := m.DB.FindProfileById(profileId)
	if err != nil {
		log.Fatal(err)
	}

	return results, err

}
