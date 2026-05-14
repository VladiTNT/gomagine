package editing_test

import (
	"testing"

	"github.com/VladiTNT/gomagine/editing"
)

// go test -run TestGrayScale ./editing -v
func TestGrayScale(t *testing.T) {
	img, _, err := editing.LoadImage("./PFP.png")
	if err != nil {
		t.Logf("Error loading image: %v\n", err)
		t.FailNow()
	}

	editing.Transform(img, editing.GrayScale)

	err = editing.WriteImage("./grayscale.png", "png", img)
	if err != nil {
		t.Logf("Error writing image: %v\n", err)
		t.Fail()
	}
}
