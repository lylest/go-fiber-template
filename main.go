package main

import (
	"drywave/connection"
	"drywave/middlewares"
	"drywave/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	err := connection.ConnectToDb()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	//middlewares
	app.Use(func(context *fiber.Ctx) error {
		return context.Next()
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use("/:action/:model",

		func(context *fiber.Ctx) error {
			return middlewares.AuthenticationMiddleware(context)
		},
		func(context *fiber.Ctx) error { return middlewares.PermissionMiddleware(context) },
		func(context *fiber.Ctx) error {
			return route.MainRouter(context)
		},
	)

	err = app.Listen(":3000")
	if err != nil {
		return
	}
}
