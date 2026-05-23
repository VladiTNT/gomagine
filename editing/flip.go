package editing

import "image"

func FlipHorizontal(img *image.RGBA64) {
	temp := cloneImage(img)
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	for x := range w {
		for y := range h {
			img.SetRGBA64(x, y, temp.RGBA64At(w-x, y))
		}
	}
}

func FlipVertical(img *image.RGBA64) {
	temp := cloneImage(img)
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	for x := range w {
		for y := range h {
			img.SetRGBA64(x, y, temp.RGBA64At(x, h-y))
		}
	}
}
