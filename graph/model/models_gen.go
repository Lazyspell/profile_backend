// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Application struct {
	ProjectName         string `json:"project_name" bson:"project_name"`
	FrontendLink        string `json:"frontend_link" bson:"frontend_link"`
	FrontendDescription string `json:"frontend_description" bson:"frontend_description"`
	BackendLink         string `json:"backend_link" bson:"backend_link"`
	BackendDescription  string `json:"backend_description" bson:"backend_description"`
}

type Categories struct {
	Frontend        []*Technologies `json:"frontend" bson:"frontend"`
	Backend         []*Technologies `json:"backend" bson:"backend"`
	MachineLearning []*Technologies `json:"machine_learning" bson:"machine_learning"`
	CloudService    []*Technologies `json:"cloud_service" bson:"cloud_service"`
}

type ContactInfo struct {
	PhoneNumber   string `json:"phone_number" bson:"phone_number"`
	Email         string `json:"email" bson:"email"`
	Github        string `json:"github" bson:"github"`
	Linkedin      string `json:"linkedin" bson:"linkedin"`
	Discord       string `json:"discord" bson:"discord"`
	AboutMe       string `json:"about_me" bson:"about_me"`
	AboutMyCareer string `json:"about_my_career" bson:"about_my_career"`
}

type DateOfBirth struct {
	Day   int    `json:"day" bson:"day"`
	Month string `json:"month" bson:"month"`
	Year  int    `json:"year" bson:"year"`
	Age   int    `json:"age" bson:"age"`
}

type ExternalEmail struct {
	EmailName    string `json:"email_name" bson:"email_name"`
	EmailAddress string `json:"email_address" bson:"email_address"`
	EmailSubject string `json:"email_subject" bson:"email_subject"`
	EmailMessage string `json:"email_message" bson:"email_message"`
}

type InputApplication struct {
	ProjectName         string `json:"project_name" bson:"project_name"`
	FrontendLink        string `json:"frontend_link" bson:"frontend_link"`
	FrontendDescription string `json:"frontend_description" bson:"frontend_description"`
	BackendLink         string `json:"backend_link" bson:"backend_link"`
	BackendDescription  string `json:"backend_description" bson:"backend_description"`
}

type InputCategories struct {
	Frontend        []*InputTechnologies `json:"frontend" bson:"frontend"`
	Backend         []*InputTechnologies `json:"backend" bson:"backend"`
	MachineLearning []*InputTechnologies `json:"machine_learning" bson:"machine_learning"`
	CloudService    []*InputTechnologies `json:"cloud_service" bson:"cloud_service"`
}

type InputContactInfo struct {
	PhoneNumber   string `json:"phone_number" bson:"phone_number"`
	Email         string `json:"email" bson:"email"`
	Github        string `json:"github" bson:"github"`
	Linkedin      string `json:"linkedin" bson:"linkedin"`
	Discord       string `json:"discord" bson:"discord"`
	AboutMe       string `json:"about_me" bson:"about_me"`
	AboutMyCareer string `json:"about_my_career" bson:"about_my_career"`
}

type InputDateOfBirth struct {
	Day   int    `json:"day" bson:"day"`
	Month string `json:"month" bson:"month"`
	Year  int    `json:"year" bson:"year"`
	Age   int    `json:"age" bson:"age"`
}

type InputExternalEmail struct {
	EmailName    string `json:"email_name" bson:"email_name"`
	EmailAddress string `json:"email_address" bson:"email_address"`
	EmailSubject string `json:"email_subject" bson:"email_subject"`
	EmailMessage string `json:"email_message" bson:"email_message"`
}

type InputJob struct {
	CompanyName     string   `json:"company_name" bson:"company_name"`
	WorkDescription string   `json:"work_description" bson:"work_description"`
	YearsWorked     int      `json:"years_worked" bson:"years_worked"`
	TechUsed        []string `json:"tech_used" bson:"tech_used"`
}

type InputLocation struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	ZipCode int    `json:"zip_code" bson:"zip_code"`
}

type InputProfile struct {
	FirstName  *string             `json:"first_name" bson:"first_name"`
	LastName   *string             `json:"last_name" bson:"last_name"`
	Location   *InputLocation      `json:"location" bson:"location"`
	Skills     *InputCategories    `json:"skills" bson:"skills"`
	Dob        *InputDateOfBirth   `json:"dob" bson:"dob"`
	Projects   []*InputApplication `json:"projects" bson:"projects"`
	Contact    *InputContactInfo   `json:"contact" bson:"contact"`
	Experience []*InputJob         `json:"experience" bson:"experience"`
	Quotes     []*InputQuote       `json:"quotes" bson:"quotes"`
}

type InputTechnologies struct {
	TechName          string `json:"tech_name" bson:"tech_name"`
	TechLink          string `json:"tech_link" bson:"tech_link"`
	ImageLink         string `json:"image_link" bson:"image_link"`
	YearsOfExperience int    `json:"years_of_experience" bson:"years_of_experience"`
	TechDescription   string `json:"tech_description" bson:"tech_description"`
	Category 		  string `json:"category" bson:"category"`
	CategoryDescription string `json:"category_description" bson:"category_description"`

}

type InputQuote struct {
	Saying     string `json:"saying" bson:"saying"`
	Source     string `json:"source" bson:"source"`
	SourceLink string `json:"source_link" bson:"source_link"`
}

type Job struct {
	CompanyName     string   `json:"company_name" bson:"company_name"`
	WorkDescription string   `json:"work_description" bson:"work_description"`
	YearsWorked     int      `json:"years_worked" bson:"years_worked"`
	TechUsed        []string `json:"tech_used" bson:"tech_used"`
}

type Location struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"city"`
	ZipCode int    `json:"zip_code" bson:"zip_code"`
}

type ProfileQl struct {
	FirstName  string         `json:"first_name" bson:"first_name"`
	LastName   string         `json:"last_name" bson:"last_name"`
	Dob        *DateOfBirth   `json:"dob" bson:"dob"`
	Location   *Location      `json:"location" bson:"location"`
	Skills     *Categories    `json:"skills" bson:"skills"`
	Projects   []*Application `json:"projects" bson:"projects"`
	Contact    *ContactInfo   `json:"contact" bson:"contact"`
	Experience []*Job         `json:"experience" bson:"experience"`
	Quotes     []*Quote  	  `json:"quotes" bson:"quotes"`
}

type Technologies struct {
	TechName          string `json:"tech_name" bson:"tech_name"`
	TechLink          string `json:"tech_link" bson:"tech_link"`
	ImageURL          string `json:"image_url" bson:"image_url"`
	YearsOfExperience int    `json:"years_of_experience" bson:"years_of_experience"`
	TechDescription   string `json:"tech_description" bson:"tech_description"`
	Category 		  string `json:"category" bson:"category"`
	CategoryDescription string `json:"category_description" bson:"category_description"`

}

type Quote struct {
	Saying     string `json:"saying" bson:"saying"`
	Source     string `json:"source" bson:"source"`
	SourceLink string `json:"source_link" bson:"source_link"`
}