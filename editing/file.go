package editing

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func LoadImage(path string) (*image.RGBA64, string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	// Decode the image
	img, imgFmt, err := image.Decode(f)
	if err != nil {
		return nil, "", err
	}

	imgRGBA64 := image.NewRGBA64(img.Bounds())

	// Copy pixel by pixel because idk
	for x := range img.Bounds().Dx() {
		for y := range img.Bounds().Dy() {
			imgRGBA64.Set(x, y, img.At(x, y))
		}
	}

	return imgRGBA64, imgFmt, nil
}

func WriteImage(path, imgFmt string, img *image.RGBA64) error {
	var buf bytes.Buffer

	switch imgFmt {
	case "png":
		if err := png.Encode(&buf, img); err != nil {
			return err
		}

		return flushBuffer(&buf, path)
	case "jpg", "jpeg":
		if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100}); err != nil {
			return err
		}

		return flushBuffer(&buf, path)
	default:
		return fmt.Errorf("gomagine: image format '%s' not supported", imgFmt)
	}
}

// flushBuffer makes a file at 'path' and flushes buf to it.
func flushBuffer(buf *bytes.Buffer, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := buf.WriteTo(f); err != nil {
		return err
	}

	return nil
}
