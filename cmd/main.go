package main

import (
	"fmt"
	"go-ds/color"
)

func main() {

	ink := color.NewRamp()
	ink.AddKey("#FEFCFF", 0)
	ink.AddKey("#64617A", 0.562)
	ink.AddKey("#020103", 1)

	intlOrange := color.NewRamp()
	intlOrange.AddKey("#FEFCFF", 0)
	intlOrange.AddKey("#FF4F01", 0.406)
	intlOrange.AddKey("#010000", 1)

	azimuth := color.NewRamp()
	azimuth.AddKey("#F9F9FB", 0)
	azimuth.AddKey("#6857DD", 0.4)
	azimuth.AddKey("#05000f", 1)

	patrick := color.NewRamp()
	patrick.AddKey("#FFFFFF", 0)
	patrick.AddKey("#DD79D8", 0.4)
	patrick.AddKey("#05000f", 1)

	kelp := color.NewRamp()
	kelp.AddKey("#EBFFE0", 0)
	kelp.AddKey("#6E9B55", 0.53)
	kelp.AddKey("#000000", 1)

	seaFoam := color.NewRamp()
	seaFoam.AddKey("#E2FFFF", 0)
	seaFoam.AddKey("#7AD0D3", 0.375)
	seaFoam.AddKey("#010511", 1)

	electric := color.NewRamp()
	electric.AddKey("#FFFFD8", 0)
	electric.AddKey("#F8FF6C", 0.375)
	electric.AddKey("#080a00", 1)

	// Semantic Colors

	// Neutral
	//

	neutral := color.Mode{}
	content := neutral.AddRole("content")
	content.AddState("default", color.UIColor{Light: ink.At(1), Dark: ink.At(1)})
	content.AddState("hover", color.UIColor{Light: ink.At(0.28), Dark: ink.At(0.78)})

}

func makeHTML(ramps []struct {
	name string
	ramp color.Ramp
}) {
	var html string
	stepCount := 32

	for _, r := range ramps {
		html += fmt.Sprintf("<div class=\"ramp\" data-name=\"%s\">\n", r.name)
		for i := 0; i <= stepCount; i++ {
			position := float64(i) / float64(stepCount)
			c := r.ramp.At(position)
			html += fmt.Sprintf("  <div style=\"background-color:%s;\" title=\"%s step %d (%.3f)\"><span style=\"display:none;\"%s</span></div>\n",
				c.ToHex(), r.name, i, position, c.ToHex())
		}
		html += "</div>\n\n"
	}

	fmt.Println(html)
}
