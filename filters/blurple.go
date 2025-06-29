package filters

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"

	_ "github.com/chai2010/webp"
)

func Blurple(img image.Image) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	const (
		blurpleR = 88
		blurpleG = 101
		blurpleB = 242
	)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			r8 := float64(r >> 8)
			g8 := float64(g >> 8)
			b8 := float64(b >> 8)
			a8 := uint8(a >> 8)

			lum := (0.299*r8 + 0.587*g8 + 0.114*b8) / 255

			var newR, newG, newB uint8

			if lum > 0.90 {
				// Zone très claire → blanc pur
				newR, newG, newB = 255, 255, 255
			} else {
				// Teinte blurple dosée selon la luminosité
				newR = uint8(math.Min(lum*blurpleR, 255))
				newG = uint8(math.Min(lum*blurpleG, 255))
				newB = uint8(math.Min(lum*blurpleB, 255))
			}

			dst.Set(x, y, color.NRGBA{R: newR, G: newG, B: newB, A: a8})
		}
	}

	return dst
}