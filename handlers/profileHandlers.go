package handlers

import (
	"fmt"
	"net/http"
)

func (m *Repository) GetAllProfiles(w http.ResponseWriter, r *http.Request) {
	m.DB.AllProfiles()

}

func (m *Repository) Working(w http.ResponseWriter, r *http.Request) {

	fmt.Println("working")
}
