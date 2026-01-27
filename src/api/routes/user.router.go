package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(router fiber.Router) {

	userRouteGroup := router.Group("/users")

	userRouteGroup.Get("/", func(c *fiber.Ctx) error { return c.SendString("") })

	userRouteGroup.Get("/:id", func(c *fiber.Ctx) error { return c.SendString("") })

	userRouteGroup.Post("/", func(c *fiber.Ctx) error { return c.SendString("") })

	userRouteGroup.Put("/:id", func(c *fiber.Ctx) error { return c.SendString("") })

	userRouteGroup.Delete("/:id", func(c *fiber.Ctx) error { return c.SendString("") })

	log.Println("user routes configured")

}
