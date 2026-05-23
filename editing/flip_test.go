package editing_test

import (
	"testing"

	"github.com/VladiTNT/gomagine/editing"
)

// go test -run TestFlipHorizontal ./editing -v
func TestFlipHorizontal(t *testing.T) {
	img, _, err := editing.LoadImage("../assets/PFP.png")
	if err != nil {
		t.Logf("Error loading image: %v\n", err)
		t.FailNow()
	}

	editing.FlipHorizontal(img)

	err = editing.WriteImage("../assets/tests/flipHorizontal.png", "png", img)
	if err != nil {
		t.Logf("Error writing image: %v\n", err)
		t.Fail()
	}
}

// go test -run TestFlipVertical ./editing -v
func TestFlipVertical(t *testing.T) {
	img, _, err := editing.LoadImage("../assets/PFP.png")
	if err != nil {
		t.Logf("Error loading image: %v\n", err)
		t.FailNow()
	}

	editing.FlipVertical(img)

	err = editing.WriteImage("../assets/tests/flipVertical.png", "png", img)
	if err != nil {
		t.Logf("Error writing image: %v\n", err)
		t.Fail()
	}
}
