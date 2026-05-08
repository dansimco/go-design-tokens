package color

import (
	"fmt"
	"testing"
)

func TestGeneratingColor(t *testing.T) {
	ramp := NewRamp()
	ramp.AddKey("#000000", 0)
	ramp.AddKey("#ffffff", 1)
	colorStart := ramp.At(0)
	color := ramp.At(0.5)
	colorEnd := ramp.At(1)

	if color.ToHex() != "#777777" {
		t.Errorf("Expected color.ToHex() to equal #777777, got %s", color.ToHex())
	}

	if colorStart.ToHex() != "#000000" {
		t.Errorf("Expected colorStart.ToHex() to equal #000000, got %s", color.ToHex())
	}

	if colorEnd.ToHex() != "#FFFFFF" {
		t.Errorf("Expected colorEnd.ToHex() to equal #FFFFFF, got %s", color.ToHex())
	}

	swatches := ramp.CreateSwatches(16)

	for _, swatch := range swatches {
		fmt.Println("<div style=\"background-color:" + swatch.Color.ToHex() + ";\" title=\"" + fmt.Sprintf("%f", swatch.Position) + "\">&nbsp;</div>")
	}

}

func TestOrangeRamp(t *testing.T) {
	ramp := NewRamp()
	ramp.AddKey("#FEFCFF", 0)
	ramp.AddKey("#FF4F01", 0.4)
	ramp.AddKey("#010000", 1)

	stepCount := 32

	for i := 0; i <= stepCount; i++ {
		position := float64(i) / float64(stepCount)
		color := ramp.At(position)
		fmt.Println("<div style=\"background-color:" + color.ToHex() + ";\" title=\"" + fmt.Sprintf("%f", position) + "\">&nbsp;</div>")
	}
}
