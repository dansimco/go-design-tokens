package theme

import (
	"go-ds/color"
	"os"
	"testing"
)

func TestThemeCSS(t *testing.T) {
	theme := NewTheme()

	theme.SpacePrefix = "sp-"
	theme.AddSpaceToken("2xs", 0.25)
	theme.AddSpaceToken("xs", 0.5)
	theme.AddSpaceToken("s", 1)
	theme.AddSpaceToken("m", 1.5)
	theme.AddSpaceToken("l", 2)
	theme.AddSpaceToken("xl", 3)
	theme.AddSpaceToken("2xl", 4)

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

	electric := color.NewRamp()
	electric.AddKey("#FFFFD8", 0)
	electric.AddKey("#F8FF6C", 0.375)
	electric.AddKey("#080a00", 1)

	kelp := color.NewRamp()
	kelp.AddKey("#EBFFE0", 0)
	kelp.AddKey("#6E9B55", 0.53)
	kelp.AddKey("#000000", 1)

	seaFoam := color.NewRamp()
	seaFoam.AddKey("#E2FFFF", 0)
	seaFoam.AddKey("#7AD0D3", 0.375)
	seaFoam.AddKey("#010511", 1)

	patrick := color.NewRamp()
	patrick.AddKey("#FFFFFF", 0)
	patrick.AddKey("#DD79D8", 0.4)
	patrick.AddKey("#05000f", 1)

	//
	// neutral
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
	// action
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
	action_trim.AddState("default", color.UIColor{Light: azimuth.At(0.1), Dark: azimuth.At(0.8125)})
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
	// action reversed
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
	action_rev_trim.AddState("default", color.UIColor{Light: azimuth.At(0.43), Dark: azimuth.At(0.46875)})
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
	// action
	// ----------------
	critical := theme.AddColorMode("critical")

	// highlight
	critical_highlight := critical.AddRole("content-highlight")
	critical_highlight.Context = "foreground"
	critical_highlight.AddState("default", color.UIColor{Light: intlOrange.At(0.5), Dark: intlOrange.At(0.3125)})
	critical_highlight.AddState("hover", color.UIColor{Light: intlOrange.At(0.4), Dark: intlOrange.At(0.2)})
	critical_highlight.AddState("pressed", color.UIColor{Light: intlOrange.At(0.4), Dark: intlOrange.At(0.15)})
	critical_highlight.AddState("focus", color.UIColor{Light: intlOrange.At(0.4), Dark: intlOrange.At(0.2)})
	// content
	critical_content := critical.AddRole("content")
	critical_content.Context = "foreground"
	critical_content.AddState("default", color.UIColor{Light: intlOrange.At(0.4), Dark: intlOrange.At(0.4375)})
	critical_content.AddState("hover", color.UIColor{Light: intlOrange.At(0.37), Dark: intlOrange.At(0.35)})
	critical_content.AddState("pressed", color.UIColor{Light: intlOrange.At(0.7), Dark: intlOrange.At(0.3)})
	critical_content.AddState("focus", color.UIColor{Light: intlOrange.At(0.6), Dark: intlOrange.At(0.35)})
	// secondary
	critical_secondary := critical.AddRole("content-secondary")
	critical_secondary.Context = "foreground"
	critical_secondary.AddState("default", color.UIColor{Light: intlOrange.At(0.33), Dark: intlOrange.At(0.58)})
	critical_secondary.AddState("hover", color.UIColor{Light: intlOrange.At(0.3), Dark: intlOrange.At(0.55)})
	critical_secondary.AddState("pressed", color.UIColor{Light: intlOrange.At(0.3), Dark: intlOrange.At(0.5)})
	critical_secondary.AddState("focus", color.UIColor{Light: intlOrange.At(0.3), Dark: intlOrange.At(0.55)})
	// trim
	critical_trim := critical.AddRole("trim")
	critical_trim.AddState("default", color.UIColor{Light: intlOrange.At(0.1), Dark: intlOrange.At(0.8125)})
	critical_trim.AddState("hover", color.UIColor{Light: intlOrange.At(0.8), Dark: intlOrange.At(0.7)})
	critical_trim.AddState("pressed", color.UIColor{Light: intlOrange.At(0.7), Dark: intlOrange.At(0.7)})
	critical_trim.AddState("focus", color.UIColor{Light: intlOrange.At(0.6), Dark: intlOrange.At(0.7)})
	// surface
	critical_surface := critical.AddRole("surface")
	critical_surface.Context = "background"
	critical_surface.AddState("default", color.UIColor{Light: ink.At(0.01), Dark: ink.At(0.975)})
	critical_surface.AddState("hover", color.UIColor{Light: intlOrange.At(0.03), Dark: intlOrange.At(0.9)})
	critical_surface.AddState("pressed", color.UIColor{Light: intlOrange.At(0.7), Dark: intlOrange.At(0.9)})
	critical_surface.AddState("focus", color.UIColor{Light: intlOrange.At(0.6), Dark: intlOrange.At(0.9)})
	// surface_low
	critical_surface_low := critical.AddRole("surface-low")
	critical_surface_low.Context = "background"
	critical_surface_low.AddState("default", color.UIColor{Light: intlOrange.At(0.025), Dark: intlOrange.At(0.95)})
	critical_surface_low.AddState("hover", color.UIColor{Light: intlOrange.At(0.9), Dark: intlOrange.At(0.95)})
	critical_surface_low.AddState("pressed", color.UIColor{Light: intlOrange.At(0.7), Dark: intlOrange.At(0.95)})
	critical_surface_low.AddState("focus", color.UIColor{Light: intlOrange.At(0.6), Dark: intlOrange.At(0.95)})
	// surface_high
	critical_surface_high := critical.AddRole("surface-high")
	critical_surface_high.Context = "background"
	critical_surface_high.AddState("default", color.UIColor{Light: intlOrange.At(0.05), Dark: intlOrange.At(0.88)})
	critical_surface_high.AddState("hover", color.UIColor{Light: intlOrange.At(0.8), Dark: intlOrange.At(0.88)})
	critical_surface_high.AddState("pressed", color.UIColor{Light: intlOrange.At(0.7), Dark: intlOrange.At(0.88)})
	critical_surface_high.AddState("focus", color.UIColor{Light: intlOrange.At(0.6), Dark: intlOrange.At(0.88)})
	// background
	critical_background := critical.AddRole("surface-bg")
	critical_background.Context = "background"
	critical_background.AddState("default", color.UIColor{Light: intlOrange.At(0.02), Dark: intlOrange.At(0.96)})
	critical_background.AddState("hover", color.UIColor{Light: intlOrange.At(0.8), Dark: intlOrange.At(0.96)})
	critical_background.AddState("pressed", color.UIColor{Light: intlOrange.At(0.7), Dark: intlOrange.At(0.96)})
	critical_background.AddState("focus", color.UIColor{Light: intlOrange.At(0.6), Dark: intlOrange.At(0.96)})

	//
	// critical reversed
	// light on dark colors for both light and dark mode
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
	critical_rev_trim.AddState("default", color.UIColor{Light: intlOrange.At(0.46875), Dark: intlOrange.At(0.7)})
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

	//
	// tint 1
	// ----------------
	tint1 := theme.AddColorMode("tint-1")

	// highlight
	tint1_highlight := tint1.AddRole("content-highlight")
	tint1_highlight.Context = "foreground"
	tint1_highlight.AddState("default", color.UIColor{Light: electric.At(0.8), Dark: electric.At(0.3125)})
	tint1_highlight.AddState("hover", color.UIColor{Light: electric.At(0.75), Dark: electric.At(0.2)})
	tint1_highlight.AddState("pressed", color.UIColor{Light: electric.At(0.4), Dark: electric.At(0.15)})
	tint1_highlight.AddState("focus", color.UIColor{Light: electric.At(0.4), Dark: electric.At(0.2)})
	// content
	tint1_content := tint1.AddRole("content")
	tint1_content.Context = "foreground"
	tint1_content.AddState("default", color.UIColor{Light: electric.At(0.75), Dark: electric.At(0.4375)})
	tint1_content.AddState("hover", color.UIColor{Light: electric.At(0.37), Dark: electric.At(0.35)})
	tint1_content.AddState("pressed", color.UIColor{Light: electric.At(0.7), Dark: electric.At(0.3)})
	tint1_content.AddState("focus", color.UIColor{Light: electric.At(0.6), Dark: electric.At(0.35)})
	// secondary
	tint1_secondary := tint1.AddRole("content-secondary")
	tint1_secondary.Context = "foreground"
	tint1_secondary.AddState("default", color.UIColor{Light: electric.At(0.65), Dark: electric.At(0.58)})
	tint1_secondary.AddState("hover", color.UIColor{Light: electric.At(0.3), Dark: electric.At(0.55)})
	tint1_secondary.AddState("pressed", color.UIColor{Light: electric.At(0.3), Dark: electric.At(0.5)})
	tint1_secondary.AddState("focus", color.UIColor{Light: electric.At(0.3), Dark: electric.At(0.55)})
	// trim
	tint1_trim := tint1.AddRole("trim")
	tint1_trim.AddState("default", color.UIColor{Light: electric.At(0.41), Dark: electric.At(0.9)})
	tint1_trim.AddState("hover", color.UIColor{Light: electric.At(0.8), Dark: electric.At(0.7)})
	tint1_trim.AddState("pressed", color.UIColor{Light: electric.At(0.7), Dark: electric.At(0.7)})
	tint1_trim.AddState("focus", color.UIColor{Light: electric.At(0.6), Dark: electric.At(0.7)})
	// surface
	tint1_surface := tint1.AddRole("surface")
	tint1_surface.Context = "background"
	tint1_surface.AddState("default", color.UIColor{Light: electric.At(0), Dark: electric.At(0.95)})
	tint1_surface.AddState("hover", color.UIColor{Light: electric.At(0.15), Dark: electric.At(0.9)})
	tint1_surface.AddState("pressed", color.UIColor{Light: electric.At(0.15), Dark: electric.At(0.9)})
	tint1_surface.AddState("focus", color.UIColor{Light: electric.At(0.15), Dark: electric.At(0.9)})
	// surface_low
	tint1_surface_low := tint1.AddRole("surface-low")
	tint1_surface_low.Context = "background"
	tint1_surface_low.AddState("default", color.UIColor{Light: electric.At(0.1), Dark: electric.At(0.96)})
	tint1_surface_low.AddState("hover", color.UIColor{Light: electric.At(0.25), Dark: electric.At(0.95)})
	tint1_surface_low.AddState("pressed", color.UIColor{Light: electric.At(0.25), Dark: electric.At(0.95)})
	tint1_surface_low.AddState("focus", color.UIColor{Light: electric.At(0.25), Dark: electric.At(0.95)})
	// surface_high
	tint1_surface_high := tint1.AddRole("surface-high")
	tint1_surface_high.Context = "background"
	tint1_surface_high.AddState("default", color.UIColor{Light: electric.At(0.2), Dark: electric.At(0.97)})
	tint1_surface_high.AddState("hover", color.UIColor{Light: electric.At(0.35), Dark: electric.At(0.88)})
	tint1_surface_high.AddState("pressed", color.UIColor{Light: electric.At(0.35), Dark: electric.At(0.88)})
	tint1_surface_high.AddState("focus", color.UIColor{Light: electric.At(0.35), Dark: electric.At(0.88)})
	// background
	tint1_background := tint1.AddRole("surface-bg")
	tint1_background.Context = "background"
	tint1_background.AddState("default", color.UIColor{Light: electric.At(0.3), Dark: electric.At(0.98)})
	tint1_background.AddState("hover", color.UIColor{Light: electric.At(0.37), Dark: electric.At(0.96)})
	tint1_background.AddState("pressed", color.UIColor{Light: electric.At(0.37), Dark: electric.At(0.96)})
	tint1_background.AddState("focus", color.UIColor{Light: electric.At(0.37), Dark: electric.At(0.96)})

	//
	// tint1 reversed
	// light on dark colors for both light and dark mode
	// ----------------
	tint1_rev := theme.AddColorMode("tint-1-rev")

	// highlight
	tint1_rev_highlight := tint1_rev.AddRole("content-highlight")
	tint1_rev_highlight.Context = "foreground"
	tint1_rev_highlight.AddState("default", color.UIColor{Light: electric.At(0.85), Dark: electric.At(0.9)})
	tint1_rev_highlight.AddState("hover", color.UIColor{Light: electric.At(1), Dark: electric.At(0)})
	tint1_rev_highlight.AddState("pressed", color.UIColor{Light: electric.At(1), Dark: electric.At(0.05)})
	tint1_rev_highlight.AddState("focus", color.UIColor{Light: electric.At(1), Dark: electric.At(0.05)})
	// content
	tint1_rev_content := tint1_rev.AddRole("content")
	tint1_rev_content.Context = "foreground"
	tint1_rev_content.AddState("default", color.UIColor{Light: electric.At(0.78), Dark: electric.At(0.78)})
	tint1_rev_content.AddState("hover", color.UIColor{Light: electric.At(0.9), Dark: electric.At(0.15)})
	tint1_rev_content.AddState("pressed", color.UIColor{Light: electric.At(0.9), Dark: electric.At(0.15)})
	tint1_rev_content.AddState("focus", color.UIColor{Light: electric.At(0.9), Dark: electric.At(0.15)})
	// secondary
	tint1_rev_secondary := tint1_rev.AddRole("content-secondary")
	tint1_rev_secondary.Context = "foreground"
	tint1_rev_secondary.AddState("default", color.UIColor{Light: electric.At(0.7), Dark: electric.At(0.7)})
	tint1_rev_secondary.AddState("hover", color.UIColor{Light: electric.At(0.86), Dark: electric.At(0.2)})
	tint1_rev_secondary.AddState("pressed", color.UIColor{Light: electric.At(0.86), Dark: electric.At(0.2)})
	tint1_rev_secondary.AddState("focus", color.UIColor{Light: electric.At(0.86), Dark: electric.At(0.2)})
	// trim
	tint1_rev_trim := tint1_rev.AddRole("trim")
	tint1_rev_trim.AddState("default", color.UIColor{Light: electric.At(0.46875), Dark: electric.At(0.66)})
	tint1_rev_trim.AddState("hover", color.UIColor{Light: electric.At(0.46875), Dark: electric.At(0.46875)})
	tint1_rev_trim.AddState("pressed", color.UIColor{Light: electric.At(0.46875), Dark: electric.At(0.46875)})
	tint1_rev_trim.AddState("focus", color.UIColor{Light: electric.At(0.46875), Dark: electric.At(0.46875)})
	// surface
	tint1_rev_surface := tint1_rev.AddRole("surface")
	tint1_rev_surface.Context = "background"
	tint1_rev_surface.AddState("default", color.UIColor{Light: electric.At(0.3), Dark: electric.At(0.4375)})
	tint1_rev_surface.AddState("hover", color.UIColor{Light: electric.At(0.35), Dark: electric.At(0.35)})
	tint1_rev_surface.AddState("pressed", color.UIColor{Light: electric.At(0.35), Dark: electric.At(0.35)})
	tint1_rev_surface.AddState("focus", color.UIColor{Light: electric.At(0.35), Dark: electric.At(0.35)})
	// surface_low
	tint1_rev_surface_low := tint1_rev.AddRole("surface-low")
	tint1_rev_surface_low.Context = "background"
	tint1_rev_surface_low.AddState("default", color.UIColor{Light: electric.At(0.4), Dark: electric.At(0.46)})
	tint1_rev_surface_low.AddState("hover", color.UIColor{Light: electric.At(0.95), Dark: electric.At(0.95)})
	tint1_rev_surface_low.AddState("pressed", color.UIColor{Light: electric.At(0.95), Dark: electric.At(0.95)})
	tint1_rev_surface_low.AddState("focus", color.UIColor{Light: electric.At(0.95), Dark: electric.At(0.95)})
	// surface_high
	tint1_rev_surface_high := tint1_rev.AddRole("surface-high")
	tint1_rev_surface_high.Context = "background"
	tint1_rev_surface_high.AddState("default", color.UIColor{Light: electric.At(0.43), Dark: electric.At(0.5)})
	tint1_rev_surface_high.AddState("hover", color.UIColor{Light: electric.At(0.8), Dark: electric.At(0.88)})
	tint1_rev_surface_high.AddState("pressed", color.UIColor{Light: electric.At(0.7), Dark: electric.At(0.88)})
	tint1_rev_surface_high.AddState("focus", color.UIColor{Light: electric.At(0.6), Dark: electric.At(0.88)})
	// background
	tint1_rev_background := tint1_rev.AddRole("surface-bg")
	tint1_rev_background.Context = "background"
	tint1_rev_background.AddState("default", color.UIColor{Light: electric.At(0.5), Dark: electric.At(0.55)})
	tint1_rev_background.AddState("hover", color.UIColor{Light: electric.At(0.96), Dark: electric.At(0.96)})
	tint1_rev_background.AddState("pressed", color.UIColor{Light: electric.At(0.96), Dark: electric.At(0.96)})
	tint1_rev_background.AddState("focus", color.UIColor{Light: electric.At(0.96), Dark: electric.At(0.96)})

	//
	// tint2
	//
	// ----------------
	tint2 := theme.AddColorMode("tint-2")

	// highlight
	tint2_highlight := tint2.AddRole("content-highlight")
	tint2_highlight.Context = "foreground"
	tint2_highlight.AddState("default", color.UIColor{Light: patrick.At(1), Dark: patrick.At(0.4)})
	tint2_highlight.AddState("hover", color.UIColor{Light: patrick.At(1), Dark: patrick.At(0.04)})
	tint2_highlight.AddState("pressed", color.UIColor{Light: patrick.At(1), Dark: patrick.At(0.05)})
	tint2_highlight.AddState("focus", color.UIColor{Light: patrick.At(1), Dark: patrick.At(0.05)})
	// content
	tint2_content := tint2.AddRole("content")
	tint2_content.Context = "foreground"
	tint2_content.AddState("default", color.UIColor{Light: patrick.At(0.9), Dark: patrick.At(0.1)})
	tint2_content.AddState("hover", color.UIColor{Light: patrick.At(0.9), Dark: patrick.At(0.15)})
	tint2_content.AddState("pressed", color.UIColor{Light: patrick.At(0.9), Dark: patrick.At(0.15)})
	tint2_content.AddState("focus", color.UIColor{Light: patrick.At(0.9), Dark: patrick.At(0.15)})
	// secondary
	tint2_secondary := tint2.AddRole("content-secondary")
	tint2_secondary.Context = "foreground"
	tint2_secondary.AddState("default", color.UIColor{Light: patrick.At(0.86), Dark: patrick.At(0.15)})
	tint2_secondary.AddState("hover", color.UIColor{Light: patrick.At(0.86), Dark: patrick.At(0.2)})
	tint2_secondary.AddState("pressed", color.UIColor{Light: patrick.At(0.86), Dark: patrick.At(0.2)})
	tint2_secondary.AddState("focus", color.UIColor{Light: patrick.At(0.86), Dark: patrick.At(0.2)})
	// trim
	tint2_trim := tint2.AddRole("trim")
	tint2_trim.AddState("default", color.UIColor{Light: patrick.At(0.1), Dark: patrick.At(0.9)})
	tint2_trim.AddState("hover", color.UIColor{Light: patrick.At(0.46875), Dark: patrick.At(0.46875)})
	tint2_trim.AddState("pressed", color.UIColor{Light: patrick.At(0.46875), Dark: patrick.At(0.46875)})
	tint2_trim.AddState("focus", color.UIColor{Light: patrick.At(0.46875), Dark: patrick.At(0.46875)})
	// surface
	tint2_surface := tint2.AddRole("surface")
	tint2_surface.Context = "background"
	tint2_surface.AddState("default", color.UIColor{Light: patrick.At(0.01), Dark: patrick.At(0.92)})
	tint2_surface.AddState("hover", color.UIColor{Light: patrick.At(0.35), Dark: patrick.At(0.35)})
	tint2_surface.AddState("pressed", color.UIColor{Light: patrick.At(0.35), Dark: patrick.At(0.35)})
	tint2_surface.AddState("focus", color.UIColor{Light: patrick.At(0.35), Dark: patrick.At(0.35)})
	// surface_low
	tint2_surface_low := tint2.AddRole("surface-low")
	tint2_surface_low.Context = "background"
	tint2_surface_low.AddState("default", color.UIColor{Light: patrick.At(0.05), Dark: patrick.At(0.93)})
	tint2_surface_low.AddState("hover", color.UIColor{Light: patrick.At(0.95), Dark: patrick.At(0.95)})
	tint2_surface_low.AddState("pressed", color.UIColor{Light: patrick.At(0.95), Dark: patrick.At(0.95)})
	tint2_surface_low.AddState("focus", color.UIColor{Light: patrick.At(0.95), Dark: patrick.At(0.95)})
	// surface_high
	tint2_surface_high := tint2.AddRole("surface-high")
	tint2_surface_high.Context = "background"
	tint2_surface_high.AddState("default", color.UIColor{Light: patrick.At(0.1), Dark: patrick.At(0.94)})
	tint2_surface_high.AddState("hover", color.UIColor{Light: patrick.At(0.8), Dark: patrick.At(0.88)})
	tint2_surface_high.AddState("pressed", color.UIColor{Light: patrick.At(0.7), Dark: patrick.At(0.88)})
	tint2_surface_high.AddState("focus", color.UIColor{Light: patrick.At(0.6), Dark: patrick.At(0.88)})
	// background
	tint2_background := tint2.AddRole("surface-bg")
	tint2_background.Context = "background"
	tint2_background.AddState("default", color.UIColor{Light: patrick.At(0.15), Dark: patrick.At(0.95)})
	tint2_background.AddState("hover", color.UIColor{Light: patrick.At(0.96), Dark: patrick.At(0.96)})
	tint2_background.AddState("pressed", color.UIColor{Light: patrick.At(0.96), Dark: patrick.At(0.96)})
	tint2_background.AddState("focus", color.UIColor{Light: patrick.At(0.96), Dark: patrick.At(0.96)})

	//
	// tint2 reversed
	// light on dark colors for both light and dark mode
	// ----------------
	tint2_rev := theme.AddColorMode("tint-2-rev")

	// highlight
	tint2_rev_highlight := tint2_rev.AddRole("content-highlight")
	tint2_rev_highlight.Context = "foreground"
	tint2_rev_highlight.AddState("default", color.UIColor{Light: patrick.At(0), Dark: patrick.At(1)})
	tint2_rev_highlight.AddState("hover", color.UIColor{Light: patrick.At(1), Dark: patrick.At(0)})
	tint2_rev_highlight.AddState("pressed", color.UIColor{Light: patrick.At(1), Dark: patrick.At(0.05)})
	tint2_rev_highlight.AddState("focus", color.UIColor{Light: patrick.At(1), Dark: patrick.At(0.05)})
	// content
	tint2_rev_content := tint2_rev.AddRole("content")
	tint2_rev_content.Context = "foreground"
	tint2_rev_content.AddState("default", color.UIColor{Light: patrick.At(0.05), Dark: patrick.At(0.85)})
	tint2_rev_content.AddState("hover", color.UIColor{Light: patrick.At(0.03), Dark: patrick.At(0.15)})
	tint2_rev_content.AddState("pressed", color.UIColor{Light: patrick.At(0.03), Dark: patrick.At(0.15)})
	tint2_rev_content.AddState("focus", color.UIColor{Light: patrick.At(0.03), Dark: patrick.At(0.15)})
	// secondary
	tint2_rev_secondary := tint2_rev.AddRole("content-secondary")
	tint2_rev_secondary.Context = "foreground"
	tint2_rev_secondary.AddState("default", color.UIColor{Light: patrick.At(0.14), Dark: patrick.At(0.8)})
	tint2_rev_secondary.AddState("hover", color.UIColor{Light: patrick.At(0.1), Dark: patrick.At(0.2)})
	tint2_rev_secondary.AddState("pressed", color.UIColor{Light: patrick.At(0.1), Dark: patrick.At(0.2)})
	tint2_rev_secondary.AddState("focus", color.UIColor{Light: patrick.At(0.1), Dark: patrick.At(0.2)})
	// trim
	tint2_rev_trim := tint2_rev.AddRole("trim")
	tint2_rev_trim.AddState("default", color.UIColor{Light: patrick.At(0.34), Dark: patrick.At(0.46875)})
	tint2_rev_trim.AddState("hover", color.UIColor{Light: patrick.At(0.46875), Dark: patrick.At(0.46875)})
	tint2_rev_trim.AddState("pressed", color.UIColor{Light: patrick.At(0.46875), Dark: patrick.At(0.46875)})
	tint2_rev_trim.AddState("focus", color.UIColor{Light: patrick.At(0.46875), Dark: patrick.At(0.46875)})
	// surface
	tint2_rev_surface := tint2_rev.AddRole("surface")
	tint2_rev_surface.Context = "background"
	tint2_rev_surface.AddState("default", color.UIColor{Light: patrick.At(0.4), Dark: patrick.At(0.4)})
	tint2_rev_surface.AddState("hover", color.UIColor{Light: patrick.At(0.35), Dark: patrick.At(0.35)})
	tint2_rev_surface.AddState("pressed", color.UIColor{Light: patrick.At(0.35), Dark: patrick.At(0.35)})
	tint2_rev_surface.AddState("focus", color.UIColor{Light: patrick.At(0.35), Dark: patrick.At(0.35)})
	// surface_low
	tint2_rev_surface_low := tint2_rev.AddRole("surface-low")
	tint2_rev_surface_low.Context = "background"
	tint2_rev_surface_low.AddState("default", color.UIColor{Light: patrick.At(0.44), Dark: patrick.At(0.42)})
	tint2_rev_surface_low.AddState("hover", color.UIColor{Light: patrick.At(0.95), Dark: patrick.At(0.95)})
	tint2_rev_surface_low.AddState("pressed", color.UIColor{Light: patrick.At(0.95), Dark: patrick.At(0.95)})
	tint2_rev_surface_low.AddState("focus", color.UIColor{Light: patrick.At(0.95), Dark: patrick.At(0.95)})
	// surface_high
	tint2_rev_surface_high := tint2_rev.AddRole("surface-high")
	tint2_rev_surface_high.Context = "background"
	tint2_rev_surface_high.AddState("default", color.UIColor{Light: patrick.At(0.48), Dark: patrick.At(0.45)})
	tint2_rev_surface_high.AddState("hover", color.UIColor{Light: patrick.At(0.8), Dark: patrick.At(0.88)})
	tint2_rev_surface_high.AddState("pressed", color.UIColor{Light: patrick.At(0.7), Dark: patrick.At(0.88)})
	tint2_rev_surface_high.AddState("focus", color.UIColor{Light: patrick.At(0.6), Dark: patrick.At(0.88)})
	// background
	tint2_rev_background := tint2_rev.AddRole("surface-bg")
	tint2_rev_background.Context = "background"
	tint2_rev_background.AddState("default", color.UIColor{Light: patrick.At(0.5), Dark: patrick.At(0.48)})
	tint2_rev_background.AddState("hover", color.UIColor{Light: patrick.At(0.96), Dark: patrick.At(0.96)})
	tint2_rev_background.AddState("pressed", color.UIColor{Light: patrick.At(0.96), Dark: patrick.At(0.96)})
	tint2_rev_background.AddState("focus", color.UIColor{Light: patrick.At(0.96), Dark: patrick.At(0.96)})

	//
	// tint3
	//
	// ----------------
	tint3 := theme.AddColorMode("tint-3")

	// highlight
	tint3_highlight := tint3.AddRole("content-highlight")
	tint3_highlight.Context = "foreground"
	tint3_highlight.AddState("default", color.UIColor{Light: seaFoam.At(1), Dark: seaFoam.At(0.4)})
	tint3_highlight.AddState("hover", color.UIColor{Light: seaFoam.At(1), Dark: seaFoam.At(0.04)})
	tint3_highlight.AddState("pressed", color.UIColor{Light: seaFoam.At(1), Dark: seaFoam.At(0.05)})
	tint3_highlight.AddState("focus", color.UIColor{Light: seaFoam.At(1), Dark: seaFoam.At(0.05)})
	// content
	tint3_content := tint3.AddRole("content")
	tint3_content.Context = "foreground"
	tint3_content.AddState("default", color.UIColor{Light: seaFoam.At(0.9), Dark: seaFoam.At(0.1)})
	tint3_content.AddState("hover", color.UIColor{Light: seaFoam.At(0.9), Dark: seaFoam.At(0.15)})
	tint3_content.AddState("pressed", color.UIColor{Light: seaFoam.At(0.9), Dark: seaFoam.At(0.15)})
	tint3_content.AddState("focus", color.UIColor{Light: seaFoam.At(0.9), Dark: seaFoam.At(0.15)})
	// secondary
	tint3_secondary := tint3.AddRole("content-secondary")
	tint3_secondary.Context = "foreground"
	tint3_secondary.AddState("default", color.UIColor{Light: seaFoam.At(0.86), Dark: seaFoam.At(0.15)})
	tint3_secondary.AddState("hover", color.UIColor{Light: seaFoam.At(0.86), Dark: seaFoam.At(0.2)})
	tint3_secondary.AddState("pressed", color.UIColor{Light: seaFoam.At(0.86), Dark: seaFoam.At(0.2)})
	tint3_secondary.AddState("focus", color.UIColor{Light: seaFoam.At(0.86), Dark: seaFoam.At(0.2)})
	// trim
	tint3_trim := tint3.AddRole("trim")
	tint3_trim.AddState("default", color.UIColor{Light: seaFoam.At(0.1), Dark: seaFoam.At(0.9)})
	tint3_trim.AddState("hover", color.UIColor{Light: seaFoam.At(0.46875), Dark: seaFoam.At(0.46875)})
	tint3_trim.AddState("pressed", color.UIColor{Light: seaFoam.At(0.46875), Dark: seaFoam.At(0.46875)})
	tint3_trim.AddState("focus", color.UIColor{Light: seaFoam.At(0.46875), Dark: seaFoam.At(0.46875)})
	// surface
	tint3_surface := tint3.AddRole("surface")
	tint3_surface.Context = "background"
	tint3_surface.AddState("default", color.UIColor{Light: seaFoam.At(0.01), Dark: seaFoam.At(0.92)})
	tint3_surface.AddState("hover", color.UIColor{Light: seaFoam.At(0.35), Dark: seaFoam.At(0.35)})
	tint3_surface.AddState("pressed", color.UIColor{Light: seaFoam.At(0.35), Dark: seaFoam.At(0.35)})
	tint3_surface.AddState("focus", color.UIColor{Light: seaFoam.At(0.35), Dark: seaFoam.At(0.35)})
	// surface_low
	tint3_surface_low := tint3.AddRole("surface-low")
	tint3_surface_low.Context = "background"
	tint3_surface_low.AddState("default", color.UIColor{Light: seaFoam.At(0.05), Dark: seaFoam.At(0.93)})
	tint3_surface_low.AddState("hover", color.UIColor{Light: seaFoam.At(0.95), Dark: seaFoam.At(0.95)})
	tint3_surface_low.AddState("pressed", color.UIColor{Light: seaFoam.At(0.95), Dark: seaFoam.At(0.95)})
	tint3_surface_low.AddState("focus", color.UIColor{Light: seaFoam.At(0.95), Dark: seaFoam.At(0.95)})
	// surface_high
	tint3_surface_high := tint3.AddRole("surface-high")
	tint3_surface_high.Context = "background"
	tint3_surface_high.AddState("default", color.UIColor{Light: seaFoam.At(0.1), Dark: seaFoam.At(0.94)})
	tint3_surface_high.AddState("hover", color.UIColor{Light: seaFoam.At(0.8), Dark: seaFoam.At(0.88)})
	tint3_surface_high.AddState("pressed", color.UIColor{Light: seaFoam.At(0.7), Dark: seaFoam.At(0.88)})
	tint3_surface_high.AddState("focus", color.UIColor{Light: seaFoam.At(0.6), Dark: seaFoam.At(0.88)})
	// background
	tint3_background := tint3.AddRole("surface-bg")
	tint3_background.Context = "background"
	tint3_background.AddState("default", color.UIColor{Light: seaFoam.At(0.15), Dark: seaFoam.At(0.95)})
	tint3_background.AddState("hover", color.UIColor{Light: seaFoam.At(0.96), Dark: seaFoam.At(0.96)})
	tint3_background.AddState("pressed", color.UIColor{Light: seaFoam.At(0.96), Dark: seaFoam.At(0.96)})
	tint3_background.AddState("focus", color.UIColor{Light: seaFoam.At(0.96), Dark: seaFoam.At(0.96)})

	//
	// tint3 reversed
	// light on dark colors for both light and dark mode
	// ----------------
	tint3_rev := theme.AddColorMode("tint-3-rev")

	// highlight
	tint3_rev_highlight := tint3_rev.AddRole("content-highlight")
	tint3_rev_highlight.Context = "foreground"
	tint3_rev_highlight.AddState("default", color.UIColor{Light: seaFoam.At(0), Dark: seaFoam.At(1)})
	tint3_rev_highlight.AddState("hover", color.UIColor{Light: seaFoam.At(1), Dark: seaFoam.At(0)})
	tint3_rev_highlight.AddState("pressed", color.UIColor{Light: seaFoam.At(1), Dark: seaFoam.At(0.05)})
	tint3_rev_highlight.AddState("focus", color.UIColor{Light: seaFoam.At(1), Dark: seaFoam.At(0.05)})
	// content
	tint3_rev_content := tint3_rev.AddRole("content")
	tint3_rev_content.Context = "foreground"
	tint3_rev_content.AddState("default", color.UIColor{Light: seaFoam.At(0.08), Dark: seaFoam.At(0.9)})
	tint3_rev_content.AddState("hover", color.UIColor{Light: seaFoam.At(0.9), Dark: seaFoam.At(0.15)})
	tint3_rev_content.AddState("pressed", color.UIColor{Light: seaFoam.At(0.9), Dark: seaFoam.At(0.15)})
	tint3_rev_content.AddState("focus", color.UIColor{Light: seaFoam.At(0.9), Dark: seaFoam.At(0.15)})
	// secondary
	tint3_rev_secondary := tint3_rev.AddRole("content-secondary")
	tint3_rev_secondary.Context = "foreground"
	tint3_rev_secondary.AddState("default", color.UIColor{Light: seaFoam.At(0.2), Dark: seaFoam.At(0.86)})
	tint3_rev_secondary.AddState("hover", color.UIColor{Light: seaFoam.At(0.86), Dark: seaFoam.At(0.2)})
	tint3_rev_secondary.AddState("pressed", color.UIColor{Light: seaFoam.At(0.86), Dark: seaFoam.At(0.2)})
	tint3_rev_secondary.AddState("focus", color.UIColor{Light: seaFoam.At(0.86), Dark: seaFoam.At(0.2)})
	// trim
	tint3_rev_trim := tint3_rev.AddRole("trim")
	tint3_rev_trim.AddState("default", color.UIColor{Light: seaFoam.At(0.46875), Dark: seaFoam.At(0.46875)})
	tint3_rev_trim.AddState("hover", color.UIColor{Light: seaFoam.At(0.46875), Dark: seaFoam.At(0.46875)})
	tint3_rev_trim.AddState("pressed", color.UIColor{Light: seaFoam.At(0.46875), Dark: seaFoam.At(0.46875)})
	tint3_rev_trim.AddState("focus", color.UIColor{Light: seaFoam.At(0.46875), Dark: seaFoam.At(0.46875)})
	// surface
	tint3_rev_surface := tint3_rev.AddRole("surface")
	tint3_rev_surface.Context = "background"
	tint3_rev_surface.AddState("default", color.UIColor{Light: seaFoam.At(0.58), Dark: seaFoam.At(0.4375)})
	tint3_rev_surface.AddState("hover", color.UIColor{Light: seaFoam.At(0.35), Dark: seaFoam.At(0.35)})
	tint3_rev_surface.AddState("pressed", color.UIColor{Light: seaFoam.At(0.35), Dark: seaFoam.At(0.35)})
	tint3_rev_surface.AddState("focus", color.UIColor{Light: seaFoam.At(0.35), Dark: seaFoam.At(0.35)})
	// surface_low
	tint3_rev_surface_low := tint3_rev.AddRole("surface-low")
	tint3_rev_surface_low.Context = "background"
	tint3_rev_surface_low.AddState("default", color.UIColor{Light: seaFoam.At(0.6), Dark: seaFoam.At(0.46)})
	tint3_rev_surface_low.AddState("hover", color.UIColor{Light: seaFoam.At(0.95), Dark: seaFoam.At(0.95)})
	tint3_rev_surface_low.AddState("pressed", color.UIColor{Light: seaFoam.At(0.95), Dark: seaFoam.At(0.95)})
	tint3_rev_surface_low.AddState("focus", color.UIColor{Light: seaFoam.At(0.95), Dark: seaFoam.At(0.95)})
	// surface_high
	tint3_rev_surface_high := tint3_rev.AddRole("surface-high")
	tint3_rev_surface_high.Context = "background"
	tint3_rev_surface_high.AddState("default", color.UIColor{Light: seaFoam.At(0.64), Dark: seaFoam.At(0.6)})
	tint3_rev_surface_high.AddState("hover", color.UIColor{Light: seaFoam.At(0.8), Dark: seaFoam.At(0.88)})
	tint3_rev_surface_high.AddState("pressed", color.UIColor{Light: seaFoam.At(0.7), Dark: seaFoam.At(0.88)})
	tint3_rev_surface_high.AddState("focus", color.UIColor{Light: seaFoam.At(0.6), Dark: seaFoam.At(0.88)})
	// background
	tint3_rev_background := tint3_rev.AddRole("surface-bg")
	tint3_rev_background.Context = "background"
	tint3_rev_background.AddState("default", color.UIColor{Light: seaFoam.At(0.66), Dark: seaFoam.At(0.62)})
	tint3_rev_background.AddState("hover", color.UIColor{Light: seaFoam.At(0.96), Dark: seaFoam.At(0.96)})
	tint3_rev_background.AddState("pressed", color.UIColor{Light: seaFoam.At(0.96), Dark: seaFoam.At(0.96)})
	tint3_rev_background.AddState("focus", color.UIColor{Light: seaFoam.At(0.96), Dark: seaFoam.At(0.96)})

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
