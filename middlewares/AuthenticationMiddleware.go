package middlewares

import (
	"drywave/models"
	"drywave/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func AuthenticationMiddleware(context *fiber.Ctx) error {
	var action = context.Params("action")
	var model = context.Params("model")

	var isFreeRoute = false

	for _, routeItem := range utils.FreeRoutes {
		if routeItem.Model == model && contains(routeItem.Action, action) {
			isFreeRoute = true
			break
		}
		isFreeRoute = false
	}

	if isFreeRoute {
		return context.Next()
	}

	TokenBody, err := verifyToken(context)
	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"messages": "Unauthorized try to login",
		})
	}
	context.Locals("userDetails", TokenBody)
	return context.Next()

}

func contains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func verifyToken(context *fiber.Ctx) (*models.UserToken, error) {
	var token = context.Cookies("token")

	if len(token) == 0 {
		bearerToken, err := validateUsingBearerToken(context)
		if err != nil {
			return nil, err
		}
		return bearerToken, nil
	}

	TokenBody, err := utils.CheckToken(token)
	if err != nil {
		return nil, err
	}
	return TokenBody, nil
}

func validateUsingBearerToken(context *fiber.Ctx) (*models.UserToken, error) {
	var bearerToken = context.Get("Authorization")
	var token string
	if len(bearerToken) == 0 {
		bearerToken = "no token"
	}

	token = strings.Split(bearerToken, " ")[1]

	TokenBody, err := utils.CheckToken(token)
	if err != nil {
		return nil, err
	}
	return TokenBody, nil
}
