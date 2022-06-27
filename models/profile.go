package models

type Profile struct {
	Id   int `json:"id" bson:"id,omitempty"`
	Name struct {
		FirstName string `json:"first_name" bson:"first_name"`
		LastName  string `json:"last_name" bson:"last_name"`
	} `json:"full_name" bson:"full_name"`
	Location struct {
		State string `json:"state" bson:"state"`
		City  string `json:"city" bson:"city"`
		Zip   int    `json:"zip_code" bson:"zip_code"`
	} `json:"location" bson:"location"`

	Email string `json:"email" bson:"email"`

	Skills []Skill `json:"skills_list" bson:"skills_list"`

	Description string `json:"career_description" bson:"career_description"`
}

type Skill struct {
	Id          int    `json:"id" bson:"id"`
	Name        string `json:"string_name" bson:"string_name"`
	Description string `json:"skill_description" bson:"skill_description"`
}
