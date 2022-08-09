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

func (m *Repository) InsertProfilesQL(newProfile *model.ProfileQl) error {
	_, err := m.DB.InsertProfilesQL(*newProfile)
	if err != nil {
		return err
	}

	return nil

}

func (m *Repository) GetProfileByIdQL(profileId string) (model.ProfileQl, error) {

	results, err := m.DB.FindProfileById(profileId)
	if err != nil {
		log.Fatal(err)
	}

	return results, err

}

func (m *Repository) UpdateSkills(technology model.Technologies, email string, category string) (model.ProfileQl, error) {

	err := m.DB.UpdateSkillsQL(technology, email, category)
	if err != nil {
		log.Fatal(err)
	}

	results, err := m.DB.FindProfileById(email)
	if err != nil {
		log.Fatal(err)
	}

	return results, err

}

func (m *Repository) UpdateProjects(application model.Application, email string) (model.ProfileQl, error) {

	err := m.DB.UpdateProjectsQL(application, email)
	if err != nil {
		log.Fatal(err)
	}

	results, err := m.DB.FindProfileById(email)
	if err != nil {
		log.Fatal(err)
	}

	return results, err

}

func (m *Repository) UpdateJob(job model.Job, email string) (model.ProfileQl, error) {

	err := m.DB.UpdateJobQL(job, email)
	if err != nil {
		log.Fatal(err)
	}

	results, err := m.DB.FindProfileById(email)
	if err != nil {
		log.Fatal(err)
	}

	return results, err

}
