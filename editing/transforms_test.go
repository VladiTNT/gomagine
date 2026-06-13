package editing_test

import (
	"image/color"
	"testing"

	"github.com/VladiTNT/gomagine/editing"
)

// go test -run TestGrayScale ./editing -v
func TestGrayScale(t *testing.T) {
	img, _, err := editing.LoadImage("../assets/PFP.png")
	if err != nil {
		t.Logf("Error loading image: %v\n", err)
		t.FailNow()
	}

	editing.Transform(img, editing.GrayScale)

	err = editing.WriteImage("../assets/tests/grayscale.png", "png", img)
	if err != nil {
		t.Logf("Error writing image: %v\n", err)
		t.Fail()
	}
}

// go test -run TestBrightness ./editing -v
func TestBrightness(t *testing.T) {
	img, _, err := editing.LoadImage("../assets/PFP.png")
	if err != nil {
		t.Logf("Error loading image: %v\n", err)
		t.FailNow()
	}

	editing.Transform(img, editing.Brightness(1.3))

	err = editing.WriteImage("../assets/tests/brightness.jpeg", "jpeg", img)
	if err != nil {
		t.Logf("Error writing image: %v\n", err)
		t.Fail()
	}
}

// go test -run TestContrast ./editing -v
func TestContrast(t *testing.T) {
	img, _, err := editing.LoadImage("../assets/PFP.png")
	if err != nil {
		t.Logf("Error loading image: %v\n", err)
		t.FailNow()
	}

	editing.Transform(img, editing.Contrast(1.3))

	err = editing.WriteImage("../assets/tests/contrast.png", "png", img)
	if err != nil {
		t.Logf("Error writing image: %v\n", err)
		t.Fail()
	}
}

// go test -run TestBlendColor ./editing -v
func TestBlendColor(t *testing.T) {
	img, _, err := editing.LoadImage("../assets/PFP.png")
	if err != nil {
		t.Logf("Error loading image: %v\n", err)
		t.FailNow()
	}

	editing.Transform(img, editing.BlendLinear(color.RGBA64{255, 0, 0, 255}, 0.5))

	err = editing.WriteImage("../assets/tests/blendColor.png", "png", img)
	if err != nil {
		t.Logf("Error writing image: %v\n", err)
		t.Fail()
	}
}
