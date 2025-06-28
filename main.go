package main

import (
	"neko-love/routes"
	"neko-love/services"

	"github.com/gofiber/fiber/v2"
)

// main is the entry point of the application. It initializes a new Fiber web server,
// starts watching for asset changes, sets up the application routes, and begins
// listening for incoming HTTP requests on port 3030.
func main() {
	app := fiber.New()

	services.WatchAssets()

	routes.SetupRoutes(app)

	app.Listen(":3030")
}
