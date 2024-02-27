package main

import (
	"github.com/gofiber/fiber/v2"
)

type car struct {
	ID             string `json:"id"`
	Make           string `json:"make"`
	Model          string `json:"model"`
	ProductionDate int    `json:"productionDate"`
}

var cars = []car{
	{ID: "1", Make: "Mercedes", Model: "E-class", ProductionDate: 2023},
	{ID: "2", Make: "BMW", Model: "5-series", ProductionDate: 2022},
	{ID: "3", Make: "Audi", Model: "A7", ProductionDate: 2021},
	{ID: "4", Make: "Mercedes", Model: "S-class", ProductionDate: 2023},
}

func getCars(c *fiber.Ctx) error {
	return c.JSON(cars)
}

func createCar(c *fiber.Ctx) error {
	var newCar car
	if err := c.BodyParser(&newCar); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	cars = append(cars, newCar)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Car created successfully", "car": newCar})

}

func main() {
	app := fiber.New()
	app.Get("/", getCars)
	app.Post("/", createCar)
	app.Listen(":5050")

}
