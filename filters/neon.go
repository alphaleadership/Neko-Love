package filters

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "github.com/chai2010/webp"
)

func Neon(img image.Image) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			r8 := float64(r >> 8)
			g8 := float64(g >> 8)
			b8 := float64(b >> 8)
			a8 := uint8(a >> 8)

			// Boost contrast
			avg := (r8 + g8 + b8) / 3
			factor := 2.0
			if avg > 180 {
				factor = 2.5
			} else if avg < 50 {
				factor = 0.5
			}

			newR := clamp8(int(r8 * factor))
			newG := clamp8(int(g8 * factor * 0.8))
			newB := clamp8(int(b8 * factor * 1.2))

			dst.Set(x, y, color.NRGBA{R: newR, G: newG, B: newB, A: a8})
		}
	}

	return dst
}
