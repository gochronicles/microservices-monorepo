package controllers

import (
	"fmt"
	"go-fiber-template/internal/domain/queries"

	"github.com/gofiber/fiber/v2"
	"go-fiber-template/pkg/utils"
	"go-fiber-template/platform/database"
)

var domain *queries.Domain

func Create(c *fiber.Ctx) error {
	domain := &queries.Domain{}
	if err := c.BodyParser(domain); err != nil {
		panic(err)
	}
	result, err := database.Write(domain.Create)
	if err != nil {
		panic(err)
	}
	return utils.SuccessResponse(c, result.(*interface{}))
}

func GetById(c *fiber.Ctx) error {
	id := c.Params("id")

	fmt.Println(id)
	result, err := domain.GetById(id)
	if err != nil {
		panic(err)
	}
	return utils.SuccessResponse(c, result.(*interface{}))
}

func GetAll(c *fiber.Ctx) error {
	result, err := database.Read(domain.GetAll)
	if err != nil {
		panic(err)
	}
	return utils.SuccessResponse(c, result.(*interface{}))
}
