package editing

import (
	"image"
	"image/color"
)

type PixelTransformFunc func(c color.Color) color.Color

func Transform(img *image.RGBA64, f PixelTransformFunc) {
	for x := range img.Bounds().Dx() {
		for y := range img.Bounds().Dy() {
			img.Set(x, y, f(img.At(x, y)))
		}
	}
}
