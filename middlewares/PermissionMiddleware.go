package middlewares

import (
	context2 "context"
	"drywave/connection"
	"drywave/models"
	"drywave/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func PermissionMiddleware(context *fiber.Ctx) error {
	var model = context.Params("model")
	var action = context.Params("action")
	var collection = connection.GetDBCollection("users")

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

	UserDetails := context.Locals("userDetails").(*models.UserToken)
	if UserDetails.Email == "" {
		return context.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"messages": "Failed to check permissions",
		})
	}
	var result models.User
	mongoErr := collection.FindOne(context2.TODO(), bson.D{
		{"_id", UserDetails.ID},
		{"permissions", bson.D{
			{"$elemMatch", bson.D{
				{"name", model},
				{"list", bson.D{
					{"$in", bson.A{action}},
				}},
			}},
		}},
	}).Decode(&result)

	if mongoErr != nil {
		return context.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"messages": "You don't have permission to access " + strings.ToUpper(model),
		})
	}

	return context.Next()
}
