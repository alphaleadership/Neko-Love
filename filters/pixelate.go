package filters

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "github.com/chai2010/webp"
)

// Pixelate applies a pixelation effect to the given image by dividing it into blocks of fixed size
// and replacing each block with its average color. The function returns a new image with the effect applied.
//
// Parameters:
//   - img: The source image to be pixelated.
//
// Returns:
//   - image.Image: A new image with the pixelation effect applied.
func Pixelate(img image.Image) image.Image {
	bounds := img.Bounds()
	dst := image.NewRGBA(bounds)
	blockSize := 6

	for y := bounds.Min.Y; y < bounds.Max.Y; y += blockSize {
		for x := bounds.Min.X; x < bounds.Max.X; x += blockSize {
			var rTotal, gTotal, bTotal, aTotal uint32
			var count uint32

			for yy := y; yy < y+blockSize && yy < bounds.Max.Y; yy++ {
				for xx := x; xx < x+blockSize && xx < bounds.Max.X; xx++ {
					r, g, b, a := img.At(xx, yy).RGBA()
					rTotal += r
					gTotal += g
					bTotal += b
					aTotal += a
					count++
				}
			}

			avgColor := color.NRGBA{
				R: uint8((rTotal / count) >> 8),
				G: uint8((gTotal / count) >> 8),
				B: uint8((bTotal / count) >> 8),
				A: uint8((aTotal / count) >> 8),
			}

			for yy := y; yy < y+blockSize && yy < bounds.Max.Y; yy++ {
				for xx := x; xx < x+blockSize && xx < bounds.Max.X; xx++ {
					dst.Set(xx, yy, avgColor)
				}
			}
		}
	}

	return dst
}
