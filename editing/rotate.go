package editing

import "image"

func Rotate90(img *image.RGBA64) {
	temp := cloneImage(img)
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	for x := range w {
		for y := range h {
			img.SetRGBA64(x, y, temp.RGBA64At(y, w-x))
		}
	}
}
