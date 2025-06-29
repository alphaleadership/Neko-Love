package filters

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "github.com/chai2010/webp"
)

// Posterize applies a posterization effect to the given image by reducing the number of color levels.
// The function processes each pixel, mapping its red, green, and blue channels to one of a fixed number
// of discrete levels (in this case, 4). The result is an image with fewer distinct colors, creating a
// stylized, poster-like appearance.
//
// Parameters:
//   img image.Image - The source image to be posterized.
//
// Returns:
//   image.Image - A new image with the posterization effect applied.
func Posterize(img image.Image) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)
	levels := 4

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			r8 := uint8((r >> 8) / uint32(256/levels) * uint32(256/levels))
			g8 := uint8((g >> 8) / uint32(256/levels) * uint32(256/levels))
			b8 := uint8((b >> 8) / uint32(256/levels) * uint32(256/levels))

			dst.Set(x, y, color.NRGBA{
				R: r8,
				G: g8,
				B: b8,
				A: uint8(a >> 8),
			})
		}
	}

	return dst
}
