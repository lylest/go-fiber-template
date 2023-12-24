package read_users

import (
	context2 "context"
	"drywave/connection"
	"drywave/models"
	"drywave/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func Login(context *fiber.Ctx) (error error, data any, message string) {
	var collection = connection.GetDBCollection("users")
	var user = new(models.UserLogin)

	isValid := context.BodyParser(user)
	if isValid != nil {
		return isValid, nil, "Invalid body values"
	}

	var result models.User
	mongoErr := collection.FindOne(context2.TODO(), bson.D{
		{"email", user.Email},
	}).Decode(&result)

	if mongoErr != nil {
		return mongoErr, nil, "User not found"
	}

	hashErr, passwordCheck := utils.CheckPasswordHash(user.Password, result.Password)

	if !passwordCheck {
		return hashErr, nil, "Incorrect Password"
	}

	var newUserToken = new(models.UserToken)
	newUserToken.ID = result.ID
	newUserToken.RoleID = result.RoleID
	newUserToken.Email = result.Email

	newToken, tokenErr := utils.CreateNewToken(*newUserToken)
	if tokenErr != nil {
		return tokenErr, nil, "Failed to authenticate user"
	}
	result.Token = newToken

	expirationTime := time.Now().Add(2 * 24 * time.Hour)
	maxAge := int(expirationTime.Sub(time.Now()).Seconds())

	context.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    newToken,
		HTTPOnly: true,
		SameSite: "Strict",
		MaxAge:   maxAge,
	})

	return nil, result, "Login success"

}
