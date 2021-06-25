package models

type Patient struct {
	PatientMRN               string `json:"patientMRN" example:"12345"`
	FirstName                string `json:"firstName" example:"Tim" validate:"required"`
	LastName                 string `json:"lastName" example:"Anderson"`
	MiddleName               string `json:"middleName" example:"James"`
	DateOfBirth              string `json:"dateOfBirth" example:"23-02-1960"`
	Gender                   string `json:"gender" example:"male"`
	Prefix                   string `json:"prefix" example:"mr"`
	Mobile                   int    `json:"mobile" example:"1234567890" validate:"required,len=10"`
	Home                     int    `json:"home" example:"1234567890" validate:"required,len=10"`
	Work                     int    `json:"work" example:"1234567890" validate:"required,len=10"`
	PrimaryEmail             string `json:"primaryEmail" example:"abcd@email.com" validate:"email"`
	SecondaryEmail           string `json:"secondaryEmail" example:"1234567890" validate:"email"`
	SpokenLanguages          string `json:"spokenLanguages" example:"english"`
	WrittenLanguages         string `json:"writtenLanguages" example:"english"`
	PrimaryInsurance         string `json:"primaryInsurance" example:"aetna"`
	SecondaryInsurance       string `json:"secondaryInsurance" example:"aetna"`
	Occupation               string `json:"occcupation" example:""`
	MaritalStatus            string `json:"maritalStatus" example:""`
	Education                string `json:"education" example:""`
	HealthSummaryDescription string `json:"healthSummaryDescription" example:""`
}

type PatientService interface {
	Create() (interface{}, error)
	GetByMRN(mrn string) (interface{}, error)
	GetAll() (interface{}, error)
}
