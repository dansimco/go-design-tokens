package typography

import (
	"regexp"
	"strconv"
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
		Weight:  WeightRegular,
		Style:   "normal",
		Display: "swap",
	}
	f.Fonts = append(f.Fonts, font)
	return &f.Fonts[len(f.Fonts)-1]
}

func (f *Family) AddFallbackFont(font_name string) {
	f.FallbackFonts = append(f.FallbackFonts, font_name)
}

var cssIdent = regexp.MustCompile(`^-?[a-zA-Z_][a-zA-Z0-9_-]*$`)

// cssFontName quotes a font name unless it is a single CSS identifier.
// Generic families (sans-serif, monospace, ...) must stay unquoted or
// browsers treat them as literal font names.
func cssFontName(name string) string {
	if cssIdent.MatchString(name) {
		return name
	}
	return `"` + name + `"`
}

// fontFamilyList renders a primary family and its fallbacks as a CSS
// font-family value.
func fontFamilyList(primary string, fallbacks []string) string {
	names := []string{cssFontName(primary)}
	for _, fallback := range fallbacks {
		names = append(names, cssFontName(fallback))
	}
	return strings.Join(names, ", ")
}

func (f *Family) ToCSS() string {
	var b strings.Builder

	for _, font := range f.Fonts {
		b.WriteString("@font-face {\n")
		b.WriteString("font-family: " + cssFontName(f.Name) + ";\n")

		var srcEntries []string
		for _, local := range font.localSrc {
			srcEntries = append(srcEntries, `local("`+local+`")`)
		}
		for _, src := range font.src {
			// Replace any existing extension with .woff2
			if !strings.HasSuffix(src, ".woff2") {
				if idx := strings.LastIndex(src, "."); idx != -1 && idx > strings.LastIndex(src, "/") {
					src = src[:idx]
				}
				src += ".woff2"
			}
			srcEntries = append(srcEntries, `url("`+src+`") format("woff2")`)
		}
		if len(srcEntries) > 0 {
			b.WriteString("src: " + strings.Join(srcEntries, ", ") + ";\n")
		}

		if font.Weight != 0 {
			b.WriteString("font-weight: " + strconv.Itoa(font.Weight) + ";\n")
		}
		if font.Style != "" {
			b.WriteString("font-style: " + font.Style + ";\n")
		}
		if font.Display != "" {
			b.WriteString("font-display: " + font.Display + ";\n")
		}

		b.WriteString("}\n")
	}

	return css_util.Format(b.String())
}

type Font struct {
	src      []string
	localSrc []string
	Weight   int // 0 = unset; use the Weight* constants
	Style    string
	Display  string
}

func (f *Font) AddSrc(src string) {
	f.src = append(f.src, src)
}

func (f *Font) AddLocalSrc(fontName string) {
	f.localSrc = append(f.localSrc, fontName)
}

func (f *Font) SetWeightNumber(weight int) {
	f.Weight = weight
}

func (f *Font) SetWeight(weight string) {
	f.Weight = WeightFromName(weight)
}

func (f *Font) SetStyle(style string) {
	f.Style = style
}

func (f *Font) SetDisplay(display string) {
	f.Display = display
}
