package main

import (
	"conferencia/goClientHttp/db"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// get method with the id parameter
	app.Post("/", func(c *fiber.Ctx) error {

		// get the body
		type bodyId struct {
			Id string `json:"id"`
		}

		var body bodyId

		if err := c.BodyParser(&body); err != nil {
			log.Println(err)
			return c.SendStatus(500)
		}

		// get data from db
		db := db.AppointmentCollection{}

		result, err := db.GetAppointments(body.Id)

		if err != nil {
			log.Println(err)
			return c.SendStatus(500)
		}

		return c.JSON(result)

	})

	app.Listen(":3002")
}
