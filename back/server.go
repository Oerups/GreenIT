package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/departements", func(c *fiber.Ctx) error {
		return c.JSON(getDepartements("data.csv"))
	})
	app.Get("/communes", func(c *fiber.Ctx) error {
		return c.JSON(getCommune("data.csv"))
	})
	app.Get("/regions", func(c *fiber.Ctx) error {
		return c.JSON(getRegion("data.csv"))
	})
	app.Get("/:city", func(c *fiber.Ctx) error {
		return c.JSON(findDataLine("data.csv", c.Params("city")))
	})

	app.Listen(":3000")
}
