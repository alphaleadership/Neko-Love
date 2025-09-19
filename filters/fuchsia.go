package filters

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "github.com/chai2010/webp"
)

// Fuchsia applies a custom "fuchsia" color filter to the given image.
// The filter maps the luminance of each pixel to a specific color palette
// with fuchsia and dark tones, creating a stylized effect. The alpha channel
// is preserved from the original image.
//
// Parameters:
//   img image.Image - The source image to be filtered.
//
// Returns:
//   image.Image - A new image with the fuchsia filter applied.
func Fuchsia(img image.Image) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			r8 := float64(r >> 8)
			g8 := float64(g >> 8)
			b8 := float64(b >> 8)
			a8 := uint8(a >> 8)

			lum := (0.299*r8 + 0.587*g8 + 0.114*b8) / 255

			var nr, ng, nb uint8
			switch {
			case lum > 0.92:
				nr, ng, nb = 255, 255, 255
			case lum > 0.7:
				nr, ng, nb = 192, 88, 168
			case lum > 0.45:
				nr, ng, nb = 152, 40, 128
			case lum >= 0.15:
				nr, ng, nb = 35, 39, 42
			default:
				nr, ng, nb = 35, 39, 42
			}

			dst.Set(x, y, color.NRGBA{R: nr, G: ng, B: nb, A: a8})
		}
	}

	return dst
}