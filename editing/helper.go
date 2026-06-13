package editing

import (
	"image"
	"image/color"
	"slices"
)

func cloneImage(img *image.RGBA64) *image.RGBA64 {
	return &image.RGBA64{
		Pix:    slices.Clone(img.Pix),
		Stride: img.Stride,
		Rect:   img.Rect,
	}
}

func clamp(val, min, max float64) float64 {
	if val < min {
		return min
	}

	if val > max {
		return max
	}

	return val
}

func rawRGBA(col color.Color) (float64, float64, float64, float64) {
	r, g, b, a := col.RGBA()
	fr, fg, fb, fa := float64(r), float64(g), float64(b), float64(a)

	if fa > 0 {
		fr = (fr * 65535) / fa
		fg = (fg * 65535) / fa
		fb = (fb * 65535) / fa
	}

	return fr, fg, fb, fa
}
