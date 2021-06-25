package models

type Domain struct {
	id               		 string `json:"id" example:"12345"`
	FirstName                string `json:"firstName" example:"Tim" validate:"required"`
	LastName                 string `json:"lastName" example:"Anderson"`
	MiddleName               string `json:"middleName" example:"James"`
}

type PatientService interface {
	Create() (interface{}, error)
	GetById(id string) (interface{}, error)
	GetAll() (interface{}, error)
}
