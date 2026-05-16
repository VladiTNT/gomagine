package editing_test

import (
	"testing"

	"github.com/VladiTNT/gomagine/editing"
)

// go test -run TestRotate90 ./editing -v
func TestRotate90(t *testing.T) {
	img, _, err := editing.LoadImage("../assets/PFP.png")
	if err != nil {
		t.Logf("Error loading image: %v\n", err)
		t.FailNow()
	}

	editing.Rotate90(img)

	err = editing.WriteImage("../assets/tests/rotate90.png", "png", img)
	if err != nil {
		t.Logf("Error writing image: %v\n", err)
		t.Fail()
	}
}
