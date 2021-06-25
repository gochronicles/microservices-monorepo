package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go-fiber-template/internal/patient/queries"
	"go-fiber-template/pkg/utils"
	"go-fiber-template/platform/database"
)

var patient *queries.Patient

func Create(c *fiber.Ctx) error {
	patient := &queries.Patient{}
	if err := c.BodyParser(patient); err != nil {
		panic(err)
	}
	result, err := database.Write(patient.Create)
	if err != nil {
		panic(err)
	}
	return utils.SuccessResponse(c, result.(*interface{}))
}

func GetByMRN(c *fiber.Ctx) error {
	mrn := c.Params("patientMRN")

	fmt.Println(mrn)
	result, err := patient.GetByMRN(mrn)
	if err != nil {
		panic(err)
	}
	return utils.SuccessResponse(c, result.(*interface{}))
}

func GetAll(c *fiber.Ctx) error {
	result, err := database.Read(patient.GetAll)
	if err != nil {
		panic(err)
	}
	return utils.SuccessResponse(c, result.(*interface{}))
}
