package editing

import (
	"image/color"
)

func GrayScale(c color.Color) color.Color {
	r, g, b, a := c.RGBA()
	y := (r + g + b + a) / 4
	return color.Gray16{Y: uint16(y)}
}
