package routes

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"io"
	"net/http"
	"strings"

	"neko-love/services"

	"github.com/gofiber/fiber/v2"
)

// RegisterFilterRoutes registers the filter-related API routes to the provided Fiber router.
// It defines a GET endpoint "/filters/:filter" that applies the specified image filter to an image
// provided via the "image" query parameter. The endpoint supports GIF images with special handling
// and applies the requested filter to other image formats. The processed image is returned in the
// original format. Returns appropriate HTTP errors for missing parameters or processing failures.
//
// Routes:
//   GET /filters/:filter?image=<image_url>
//
// Parameters:
//   - filter: The name of the filter to apply (path parameter).
//   - image:  The URL of the image to process (query parameter).
func RegisterFilterRoutes(router fiber.Router) {
	router.Get("/filters/:filter", func(c *fiber.Ctx) error {
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
		c.Locals("noCache", true)

		if strings.HasPrefix(format, "image/gif") {
			return handleGIF(c, filter, data)
		}

		srcImg, formatStr, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to decode image")
		}

		result := services.ApplyFilter(filter, srcImg)
		return services.EncodeAndSetContentType(c, result, formatStr)
	})
}

// fetchImage retrieves the content of the image from the specified URL.
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

// handleGIF processes a GIF image using the specified filter and writes the filtered GIF to the response.
// It decodes the input GIF data, applies the filter via services.ProcessGIF, and encodes the result back to the client.
// Returns a Fiber error if decoding or processing fails.
//
// Parameters:
//   - c: Fiber context for the HTTP request and response.
//   - filter: The name of the filter to apply to the GIF.
//   - data: The raw GIF image data as a byte slice.
//
// Returns:
//   - error: An error if the GIF cannot be decoded, processed, or encoded; otherwise, nil.
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
