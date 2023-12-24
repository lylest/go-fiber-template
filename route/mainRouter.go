package route

import (
	"drywave/controllers/create"
	"drywave/controllers/read"
	"github.com/gofiber/fiber/v2"
)

func MainRouter(context *fiber.Ctx) error {
	var action = context.Params("action")

	switch action {

	case "create":
		err, message, data := create.Create(context)
		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": message,
				"error":   err.Error(),
			})
		}
		return context.Status(fiber.StatusOK).JSON(fiber.Map{"message": message, "data": data})

	case "read":
		err, message, data := read.Read(context)
		if err != nil {
			return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": message,
				"error":   err.Error(),
			})
		}
		return context.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": message,
			"data":    data,
		})

	default:
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "No matching action route", "data": action})

	}
}
