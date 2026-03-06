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
	intlOrange.AddKey("#FAF9FF", 0)
	intlOrange.AddKey("#FF4F01", 0.406)
	intlOrange.AddKey("#010000", 1)

	azimuth := color.NewRamp()
	azimuth.AddKey("#F9F9FB", 0)
	azimuth.AddKey("#6857DD", 0.4)
	azimuth.AddKey("#05000f", 1)

	//
	// NEUTRAL
	// ----------------
	neutral := theme.AddColorMode("neutral")

	// highlight
	neutral_highlight := neutral.AddRole("content-highlight")
	neutral_highlight.Context = "foreground"
	neutral_highlight.AddState("default", color.UIColor{Light: ink.At(1), Dark: ink.At(0.1)})
	neutral_highlight.AddState("hover", color.UIColor{Light: ink.At(0.8), Dark: ink.At(0.2)})
	neutral_highlight.AddState("pressed", color.UIColor{Light: ink.At(0.7), Dark: ink.At(0.2)})
	neutral_highlight.AddState("focus", color.UIColor{Light: ink.At(0.6), Dark: ink.At(0.2)})
	// content
	neutral_content := neutral.AddRole("content")
	neutral_content.Context = "foreground"
	neutral_content.AddState("default", color.UIColor{Light: ink.At(0.75), Dark: ink.At(0.3124)})
	neutral_content.AddState("hover", color.UIColor{Light: ink.At(0.8), Dark: ink.At(0.219)})
	neutral_content.AddState("pressed", color.UIColor{Light: ink.At(0.7), Dark: ink.At(0.1)})
	neutral_content.AddState("focus", color.UIColor{Light: ink.At(0.6), Dark: ink.At(0.219)})
	// secondary
	neutral_secondary := neutral.AddRole("content-secondary")
	neutral_secondary.Context = "foreground"
	neutral_secondary.AddState("default", color.UIColor{Light: ink.At(0.45), Dark: ink.At(0.5)})
	neutral_secondary.AddState("hover", color.UIColor{Light: ink.At(0.5), Dark: ink.At(0.156)})
	neutral_secondary.AddState("pressed", color.UIColor{Light: ink.At(0.7), Dark: ink.At(0.156)})
	neutral_secondary.AddState("focus", color.UIColor{Light: ink.At(0.5), Dark: ink.At(0.156)})
	// trim
	neutral_trim := neutral.AddRole("trim")
	neutral_trim.AddState("default", color.UIColor{Light: ink.At(0.1), Dark: ink.At(0.8125)})
	neutral_trim.AddState("hover", color.UIColor{Light: ink.At(0.1), Dark: ink.At(0.8125)})
	neutral_trim.AddState("pressed", color.UIColor{Light: ink.At(0.1), Dark: ink.At(0.8125)})
	neutral_trim.AddState("focus", color.UIColor{Light: ink.At(0.1), Dark: ink.At(0.8125)})
	// surface
	neutral_surface := neutral.AddRole("surface")
	neutral_surface.Context = "background"
	neutral_surface.AddState("default", color.UIColor{Light: ink.At(0.01), Dark: ink.At(0.9)})
	neutral_surface.AddState("hover", color.UIColor{Light: ink.At(0.05), Dark: ink.At(0.875)})
	neutral_surface.AddState("pressed", color.UIColor{Light: ink.At(0.05), Dark: ink.At(0.875)})
	neutral_surface.AddState("focus", color.UIColor{Light: ink.At(0.05), Dark: ink.At(0.875)})
	// surface_low
	neutral_surface_low := neutral.AddRole("surface-low")
	neutral_surface_low.Context = "background"
	neutral_surface_low.AddState("default", color.UIColor{Light: ink.At(0.05), Dark: ink.At(0.906)})
	neutral_surface_low.AddState("hover", color.UIColor{Light: ink.At(0.8), Dark: ink.At(0.916)})
	neutral_surface_low.AddState("pressed", color.UIColor{Light: ink.At(0.7), Dark: ink.At(0.93)})
	neutral_surface_low.AddState("focus", color.UIColor{Light: ink.At(0.6), Dark: ink.At(0.906)})
	// surface_high
	neutral_surface_high := neutral.AddRole("surface-high")
	neutral_surface_high.Context = "background"
	neutral_surface_high.AddState("default", color.UIColor{Light: ink.At(0.08), Dark: ink.At(0.92)})
	neutral_surface_high.AddState("hover", color.UIColor{Light: ink.At(0.8), Dark: ink.At(0.90)})
	neutral_surface_high.AddState("pressed", color.UIColor{Light: ink.At(0.7), Dark: ink.At(0.91)})
	neutral_surface_high.AddState("focus", color.UIColor{Light: ink.At(0.6), Dark: ink.At(0.90)})
	// background
	neutral_background := neutral.AddRole("surface-bg")
	neutral_background.Context = "background"
	neutral_background.AddState("default", color.UIColor{Light: ink.At(0.03), Dark: ink.At(0.94)})
	neutral_background.AddState("hover", color.UIColor{Light: ink.At(0.8), Dark: ink.At(0.94)})
	neutral_background.AddState("pressed", color.UIColor{Light: ink.At(0.7), Dark: ink.At(0.94)})
	neutral_background.AddState("focus", color.UIColor{Light: ink.At(0.6), Dark: ink.At(0.94)})

	//
	// ACTION
	// ----------------
	action := theme.AddColorMode("action")

	// highlight
	action_highlight := action.AddRole("content-highlight")
	action_highlight.Context = "foreground"
	action_highlight.AddState("default", color.UIColor{Light: azimuth.At(0.5), Dark: azimuth.At(0.3125)})
	action_highlight.AddState("hover", color.UIColor{Light: azimuth.At(0.4), Dark: azimuth.At(0.2)})
	action_highlight.AddState("pressed", color.UIColor{Light: azimuth.At(0.4), Dark: azimuth.At(0.15)})
	action_highlight.AddState("focus", color.UIColor{Light: azimuth.At(0.4), Dark: azimuth.At(0.2)})
	// content
	action_content := action.AddRole("content")
	action_content.Context = "foreground"
	action_content.AddState("default", color.UIColor{Light: azimuth.At(0.4), Dark: azimuth.At(0.4375)})
	action_content.AddState("hover", color.UIColor{Light: azimuth.At(0.37), Dark: azimuth.At(0.35)})
	action_content.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.3)})
	action_content.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.35)})
	// secondary
	action_secondary := action.AddRole("content-secondary")
	action_secondary.Context = "foreground"
	action_secondary.AddState("default", color.UIColor{Light: azimuth.At(0.33), Dark: azimuth.At(0.58)})
	action_secondary.AddState("hover", color.UIColor{Light: azimuth.At(0.3), Dark: azimuth.At(0.55)})
	action_secondary.AddState("pressed", color.UIColor{Light: azimuth.At(0.3), Dark: azimuth.At(0.5)})
	action_secondary.AddState("focus", color.UIColor{Light: azimuth.At(0.3), Dark: azimuth.At(0.55)})
	// trim
	action_trim := action.AddRole("trim")
	action_trim.AddState("default", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.8125)})
	action_trim.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.7)})
	action_trim.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.7)})
	action_trim.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.7)})
	// surface
	action_surface := action.AddRole("surface")
	action_surface.Context = "background"
	action_surface.AddState("default", color.UIColor{Light: ink.At(0.01), Dark: ink.At(0.975)})
	action_surface.AddState("hover", color.UIColor{Light: azimuth.At(0.03), Dark: azimuth.At(0.9)})
	action_surface.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.9)})
	action_surface.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.9)})
	// surface_low
	action_surface_low := action.AddRole("surface-low")
	action_surface_low.Context = "background"
	action_surface_low.AddState("default", color.UIColor{Light: azimuth.At(0.025), Dark: azimuth.At(0.95)})
	action_surface_low.AddState("hover", color.UIColor{Light: azimuth.At(0.9), Dark: azimuth.At(0.95)})
	action_surface_low.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.95)})
	action_surface_low.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.95)})
	// surface_high
	action_surface_high := action.AddRole("surface-high")
	action_surface_high.Context = "background"
	action_surface_high.AddState("default", color.UIColor{Light: azimuth.At(0.05), Dark: azimuth.At(0.88)})
	action_surface_high.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.88)})
	action_surface_high.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.88)})
	action_surface_high.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.88)})
	// background
	action_background := action.AddRole("surface-bg")
	action_background.Context = "background"
	action_background.AddState("default", color.UIColor{Light: azimuth.At(0.02), Dark: azimuth.At(0.96)})
	action_background.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.96)})
	action_background.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.96)})
	action_background.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.96)})

	//
	// ACTION REVERSED
	// Light on dark colors for both light and dark mode
	// ----------------
	action_rev := theme.AddColorMode("action-rev")

	// highlight
	action_rev_highlight := action_rev.AddRole("content-highlight")
	action_rev_highlight.Context = "foreground"
	action_rev_highlight.AddState("default", color.UIColor{Light: azimuth.At(0), Dark: azimuth.At(0.1)})
	action_rev_highlight.AddState("hover", color.UIColor{Light: azimuth.At(0.05), Dark: azimuth.At(0)})
	action_rev_highlight.AddState("pressed", color.UIColor{Light: azimuth.At(0.05), Dark: azimuth.At(0.05)})
	action_rev_highlight.AddState("focus", color.UIColor{Light: azimuth.At(0.05), Dark: azimuth.At(0.05)})
	// content
	action_rev_content := action_rev.AddRole("content")
	action_rev_content.Context = "foreground"
	action_rev_content.AddState("default", color.UIColor{Light: azimuth.At(0.08), Dark: azimuth.At(0.15)})
	action_rev_content.AddState("hover", color.UIColor{Light: azimuth.At(0.15), Dark: azimuth.At(0.15)})
	action_rev_content.AddState("pressed", color.UIColor{Light: azimuth.At(0.15), Dark: azimuth.At(0.15)})
	action_rev_content.AddState("focus", color.UIColor{Light: azimuth.At(0.15), Dark: azimuth.At(0.15)})
	// secondary
	action_rev_secondary := action_rev.AddRole("content-secondary")
	action_rev_secondary.Context = "foreground"
	action_rev_secondary.AddState("default", color.UIColor{Light: azimuth.At(0.15), Dark: azimuth.At(0.2)})
	action_rev_secondary.AddState("hover", color.UIColor{Light: azimuth.At(0.2), Dark: azimuth.At(0.2)})
	action_rev_secondary.AddState("pressed", color.UIColor{Light: azimuth.At(0.2), Dark: azimuth.At(0.2)})
	action_rev_secondary.AddState("focus", color.UIColor{Light: azimuth.At(0.2), Dark: azimuth.At(0.2)})
	// trim
	action_rev_trim := action_rev.AddRole("trim")
	action_rev_trim.AddState("default", color.UIColor{Light: azimuth.At(0.46875), Dark: azimuth.At(0.46875)})
	action_rev_trim.AddState("hover", color.UIColor{Light: azimuth.At(0.46875), Dark: azimuth.At(0.46875)})
	action_rev_trim.AddState("pressed", color.UIColor{Light: azimuth.At(0.46875), Dark: azimuth.At(0.46875)})
	action_rev_trim.AddState("focus", color.UIColor{Light: azimuth.At(0.46875), Dark: azimuth.At(0.46875)})
	// surface
	action_rev_surface := action_rev.AddRole("surface")
	action_rev_surface.Context = "background"
	action_rev_surface.AddState("default", color.UIColor{Light: azimuth.At(0.4375), Dark: azimuth.At(0.4375)})
	action_rev_surface.AddState("hover", color.UIColor{Light: azimuth.At(0.35), Dark: azimuth.At(0.35)})
	action_rev_surface.AddState("pressed", color.UIColor{Light: azimuth.At(0.35), Dark: azimuth.At(0.35)})
	action_rev_surface.AddState("focus", color.UIColor{Light: azimuth.At(0.35), Dark: azimuth.At(0.35)})
	// surface_low
	action_rev_surface_low := action_rev.AddRole("surface-low")
	action_rev_surface_low.Context = "background"
	action_rev_surface_low.AddState("default", color.UIColor{Light: azimuth.At(0.5), Dark: azimuth.At(0.5375)})
	action_rev_surface_low.AddState("hover", color.UIColor{Light: azimuth.At(0.5), Dark: azimuth.At(0.95)})
	action_rev_surface_low.AddState("pressed", color.UIColor{Light: azimuth.At(0.5), Dark: azimuth.At(0.95)})
	action_rev_surface_low.AddState("focus", color.UIColor{Light: azimuth.At(0.5), Dark: azimuth.At(0.95)})
	// surface_high
	action_rev_surface_high := action_rev.AddRole("surface-high")
	action_rev_surface_high.Context = "background"
	action_rev_surface_high.AddState("default", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.6)})
	action_rev_surface_high.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.88)})
	action_rev_surface_high.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.88)})
	action_rev_surface_high.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.88)})
	// background
	action_rev_background := action_rev.AddRole("surface-bg")
	action_rev_background.Context = "background"
	action_rev_background.AddState("default", color.UIColor{Light: azimuth.At(0.55), Dark: azimuth.At(0.96)})
	action_rev_background.AddState("hover", color.UIColor{Light: azimuth.At(0.8), Dark: azimuth.At(0.96)})
	action_rev_background.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.96)})
	action_rev_background.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.96)})

	//
	// critical REVERSED
	// Light on dark colors for both light and dark mode
	// ----------------
	critical_rev := theme.AddColorMode("critical-rev")

	// highlight
	critical_rev_highlight := critical_rev.AddRole("content-highlight")
	critical_rev_highlight.Context = "foreground"
	critical_rev_highlight.AddState("default", color.UIColor{Light: intlOrange.At(0), Dark: intlOrange.At(0.1)})
	critical_rev_highlight.AddState("hover", color.UIColor{Light: intlOrange.At(0.05), Dark: intlOrange.At(0)})
	critical_rev_highlight.AddState("pressed", color.UIColor{Light: intlOrange.At(0.05), Dark: intlOrange.At(0.05)})
	critical_rev_highlight.AddState("focus", color.UIColor{Light: intlOrange.At(0.05), Dark: intlOrange.At(0.05)})
	// content
	critical_rev_content := critical_rev.AddRole("content")
	critical_rev_content.Context = "foreground"
	critical_rev_content.AddState("default", color.UIColor{Light: intlOrange.At(0.08), Dark: intlOrange.At(0.15)})
	critical_rev_content.AddState("hover", color.UIColor{Light: intlOrange.At(0.15), Dark: intlOrange.At(0.15)})
	critical_rev_content.AddState("pressed", color.UIColor{Light: intlOrange.At(0.15), Dark: intlOrange.At(0.15)})
	critical_rev_content.AddState("focus", color.UIColor{Light: intlOrange.At(0.15), Dark: intlOrange.At(0.15)})
	// secondary
	critical_rev_secondary := critical_rev.AddRole("content-secondary")
	critical_rev_secondary.Context = "foreground"
	critical_rev_secondary.AddState("default", color.UIColor{Light: intlOrange.At(0.15), Dark: intlOrange.At(0.2)})
	critical_rev_secondary.AddState("hover", color.UIColor{Light: intlOrange.At(0.2), Dark: intlOrange.At(0.2)})
	critical_rev_secondary.AddState("pressed", color.UIColor{Light: intlOrange.At(0.2), Dark: intlOrange.At(0.2)})
	critical_rev_secondary.AddState("focus", color.UIColor{Light: intlOrange.At(0.2), Dark: intlOrange.At(0.2)})
	// trim
	critical_rev_trim := critical_rev.AddRole("trim")
	critical_rev_trim.AddState("default", color.UIColor{Light: intlOrange.At(0.46875), Dark: intlOrange.At(0.46875)})
	critical_rev_trim.AddState("hover", color.UIColor{Light: intlOrange.At(0.46875), Dark: intlOrange.At(0.46875)})
	critical_rev_trim.AddState("pressed", color.UIColor{Light: intlOrange.At(0.46875), Dark: intlOrange.At(0.46875)})
	critical_rev_trim.AddState("focus", color.UIColor{Light: intlOrange.At(0.46875), Dark: intlOrange.At(0.46875)})
	// surface
	critical_rev_surface := critical_rev.AddRole("surface")
	critical_rev_surface.Context = "background"
	critical_rev_surface.AddState("default", color.UIColor{Light: intlOrange.At(0.4375), Dark: intlOrange.At(0.4375)})
	critical_rev_surface.AddState("hover", color.UIColor{Light: intlOrange.At(0.35), Dark: intlOrange.At(0.35)})
	critical_rev_surface.AddState("pressed", color.UIColor{Light: intlOrange.At(0.35), Dark: intlOrange.At(0.35)})
	critical_rev_surface.AddState("focus", color.UIColor{Light: intlOrange.At(0.35), Dark: intlOrange.At(0.35)})
	// surface_low
	critical_rev_surface_low := critical_rev.AddRole("surface-low")
	critical_rev_surface_low.Context = "background"
	critical_rev_surface_low.AddState("default", color.UIColor{Light: intlOrange.At(0.5), Dark: intlOrange.At(0.5375)})
	critical_rev_surface_low.AddState("hover", color.UIColor{Light: intlOrange.At(0.5), Dark: intlOrange.At(0.95)})
	critical_rev_surface_low.AddState("pressed", color.UIColor{Light: intlOrange.At(0.5), Dark: intlOrange.At(0.95)})
	critical_rev_surface_low.AddState("focus", color.UIColor{Light: intlOrange.At(0.5), Dark: intlOrange.At(0.95)})
	// surface_high
	critical_rev_surface_high := critical_rev.AddRole("surface-high")
	critical_rev_surface_high.Context = "background"
	critical_rev_surface_high.AddState("default", color.UIColor{Light: intlOrange.At(0.6), Dark: intlOrange.At(0.6)})
	critical_rev_surface_high.AddState("hover", color.UIColor{Light: intlOrange.At(0.8), Dark: intlOrange.At(0.88)})
	critical_rev_surface_high.AddState("pressed", color.UIColor{Light: intlOrange.At(0.7), Dark: intlOrange.At(0.88)})
	critical_rev_surface_high.AddState("focus", color.UIColor{Light: intlOrange.At(0.6), Dark: intlOrange.At(0.88)})
	// background
	critical_rev_background := critical_rev.AddRole("surface-bg")
	critical_rev_background.Context = "background"
	critical_rev_background.AddState("default", color.UIColor{Light: intlOrange.At(0.55), Dark: intlOrange.At(0.96)})
	critical_rev_background.AddState("hover", color.UIColor{Light: intlOrange.At(0.8), Dark: intlOrange.At(0.96)})
	critical_rev_background.AddState("pressed", color.UIColor{Light: intlOrange.At(0.7), Dark: intlOrange.At(0.96)})
	critical_rev_background.AddState("focus", color.UIColor{Light: intlOrange.At(0.6), Dark: intlOrange.At(0.96)})

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
