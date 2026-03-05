package theme

import (
	"go-ds/color"
	"os"
	"testing"
)

func TestThemeCSS(t *testing.T) {
	theme := NewTheme()

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

	//
	// NEUTRAL
	// ----------------

	// content
	neutral := theme.AddColorMode("neutral")
	neutral_content := neutral.AddRole("content")
	neutral_content.AddState("default", color.UIColor{Light: ink.At(0.9), Dark: ink.At(0.3124)})
	neutral_content.AddState("hover", color.UIColor{Light: ink.At(0.8), Dark: ink.At(0.219)})
	neutral_content.AddState("pressed", color.UIColor{Light: ink.At(0.7), Dark: ink.At(0.1)})
	neutral_content.AddState("focus", color.UIColor{Light: ink.At(0.6), Dark: ink.At(0.219)})
	neutral_content.AddState("disabled", color.UIColor{Light: ink.At(0.5), Dark: ink.At(0.406)})
	// secondary
	neutral_secondary := neutral.AddRole("content-secondary")
	neutral_secondary.AddState("default", color.UIColor{Light: ink.At(0.9), Dark: ink.At(0.5)})
	neutral_secondary.AddState("hover", color.UIColor{Light: ink.At(0.8), Dark: ink.At(0.156)})
	neutral_secondary.AddState("pressed", color.UIColor{Light: ink.At(0.7), Dark: ink.At(0.156)})
	neutral_secondary.AddState("focus", color.UIColor{Light: ink.At(0.6), Dark: ink.At(0.156)})
	neutral_secondary.AddState("disabled", color.UIColor{Light: ink.At(0.5), Dark: ink.At(0.8)})
	// highlight
	neutral_highlight := neutral.AddRole("content-highlight")
	neutral_highlight.AddState("default", color.UIColor{Light: ink.At(0.9), Dark: ink.At(0.3125)})
	neutral_highlight.AddState("hover", color.UIColor{Light: ink.At(0.8), Dark: ink.At(0.2)})
	neutral_highlight.AddState("pressed", color.UIColor{Light: ink.At(0.7), Dark: ink.At(0.2)})
	neutral_highlight.AddState("focus", color.UIColor{Light: ink.At(0.6), Dark: ink.At(0.2)})
	neutral_highlight.AddState("disabled", color.UIColor{Light: ink.At(0.5), Dark: ink.At(0.2)})
	// trim
	neutral_trim := neutral.AddRole("trim")
	neutral_trim.AddState("default", color.UIColor{Light: ink.At(0.9), Dark: ink.At(0.8125)})
	neutral_trim.AddState("hover", color.UIColor{Light: ink.At(0.8), Dark: ink.At(0.8125)})
	neutral_trim.AddState("pressed", color.UIColor{Light: ink.At(0.7), Dark: ink.At(0.8125)})
	neutral_trim.AddState("focus", color.UIColor{Light: ink.At(0.6), Dark: ink.At(0.8125)})
	neutral_trim.AddState("disabled", color.UIColor{Light: ink.At(0.5), Dark: ink.At(0.8125)})
	// surface
	neutral_surface := neutral.AddRole("surface")
	neutral_surface.AddState("default", color.UIColor{Light: ink.At(0.9), Dark: ink.At(0.875)})
	neutral_surface.AddState("hover", color.UIColor{Light: ink.At(0.8), Dark: ink.At(0.843)})
	neutral_surface.AddState("pressed", color.UIColor{Light: ink.At(0.7), Dark: ink.At(0.438)})
	neutral_surface.AddState("focus", color.UIColor{Light: ink.At(0.6), Dark: ink.At(0.438)})
	neutral_surface.AddState("disabled", color.UIColor{Light: ink.At(0.5), Dark: ink.At(0.875)})
	// surface_low
	neutral_surface_low := neutral.AddRole("surface-low")
	neutral_surface_low.AddState("default", color.UIColor{Light: ink.At(0.9), Dark: ink.At(0.906)})
	neutral_surface_low.AddState("hover", color.UIColor{Light: ink.At(0.8), Dark: ink.At(0.906)})
	neutral_surface_low.AddState("pressed", color.UIColor{Light: ink.At(0.7), Dark: ink.At(0.906)})
	neutral_surface_low.AddState("focus", color.UIColor{Light: ink.At(0.6), Dark: ink.At(0.906)})
	neutral_surface_low.AddState("disabled", color.UIColor{Light: ink.At(0.5), Dark: ink.At(0.906)})
	// surface_high
	neutral_surface_high := neutral.AddRole("surface-high")
	neutral_surface_high.AddState("default", color.UIColor{Light: ink.At(0.9), Dark: ink.At(0.94)})
	neutral_surface_high.AddState("hover", color.UIColor{Light: ink.At(0.8), Dark: ink.At(0.94)})
	neutral_surface_high.AddState("pressed", color.UIColor{Light: ink.At(0.7), Dark: ink.At(0.94)})
	neutral_surface_high.AddState("focus", color.UIColor{Light: ink.At(0.6), Dark: ink.At(0.94)})
	neutral_surface_high.AddState("disabled", color.UIColor{Light: ink.At(0.5), Dark: ink.At(0.94)})
	// background
	neutral_background := neutral.AddRole("surface-bg")
	neutral_background.AddState("default", color.UIColor{Light: ink.At(0.9), Dark: ink.At(0.94)})
	neutral_background.AddState("hover", color.UIColor{Light: ink.At(0.8), Dark: ink.At(0.94)})
	neutral_background.AddState("pressed", color.UIColor{Light: ink.At(0.7), Dark: ink.At(0.94)})
	neutral_background.AddState("focus", color.UIColor{Light: ink.At(0.6), Dark: ink.At(0.94)})
	neutral_background.AddState("disabled", color.UIColor{Light: ink.At(0.5), Dark: ink.At(0.94)})

	//
	// ACTION
	// ----------------

	// content
	action := theme.AddColorMode("action")
	action_content := action.AddRole("content")
	action_content.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.4375)})
	action_content.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.35)})
	action_content.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.3)})
	action_content.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.35)})
	action_content.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: ink.At(0.4375)})
	// secondary
	action_secondary := action.AddRole("content-secondary")
	action_secondary.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.6)})
	action_secondary.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.55)})
	action_secondary.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.5)})
	action_secondary.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.55)})
	action_secondary.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: ink.At(0.4375)})
	// highlight
	action_highlight := action.AddRole("content-highlight")
	action_highlight.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.3125)})
	action_highlight.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.2)})
	action_highlight.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.15)})
	action_highlight.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.2)})
	action_highlight.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: ink.At(0.6)})
	// trim
	action_trim := action.AddRole("trim")
	action_trim.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.8125)})
	action_trim.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.7)})
	action_trim.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.7)})
	action_trim.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.7)})
	action_trim.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: ink.At(0.8125)})
	// surface
	action_surface := action.AddRole("surface")
	action_surface.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.9)})
	action_surface.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.9)})
	action_surface.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.9)})
	action_surface.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.9)})
	action_surface.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: azimuth.At(0.9)})
	// surface_low
	action_surface_low := action.AddRole("surface-low")
	action_surface_low.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.95)})
	action_surface_low.AddState("hover", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.95)})
	action_surface_low.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.95)})
	action_surface_low.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.95)})
	action_surface_low.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: azimuth.At(0.95)})
	// surface_high
	action_surface_high := action.AddRole("surface-high")
	action_surface_high.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.88)})
	action_surface_high.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.88)})
	action_surface_high.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.88)})
	action_surface_high.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.88)})
	action_surface_high.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: azimuth.At(0.88)})
	// background
	action_background := action.AddRole("surface-bg")
	action_background.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.96)})
	action_background.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.96)})
	action_background.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.96)})
	action_background.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.96)})
	action_background.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: azimuth.At(0.96)})

	//
	// ACTION REVERSED
	// Light on dark colors for both light and dark mode
	// ----------------

	// content
	action_rev := theme.AddColorMode("action-rev")
	action_rev_content := action_rev.AddRole("content")
	action_rev_content.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.15)})
	action_rev_content.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.15)})
	action_rev_content.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.15)})
	action_rev_content.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.15)})
	action_rev_content.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: ink.At(0.4)})
	// secondary
	action_rev_secondary := action_rev.AddRole("content-secondary")
	action_rev_secondary.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.2)})
	action_rev_secondary.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.2)})
	action_rev_secondary.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.2)})
	action_rev_secondary.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.2)})
	action_rev_secondary.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: ink.At(0.2)})
	// highlight
	action_rev_highlight := action_rev.AddRole("content-highlight")
	action_rev_highlight.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.1)})
	action_rev_highlight.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0)})
	action_rev_highlight.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.05)})
	action_rev_highlight.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.05)})
	action_rev_highlight.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: ink.At(0.6)})
	// trim
	action_rev_trim := action_rev.AddRole("trim")
	action_rev_trim.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.46875)})
	action_rev_trim.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.46875)})
	action_rev_trim.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.46875)})
	action_rev_trim.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.46875)})
	action_rev_trim.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: ink.At(0.8125)})
	// surface
	action_rev_surface := action_rev.AddRole("surface")
	action_rev_surface.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.4375)})
	action_rev_surface.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.35)})
	action_rev_surface.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.35)})
	action_rev_surface.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.35)})
	action_rev_surface.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: ink.At(0.8)})
	// surface_low
	action_rev_surface_low := action_rev.AddRole("surface-low")
	action_rev_surface_low.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.95)})
	action_rev_surface_low.AddState("hover", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.95)})
	action_rev_surface_low.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.95)})
	action_rev_surface_low.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.95)})
	action_rev_surface_low.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: azimuth.At(0.95)})
	// surface_high
	action_rev_surface_high := action_rev.AddRole("surface-high")
	action_rev_surface_high.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.88)})
	action_rev_surface_high.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.88)})
	action_rev_surface_high.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.88)})
	action_rev_surface_high.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.88)})
	action_rev_surface_high.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: azimuth.At(0.88)})
	// background
	action_rev_background := action_rev.AddRole("surface-bg")
	action_rev_background.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.96)})
	action_rev_background.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.96)})
	action_rev_background.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.96)})
	action_rev_background.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.96)})
	action_rev_background.AddState("disabled", color.UIColor{Light: azimuth.At(0.5), Dark: azimuth.At(0.96)})

	css := theme.ToCSS()
	println(css)

	html := theme.GenerateHTMLPreview()

	// Write HTML to file
	err := os.WriteFile("theme_preview.html", []byte(html), 0644)
	if err != nil {
		t.Fatalf("Failed to write HTML file: %v", err)
	}
	t.Log("HTML preview written to theme_preview.html")
}
