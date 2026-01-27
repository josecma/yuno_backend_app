package main

import (
	"yuno/src/api/routes"
	"yuno/src/common/config"
	"yuno/src/common/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	/* config */
	cfg := config.Load()

	/* database */
	db := database.Connect(&config.Load().Database)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	/* setup routes */
	routes.SetupUserRoutes(v1)

	app.Listen(":3001")

}
