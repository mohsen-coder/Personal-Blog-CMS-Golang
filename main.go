package main

import (
	"log"
	"mohsen-coder/PersonalBlog/config"

	"github.com/gofiber/fiber/v2"
)

func homeRoute(c *fiber.Ctx) error {
	return c.SendString("hello world")
}

func main() {
	config.ConnectDb()
	app := fiber.New()

	app.Get("/", homeRoute)

	log.Fatal(app.Listen(":3000"))
}
