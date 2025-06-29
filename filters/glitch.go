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

func Glitch(img image.Image) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	height := bounds.Dy()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		// Décalages aléatoires R, G, B pour chaque ligne
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

	// Ajout de bandes aléatoires
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

func clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func clamp8(v int) uint8 {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return uint8(v)
}
