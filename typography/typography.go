package typography

import (
	"fmt"
	"strings"

	"github.com/dansimco/go-design-tokens/css_util"
)

type Family struct {
	Name          string
	Fonts         []Font
	FallbackFonts []string
}

func NewFontFamily(name string) Family {
	return Family{
		Name: name,
	}
}

func (f *Family) AddFont() *Font {
	font := Font{
		Weight:            "normal",
		Style:             "normal",
		UseNumberedWeight: false,
	}
	f.Fonts = append(f.Fonts, font)
	return &f.Fonts[len(f.Fonts)-1]
}

func (f *Family) AddFallbackFont(font_name string) {
	f.FallbackFonts = append(f.FallbackFonts, font_name)
}

func (f *Family) ToCSS() string {
	css := "\n"

	for i, font := range f.Fonts {
		if i == 0 {
			css += "\t"
		} else {
			css += "\n   \t"
		}
		css += "@font-face {\n"
		css += "\t  font-family: \"" + f.Name + "\";\n"

		var srcEntries []string

		// Add local sources
		for _, local := range font.localSrc {
			srcEntries = append(srcEntries, "local(\""+local+"\")")
		}

		// Add URL sources
		for _, src := range font.src {
			// Replace any existing extension with .woff2
			if !strings.HasSuffix(src, ".woff2") {
				if idx := strings.LastIndex(src, "."); idx != -1 && idx > strings.LastIndex(src, "/") {
					src = src[:idx]
				}
				src += ".woff2"
			}
			srcEntries = append(srcEntries, "url(\""+src+"\")")
		}

		// Only emit src if there's at least one source, and always
		// terminate the last entry with a semicolon.
		if len(srcEntries) > 0 {
			css += "\t  src:\n"
			for j, entry := range srcEntries {
				css += "        " + entry
				if j < len(srcEntries)-1 {
					css += ","
				} else {
					css += ";"
				}
				css += "\n"
			}
		}

		if font.UseNumberedWeight && font.WeightNumber != 0 {
			css += fmt.Sprintf("  font-weight: %d;\n", font.WeightNumber)
		}

		if !font.UseNumberedWeight && font.Weight != "" {
			css += "  font-weight: " + font.Weight + ";\n"
		}

		// Add font-style if specified
		if font.Style != "" {
			css += "      font-style: " + font.Style + ";\n"
		}

		css += "    }"

		// Add newline after closing brace except for the last font
		if i < len(f.Fonts)-1 {
			css += "\n"
		}
	}

	css += "\n\t"
	css = css_util.Format(css)
	return css
}

type Font struct {
	src               []string
	localSrc          []string
	Weight            string
	WeightNumber      int
	UseNumberedWeight bool
	Style             string
}

func (f *Font) AddSrc(src string) {
	f.src = append(f.src, src)
}

func (f *Font) AddLocalSrc(fontName string) {
	f.localSrc = append(f.localSrc, fontName)
}

func (f *Font) SetWeightNumber(weight int) {
	f.WeightNumber = weight
	f.UseNumberedWeight = true
}

func (f *Font) SetWeight(weight string) {
	f.Weight = weight
}

func (f *Font) SetStyle(style string) {
	f.Style = style
}
