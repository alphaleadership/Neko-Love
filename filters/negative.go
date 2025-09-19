package filters

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "github.com/chai2010/webp"
)

// Negative returns a new image that is the negative (color-inverted) version of the input image.
// Each pixel's red, green, and blue channels are inverted, while the alpha channel is preserved.
// The function supports any image.Image input and outputs an *image.RGBA.
func Negative(img image.Image) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			a8 := uint8(a >> 8)

			newR := uint8(255 - (r >> 8))
			newG := uint8(255 - (g >> 8))
			newB := uint8(255 - (b >> 8))

			dst.Set(x, y, color.NRGBA{R: newR, G: newG, B: newB, A: a8})
		}
	}

	return dst
}