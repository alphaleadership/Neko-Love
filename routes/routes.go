package routes

import (
	"neko-love/middlewares"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures the main application routes for the Fiber app.
// It applies middleware to disable caching, sets up API version 4 routes for images and filters,
// and registers debug routes under the "/debug" path.
//
// Parameters:
//   - app: The Fiber application instance to which the routes will be attached.
func SetupRoutes(app *fiber.App) {
	app.Use(middlewares.NoCache())

	api := app.Group("/api/v4")
	RegisterImageRoutes(api)
	RegisterFilterRoutes(api)

	debug := app.Group("/debug")
	RegisterDebugRoutes(debug)
}
