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

func GrayScale(c color.Color) color.Color {
	r, g, b, a := c.RGBA()
	y := (r + g + b + a) / 4
	return color.Gray16{Y: uint16(y)}
}

func Brightness(factor float64) PixelTransformFunc {
	return func(c color.Color) color.Color {
		r, g, b, a := c.RGBA()

		newR := float64(r) * factor
		newG := float64(g) * factor
		newB := float64(b) * factor

		if newR > float64(a) {
			newR = float64(a)
		}

		if newG > float64(a) {
			newG = float64(a)
		}

		if newB > float64(a) {
			newB = float64(a)
		}

		return color.RGBA64{
			R: uint16(newR),
			G: uint16(newG),
			B: uint16(newB),
			A: uint16(a),
		}
	}
}

func Contrast(factor float64) PixelTransformFunc {
	return func(c color.Color) color.Color {
		r, g, b, a := c.RGBA()

		fa := float64(a)

		// Divide by alpha to get the raw channels
		newR := float64(r) / fa
		newG := float64(g) / fa
		newB := float64(b) / fa

		// Apply standard contrast formula
		newR = (newR-0.5)*factor + 0.5
		newG = (newG-0.5)*factor + 0.5
		newB = (newB-0.5)*factor + 0.5

		// Clamp invalid values
		newR = clamp(newR, 0.0, 1.0)
		newG = clamp(newG, 0.0, 1.0)
		newB = clamp(newB, 0.0, 1.0)

		return color.RGBA64{
			R: uint16(newR * fa),
			G: uint16(newG * fa),
			B: uint16(newB * fa),
			A: uint16(a),
		}
	}
}
