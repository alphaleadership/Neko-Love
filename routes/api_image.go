package routes

import (
	"fmt"
	"neko-love/services/cache"

	"github.com/gofiber/fiber/v2"
)

// RegisterImageRoutes registers image-related API routes to the provided Fiber router.
// 
// Routes:
//   GET /:category
//     - Returns a random image URL from the specified category as JSON.
//     - Response: { "url": "<image_url>" }
//     - Middleware is expected to provide "cacheAssets" (of type *cache.ImageCache) in Locals.
//     - Sets "noCache" local for middleware to handle cache headers.
//
//   GET /images/:category/:name
//     - Serves the image file with the given name from the specified category.
//     - Returns 404 if the image is not found.
//     - Middleware is expected to provide "cacheAssets" (of type *cache.ImageCache) in Locals.
//
// Parameters:
//   router fiber.Router - The Fiber router to which the routes will be registered.
func RegisterImageRoutes(router fiber.Router) {
	router.Get("/:category", func(c *fiber.Ctx) error {
		category := c.Params("category")
		path, _ := c.Locals("cacheAssets").(*cache.ImageCache).GetRandom(category)

		imageURL := fmt.Sprintf("%s/api/v4/images/%s/%s", c.BaseURL(), category, path)
		c.Locals("noCache", true) // middleware will handle headers

		return c.JSON(fiber.Map{
			"url": imageURL,
		})
	})

	router.Get("/images/:category/:name", func(c *fiber.Ctx) error {
		category := c.Params("category")
		name := c.Params("name")

		cache := c.Locals("cacheAssets").(*cache.ImageCache)
		path, ok := cache.GetImagePath(category, name)
		if !ok {
			return fiber.ErrNotFound
		}
		return c.SendFile(path)
	})
}
