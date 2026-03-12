package typography

import (
	"fmt"

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
		css += "\t  src:\n"

		// Add local sources
		for _, local := range font.localSrc {
			css += "        local(\"" + local + "\"),\n"
		}

		// Add URL sources
		for j, src := range font.src {
			// Add .woff2 extension if not present
			if len(src) < 6 || src[len(src)-6:] != ".woff2" {
				src += ".woff2"
			}
			css += "        url(\"" + src + "\")"
			if j < len(font.src)-1 {
				css += ","
			} else {
				css += ";"
			}
			css += "\n"
		}

		if font.UseNumberedWeight && font.WeightNumber != 0 {
			css += `  font-weight: ` + string(font.WeightNumber) + ";\n"
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
}

func (f *Font) SetWeight(weight string) {
	f.Weight = weight
}

func (f *Font) SetStyle(style string) {
	f.Style = style
}

type Style struct {
	Name              string
	Family            *Family
	Size              float32
	LineHeight        float32
	Tracking          float32
	Weight            string
	WeightNumber      int
	Style             string
	cssClass          string
	UseNumberedWeight bool
}

func NewTypeStyle(name string) Style {
	s := Style{
		Name: name,
	}
	return s
}

func (s *Style) SetCSSClass(className string) {
	s.cssClass = className
}

func (s *Style) SetFamily(f *Family) {
	s.Family = f
}

func (s *Style) SetSize(size float32) {
	s.Size = size
}

func (s *Style) SetLineHeight(lineHeight float32) {
	s.LineHeight = lineHeight
}

func (s *Style) SetTracking(tracking float32) {
	s.Tracking = tracking
}

func (s *Style) SetWeight(weight string) {
	s.Weight = weight
	s.UseNumberedWeight = false
}

func (s *Style) SetWeightNumber(weight int) {
	s.WeightNumber = weight
	s.UseNumberedWeight = true
}

func (s *Style) SetStyle(style string) {
	s.Style = style
}

func (s *Style) ToCSS() string {

	className := s.Name
	if s.cssClass != "" {
		className = s.cssClass
	}

	css := "." + className + " {\n"

	// Add font-family if specified
	if s.Family != nil && s.Family.Name != "" {
		css += "  font-family: \"" + s.Family.Name + "\""
		for _, fallback := range s.Family.FallbackFonts {
			css += ", \"" + fallback + "\""
		}
		css += ";\n"
	}

	// Add font-size if specified
	if s.Size > 0 {
		css += fmt.Sprintf("  font-size: %grem;\n", s.Size)
	}

	// Add line-height if specified
	if s.LineHeight > 0 {
		css += fmt.Sprintf("  line-height: %grem;\n", s.LineHeight)
	}

	// Add letter-spacing (tracking) if specified
	if s.Tracking != 0 {
		css += fmt.Sprintf("  letter-spacing: %grem;\n", s.Tracking)
	}

	// Add font-weight if specified

	if s.UseNumberedWeight && s.WeightNumber != 0 {
		css += `  font-weight: ` + string(s.WeightNumber) + ";\n"
	}

	if !s.UseNumberedWeight && s.Weight != "" {
		css += `  font-weight: ` + s.Weight + ";\n"
	}

	// Add font-style if specified
	if s.Style != "" {
		css += "  font-style: " + s.Style + ";\n"
	}

	css += "}"

	css = css_util.Format(css)
	return css
}
