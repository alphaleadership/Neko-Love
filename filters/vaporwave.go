package filters

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "github.com/chai2010/webp"
)

func Vaporwave(img image.Image) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			r8 := float64(r >> 8)
			g8 := float64(g >> 8)
			b8 := float64(b >> 8)

			// Boost rose et cyan, réduire vert
			newR := clamp8(int(r8*1.2 + 30))
			newG := clamp8(int(g8 * 0.9))
			newB := clamp8(int(b8*1.2 + 20))

			dst.Set(x, y, color.NRGBA{R: newR, G: newG, B: newB, A: uint8(a >> 8)})
		}
	}

	return dst
}
