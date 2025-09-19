package main

import (
	"neko-love/routes"
	"neko-love/services/cache"

	"github.com/gofiber/fiber/v2"
)

// main is the entry point of the application. It initializes a new Fiber web server,
// starts watching for asset changes, sets up the application routes, and begins
// listening for incoming HTTP requests on port 3030.
func main() {
	app := fiber.New()

	cacheAssets, err := cache.New("./assets")
	if err != nil {
		panic("Failed to initialize image cache: " + err.Error())
	}

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("cacheAssets", cacheAssets)
		return c.Next()
	})

	routes.SetupRoutes(app)
	app.Listen(":3030")
}
