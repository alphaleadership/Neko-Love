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

// Deepfry applies a "deep-fry" effect to the given image by increasing saturation,
// contrast, and shifting the color balance towards red-orange tones. This effect
// is achieved by manipulating the RGB channels of each pixel, resulting in a
// visually exaggerated and stylized image. The function returns a new image with
// the applied effect, preserving the original image's dimensions.
func Deepfry(img image.Image) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			r8 := float64(r >> 8)
			g8 := float64(g >> 8)
			b8 := float64(b >> 8)

			r8 = math.Min(255, r8*1.8+50)
			g8 = math.Min(255, g8*1.4)
			b8 = math.Min(255, b8*0.8)

			dst.Set(x, y, color.NRGBA{
				R: uint8(r8),
				G: uint8(g8),
				B: uint8(b8),
				A: uint8(a >> 8),
			})
		}
	}

	return dst
}
