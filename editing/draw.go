package editing

import (
	"image"
	"image/draw"
)

func DrawWhere(dst, src *image.RGBA64, p image.Point) {
	draw.Draw(dst, dst.Rect.Add(p), src, image.Pt(0, 0), draw.Over)
}
