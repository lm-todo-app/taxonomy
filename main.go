package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/lm-todo-app/taxonomy/db"
	"github.com/lm-todo-app/taxonomy/resources"
)

func main() {
	db.InitDatabase()
	appConfig := resources.InitConfig()
	app := fiber.New(appConfig)
	resources.SetupRoutes(app)
	app.Listen(":8090")
}
