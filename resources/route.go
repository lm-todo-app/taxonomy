package resources

import (
	"github.com/gofiber/fiber/v2"
)

func StartServer(port string) {
	app := fiber.New()
	SetupRoutes(app)
	app.Listen(port)
}

func SetupRoutes(app *fiber.App) {
	// add version: v1
	app.Get("/taxonomy", getTaxonomies)
	app.Post("/taxonomy", createTaxonomy)
	app.Get("/taxonomy/:id", getTaxonomy)
	app.Put("/taxonomy/:id", updateTaxonomy)
	app.Delete("/taxonomy/:id", deleteTaxonomy)
	// app.Post("/taxonomy/bulk", createTaxonomies)
}
