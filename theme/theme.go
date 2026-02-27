package theme

import "go-ds/color"

type Theme struct {
	ColorModes []color.Mode
}

func NewTheme() Theme {
	return Theme{}
}

func (t *Theme) AddColorMode(mode color.Mode) {
	t.ColorModes = append(t.ColorModes, mode)
}

func (t *Theme) ToCSS() string {
	css := ":root {\n"
	for _, mode := range t.ColorModes {
		css += mode.ToCSS()
	}
	css += "}"
	return css
}
