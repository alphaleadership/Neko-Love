package routes

import (
	"fmt"
	"neko-love/services/cache"

	"github.com/gofiber/fiber/v2"
)

type ImageHandler struct {
	cache *cache.ImageCache
}

// NewImageHandler creates and returns a new ImageHandler instance with the provided ImageCache.
// The cache parameter is used to store and retrieve image data efficiently.
func NewImageHandler(c *cache.ImageCache) *ImageHandler {
	return &ImageHandler{cache: c}
}

// GetRandomImage handles HTTP requests to retrieve a random image from a specified category.
// It extracts the category from the route parameters, fetches a random image name from the cache,
// and retrieves the corresponding image path. If successful, it sets appropriate response headers
// and serves the image file. Returns a 404 error if the category or image is not found.
func (h *ImageHandler) GetRandomImage(c *fiber.Ctx) error {
	category := c.Params("category")

	name, err := h.cache.GetRandom(category)
	if err != nil {
		return fiber.ErrNotFound
	}

	image, ok := h.cache.GetImagePath(category, name)
	if !ok {
		return fiber.ErrNotFound
	}

	c.Locals("noCache", true)
	c.Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, name))
	return c.SendFile(image)
}

// GetRandomImageMeta handles the HTTP request to retrieve metadata for a random image
// within a specified category. It extracts the category from the route parameters,
// fetches a random image name from the cache, and retrieves its metadata. If successful,
// it returns a JSON response containing the image name, category, API path, human-readable
// size, size in bytes, last modified timestamp, and MIME type. If the category or image
// is not found, it responds with a 404 Not Found error.
//
// Route Params:
//   - category: string representing the image category.
//
// Response JSON:
//   - name:        string, image file name
//   - category:    string, image category
//   - path:        string, API path to the image
//   - size:        string, human-readable file size
//   - size_bytes:  int64, file size in bytes
//   - modified_at: time.Time, last modified timestamp
//   - mime_type:   string, MIME type of the image
//
// Returns 404 if the category or image is not found.
func (h *ImageHandler) GetRandomImageMeta(c *fiber.Ctx) error {
	category := c.Params("category")

	name, err := h.cache.GetRandom(category)
	if err != nil {
		return fiber.ErrNotFound
	}

	meta, ok := h.cache.GetImageMeta(category, name)
	if !ok {
		return fiber.ErrNotFound
	}

	return c.JSON(fiber.Map{
		"name":        name,
		"category":    category,
		"path":        fmt.Sprintf("/api/v4/images/%s/%s", category, name),
		"size":        meta.Readable,
		"size_bytes":  meta.Size,
		"modified_at": meta.ModifiedAt,
		"mime_type":   meta.MimeType,
	})
}

// ServeImage handles HTTP requests to serve an image file based on the provided
// category and name parameters in the URL. It retrieves the image path from the
// cache and sends the file as a response. If the image is not found in the cache,
// it returns a 404 Not Found error.
func (h *ImageHandler) ServeImage(c *fiber.Ctx) error {
	category := c.Params("category")
	name := c.Params("name")

	path, ok := h.cache.GetImagePath(category, name)
	if !ok {
		return fiber.ErrNotFound
	}

	return c.SendFile(path)
}

// RegisterImageRoutes registers image-related API routes to the provided Fiber router.
// It sets up middleware to inject an ImageHandler into the request context using a cached image asset store.
// The following endpoints are registered:
//   - GET /:category: Returns random image metadata for the specified category.
//   - GET /images/:category/:name: Serves the image file for the given category and image name.
func RegisterImageRoutes(router fiber.Router) {
	router.Use(func(c *fiber.Ctx) error {
		c.Locals("handler", NewImageHandler(c.Locals("cacheAssets").(*cache.ImageCache)))
		return c.Next()
	})

	router.Get("/:category", func(c *fiber.Ctx) error {
		handler := c.Locals("handler").(*ImageHandler)
		return handler.GetRandomImageMeta(c)
	})

	router.Get("/images/:category/:name", func(c *fiber.Ctx) error {
		handler := c.Locals("handler").(*ImageHandler)
		return handler.ServeImage(c)
	})
}
