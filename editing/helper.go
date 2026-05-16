package editing

import (
	"image"
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
