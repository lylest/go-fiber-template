package utils

import "github.com/gofiber/fiber/v2"

func DecodeJSON(context *fiber.Ctx, modelInterface any) error {
	err := context.BodyParser(modelInterface)
	if err != nil {
		return err
	}
	return nil
}
