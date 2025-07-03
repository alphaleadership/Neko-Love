package middlewares

import "github.com/gofiber/fiber/v2"

// NoCache is a Fiber middleware that sets HTTP headers to prevent client-side caching
// if the "noCache" local variable is set to true in the request context. It sets the
// "Cache-Control", "Pragma", and "Expires" headers to instruct browsers and proxies
// not to cache the response.
func NoCache() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		if noCache, ok := c.Locals("noCache").(bool); ok && noCache {
			c.Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
			c.Set("Pragma", "no-cache")
			c.Set("Expires", "0")
		}

		return err
	}
}
