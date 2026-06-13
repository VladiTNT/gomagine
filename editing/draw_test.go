package editing_test

import (
	"image"
	"testing"

	"github.com/VladiTNT/gomagine/editing"
)

// go test -run TestDraw ./editing -v
func TestDraw(t *testing.T) {
	baseImg, _, err := editing.LoadImage("../assets/PFP.png")
	if err != nil {
		t.Logf("Error loading image: %v\n", err)
		t.FailNow()
	}

	drawImg, _, err := editing.LoadImage("../assets/drawSprite.png")
	if err != nil {
		t.Logf("Error loading image: %v\n", err)
		t.FailNow()
	}

	editing.DrawWhere(baseImg, drawImg, image.Pt(108, 108))

	err = editing.WriteImage("../assets/tests/draw.png", "png", baseImg)
	if err != nil {
		t.Logf("Error writing image: %v\n", err)
		t.Fail()
	}
}
