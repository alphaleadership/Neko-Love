package filters

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math/rand"

	_ "github.com/chai2010/webp"
)

// Glitch applies a "glitch" visual effect to the given image by randomly shifting the red, green, and blue channels
// horizontally on each scanline, and by adding several random horizontal color bands. This creates a distorted,
// glitch-art appearance. The function returns a new image with the effect applied.
//
// Parameters:
//   img image.Image - The source image to which the glitch effect will be applied.
//
// Returns:
//   image.Image - A new image with the glitch effect applied.
func Glitch(img image.Image) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	height := bounds.Dy()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		offsetR := rand.Intn(6) - 3
		offsetG := rand.Intn(6) - 3
		offsetB := rand.Intn(6) - 3

		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rx := clamp(x+offsetR, bounds.Min.X, bounds.Max.X-1)
			gx := clamp(x+offsetG, bounds.Min.X, bounds.Max.X-1)
			bx := clamp(x+offsetB, bounds.Min.X, bounds.Max.X-1)

			r, _, _, _ := img.At(rx, y).RGBA()
			_, g, _, _ := img.At(gx, y).RGBA()
			_, _, b, a := img.At(bx, y).RGBA()

			dst.Set(x, y, color.NRGBA{
				R: uint8(r >> 8),
				G: uint8(g >> 8),
				B: uint8(b >> 8),
				A: uint8(a >> 8),
			})
		}
	}

	for i := 0; i < 5; i++ {
		yStart := rand.Intn(height)
		bandHeight := rand.Intn(10) + 5
		colorShift := uint8(rand.Intn(100))

		for y := yStart; y < yStart+bandHeight && y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, a := dst.At(x, y).RGBA()
				dst.Set(x, y, color.NRGBA{
					R: uint8(r>>8) ^ colorShift,
					G: uint8(g>>8) ^ colorShift,
					B: uint8(b>>8) ^ colorShift,
					A: uint8(a >> 8),
				})
			}
		}
	}

	return dst
}