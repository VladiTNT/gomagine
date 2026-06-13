package editing

import "image/color"

func LinearInterpolationBlending(col1, col2 color.Color, t float64) color.Color {
	r1, g1, b1, a1 := rawRGBA(col1)
	r2, g2, b2, _ := rawRGBA(col2)

	// Linear interpolation
	bR := (1-t)*r1 + t*r2
	bG := (1-t)*g1 + t*g2
	bB := (1-t)*b1 + t*b2
	bA := a1 // (1-t)*a1 + t*a2

	// Re-premultiply the colors back
	R := uint16(bR * bA / 65535)
	G := uint16(bG * bA / 65535)
	B := uint16(bB * bA / 65535)
	A := uint16(bA)

	return color.RGBA64{R, G, B, A}
}
