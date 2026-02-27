package theme

import (
	"go-ds/color"
	"testing"
)

func TestThemeCSS(t *testing.T) {
	theme := NewTheme()

	ink := color.NewRamp()
	ink.AddKey("#FEFCFF", 0)
	ink.AddKey("#64617A", 0.562)
	ink.AddKey("#020103", 1)

	neutral := color.NewMode("neutral")
	neutral.AddRole(color.Role{
		Name:    "content",
		Default: color.UIColor{Light: ink.At(0.3), Dark: ink.At(0.8)},
		Hover:   color.UIColor{Light: ink.At(0.28), Dark: ink.At(0.78)},
	})
	theme.AddColorMode(neutral)

	css := theme.ToCSS()
	println(css)
}
