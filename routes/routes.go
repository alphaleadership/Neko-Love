package routes

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"io"
	"neko-love/services/cache"
	"net/http"
	"strings"

	"neko-love/services"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures the HTTP routes for the Fiber application.
// It registers the following endpoints:
//   - GET /api/v4/:category: Returns a random image URL from the specified category.
//   - GET /api/v4/images/:category/:name: Serves the image file for the given category and name.
//   - GET /api/v4/filters/:filter: Applies the specified filter to an image provided via query parameter and returns the result.
//   - GET /debug/cache/:category: Returns debug information about cached files in the specified category.
//
// The function expects a "cacheAssets" local variable of type *cache.ImageCache to be set in the Fiber context.
func SetupRoutes(app *fiber.App) {
	app.Get("/api/v4/:category", func(c *fiber.Ctx) error {
		category := c.Params("category")
		path, _ := c.Locals("cacheAssets").(*cache.ImageCache).GetRandom(category)

		imageURL := fmt.Sprintf("%s/api/v4/images/%s/%s", c.BaseURL(), category, path)

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

		cache := c.Locals("cacheAssets").(*cache.ImageCache)

		path, ok := cache.GetImagePath(category, name)
		if !ok {
			return fiber.ErrNotFound
		}

		return c.SendFile(path)
	})

	app.Get("/api/v4/filters/:filter", func(c *fiber.Ctx) error {
		filter := c.Params("filter")
		if filter == "" {
			return fiber.ErrNotFound
		}

		imageURL := c.Query("image")
		if imageURL == "" {
			return fiber.NewError(fiber.StatusBadRequest, "Image URL is required")
		}

		data, err := fetchImage(imageURL)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch image")
		}

		format := http.DetectContentType(data)
		if strings.HasPrefix(format, "image/gif") {
			return handleGIF(c, filter, data)
		}

		srcImg, formatStr, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to decode image")
		}

		result := services.ApplyFilter(filter, srcImg)

		c.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")

		return services.EncodeAndSetContentType(c, result, formatStr)
	})

	app.Get("/debug/cache/:category", func(c *fiber.Ctx) error {
		category := c.Params("category")
		files := c.Locals("cacheAssets").(*cache.ImageCache).GetFiles(category)
		return c.JSON(fiber.Map{
			"category": category,
			"count":    len(files),
			"files":    files,
		})
	})
}

// fetchImage retrieves the content of an image from the specified URL.
// It performs an HTTP GET request and returns the image data as a byte slice.
// If the request fails or the response status is not 200 OK, it returns an error.
func fetchImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch image, status %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

// handleGIF processes a GIF image with the specified filter and writes the result to the HTTP response.
// It decodes the input GIF data, applies the given filter using the services.ProcessGIF function,
// and encodes the filtered GIF back to the response. If any error occurs during decoding or processing,
// it returns a 500 Internal Server Error with an appropriate message.
//
// Parameters:
//   - c: Fiber context used to manage the HTTP request and response.
//   - filter: The name of the filter to apply to the GIF.
//   - data: The raw GIF image data as a byte slice.
//
// Returns:
//   - error: An error if decoding, processing, or encoding fails; otherwise, nil.
func handleGIF(c *fiber.Ctx, filter string, data []byte) error {
	gifReader := bytes.NewReader(data)
	gifData, err := gif.DecodeAll(gifReader)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to decode GIF fully")
	}
	filteredGIF, err := services.ProcessGIF(filter, gifData)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to process GIF")
	}
	c.Set("Content-Type", "image/gif")
	return gif.EncodeAll(c.Context().Response.BodyWriter(), filteredGIF)
}