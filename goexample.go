package main

import (
	"github.com/atolafsson/goexample/database"
	"github.com/atolafsson/goexample/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)
}

func main() {
	database.ConnectDb()
	//database.DB.Db.AutoMigrate(&models.Fact{})
	app := fiber.New()
	setupRoutes(app)

	app.Listen(":3000")
}
