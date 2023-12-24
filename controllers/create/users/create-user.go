package users

import (
	context2 "context"
	"drywave/connection"
	"drywave/models"
	"drywave/utils"
	"drywave/validation"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func CreateUser(context *fiber.Ctx) (error error, data any, message string) {
	var newUser = new(models.User)
	var collection = connection.GetDBCollection("users")

	decoderErr := utils.DecodeJSON(context, newUser)
	if decoderErr != nil {
		return decoderErr, nil, "Failed to decode json data"
	}

	validationMessage, validationErr := validation.Validate(newUser)
	if validationErr != nil {
		return validationErr, nil, validationMessage
	}

	hashedPassword, _ := utils.HashPassword(newUser.Password)

	newUser.ID = primitive.NewObjectID()
	newUser.CreatedAt = time.Now().UTC()
	newUser.UpdatedAt = time.Now().UTC()
	newUser.Shops = make([]primitive.ObjectID, 0)
	newUser.Permissions = utils.GeneratePermissions(newUser.RoleID)
	newUser.Password = hashedPassword

	var userExist models.User
	mongoErr := collection.FindOne(context2.TODO(), bson.D{
		{"email", newUser.Email},
	}).Decode(&userExist)

	if mongoErr == nil {
		return mongoErr, nil, "Email already used"
	}

	var result, err = collection.InsertOne(context.Context(), newUser)

	if err != nil {
		return err, nil, "Failed to register user"
	}
	return nil, result, "User registered successful"

}
