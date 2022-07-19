package models

type Application struct {
	ProjectName         string `json:"project_name" bson:"project_name"`
	FrontendLink        string `json:"frontend_link" bson:"frontend_link"`
	FrontendDescription string `json:"frontend_description" bson:"frontend_description"`
	BackendLink         string `json:"backend_link" bson:"backend_link"`
	BackendDescription  string `json:"backend_description" bson:"backend_description"`
}

type ContactInfo struct {
	PhoneNumber   string `json:"phone_number" bson:"phone_number"`
	Email         string `json:"email" bson:"email"`
	Github        string `json:"github" bson:"github"`
	LinkedIn      string `json:"linkedin" bson:"linkedin"`
	Discord       string `json:"discord" bson:"discord"`
	AboutMe       string `json:"about_me" bson:"about_me"`
	AboutMyCareer string `json:"about_my_career" bson:"about_my_career"`
}
type DateOfBirth struct {
	Day   int    `json:"day" bson:"day"`
	Month string `json:"month" bson:"month"`
	Year  int    `json:"year" bson:"month"`
	Age   int    `json:"age" bson:"age"`
}

type Profile struct {
	FirstName  string         `json:"first_name" bson:"first_name"`
	LastName   string         `json:"last_name" bson:"last_name"`
	DOB        DateOfBirth    `json:"dob" bson:"dob"`
	Location   Location       `json:"location" bson:"location"`
	Skills     []Technologies `json:"skills" bson:"skills"`
	Projects   []Application  `json:"project" bson:"project"`
	Contact    ContactInfo    `json:"contact" bson:"contact"`
	Experience []Job          `json:"experience" bson:"experience"`
}

type Job struct {
	CompanyName     string `json:"company_name" bson:"company_name"`
	WorkDescription string `json:"work_description" bson:"work_description"`
	YearsWorked     int    `json:"years_worked" bson:"years_worked"`
}

type Technologies struct {
	TechName          string `json:"tech_name" bson:"tech_name"`
	TechLink          string `json:"tech_link" bson:"tech_link"`
	YearsOfExperience int    `json:"year_of_experience" bson:"year_of_experience"`
	TechDescription   string `json:"tech_description" bson:"tech_description"`
}

type Location struct {
	State   string `json:"state" bson:"state"`
	City    string `json:"city" bson:"state"`
	ZipCode int    `json:"zip_code" bson:"zip_code"`
}

type Email struct {
	Name         string `json:"email_name" bson:"email_name"`
	EmailAddress string `json:"email_address" bson:"email_address"`
	EmailSubject string `json:"email_subject" bson:"email_subject"`
	EmailMessage string `json:"email_message" bson:"email_message"`
}
