package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/data", func(c *fiber.Ctx) error {
		return c.JSON(getData("data.csv"))
	})

	app.Get("/search", func(c *fiber.Ctx) error {
		regions := c.Query("regions")
		departments := c.Query("departments")
		municipalities := c.Query("municipalities")
		return c.JSON(search(regions, departments, municipalities))
	})

	app.Listen(":3000")
}
