package utils

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ConvertToMap(model interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	res, _ := json.Marshal(model)
	json.Unmarshal(res, &ret)
	fmt.Println(ret)
	return ret
}

func SuccessResponse(c *fiber.Ctx, response *interface{}) error {
	success := map[string]interface{}{"data": response, "error": nil}
	return c.JSON(success)
}
