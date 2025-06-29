package filters

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "github.com/chai2010/webp"
)

// Greyscale converts the given image to greyscale using standard luminance calculation.
// It iterates over each pixel, computes the luminance based on the RGB values, and sets
// the resulting pixel to a shade of grey with the original alpha value preserved.
//
// Parameters:
//   img image.Image - The source image to be converted to greyscale.
//
// Returns:
//   image.Image - A new image in greyscale.
func Greyscale(img image.Image) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			a8 := uint8(a >> 8)

			lum := uint8((0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8)))

			dst.Set(x, y, color.NRGBA{R: lum, G: lum, B: lum, A: a8})
		}
	}

	return dst
}