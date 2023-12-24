package create

import (
	"drywave/controllers/create/files"
	"drywave/controllers/create/users"
	"github.com/gofiber/fiber/v2"
)

func Create(context *fiber.Ctx) (error error, message string, data any) {
	var model = context.Params("model")

	switch model {

	case "users":
		err, data, message := users.CreateUser(context)
		if err != nil {
			return err, message, data
		}
		return nil, message, data

	case "files":
		err, data, message := files.UploadFile(context)
		if err != nil {
			return err, message, data
		}
		return nil, message, data

	default:
		return nil, "No matching controller found", nil

	}
}
