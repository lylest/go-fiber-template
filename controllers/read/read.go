package read

import (
	"drywave/controllers/read/read_users"
	"github.com/gofiber/fiber/v2"
)

func Read(context *fiber.Ctx) (error error, message string, data any) {
	var model = context.Params("model")

	switch model {
	case "users":
		err, data, message := read_users.Login(context)
		if err != nil {
			return err, message, data
		}
		return nil, message, data

	default:
		return nil, "No matching controller found", nil
	}
}
