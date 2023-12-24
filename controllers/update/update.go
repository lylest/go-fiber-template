package update

import "github.com/gofiber/fiber/v2"

func update(context *fiber.Ctx) error {
	return context.Status(200).JSON(fiber.Map{"message": "Yes we can update"})
}
