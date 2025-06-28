package routes

import (
	"fmt"
	"neko-love/handlers"
	"neko-love/services"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures the HTTP routes for the Fiber application.
// It registers the following endpoints:
//   - GET /api/v4/:category: Returns a random image URL from the specified category as JSON.
//     Responds with 404 if the category does not exist or no image is found.
//   - GET /api/v4/images/:category/:name: Serves the image file with the given name from the specified category.
//     Responds with 404 if the category does not exist.
//
// The function also sets appropriate cache control headers for the random image endpoint.
func SetupRoutes(app *fiber.App) {
	app.Get("/api/v4/:category", func(c *fiber.Ctx) error {
		category := c.Params("category")
		path, exists := services.GetCategoryPath(category)
		if !exists {
			return fiber.ErrNotFound
		}

		imageName, err := handlers.PickRandomImageName(path)
		if err != nil {
			return fiber.ErrNotFound
		}

		imageURL := fmt.Sprintf("%s/api/v4/images/%s/%s", c.BaseURL(), category, imageName)

		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")

		return c.JSON(fiber.Map{
			"url": imageURL,
		})
	})

	app.Get("/api/v4/images/:category/:name", func(c *fiber.Ctx) error {
		category := c.Params("category")
		name := c.Params("name")
		path, exists := services.GetCategoryPath(category)
		if !exists {
			return fiber.ErrNotFound
		}

		imagePath := fmt.Sprintf("%s/%s", path, name)
		return c.SendFile(imagePath)
	})
}
