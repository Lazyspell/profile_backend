package models

type Profile struct {
	Id   int `json:"id"`
	Name struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"full_name"`

	Location struct {
		State string `json:"state"`
		City  string `json:"city"`
		Zip   int    `json:"zip_code"`
	} `json:"location"`

	Email string `json:"email"`

	Skills []Skill `json:"skills_list"`

	Description string `json:"career_description"`
}

type Skill struct {
	Id          int    `json:"id"`
	Name        string `json:"string_name"`
	Description string `json:"skill_description"`
}
