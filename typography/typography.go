package typography

import "go-ds/css_util"

type Family struct {
	Name  string
	Fonts []Font
}

func NewFontFamily(name string) Family {
	return Family{
		Name: name,
	}
}

func (f *Family) AddFont() *Font {
	font := Font{
		Weight:       "normal",
		Style:        "regular",
		WeightNumber: 400,
	}
	f.Fonts = append(f.Fonts, font)
	return &f.Fonts[len(f.Fonts)-1]
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

		css += "      weight: "
		if font.Weight != "" && font.Weight != "normal" {
			// Use the Weight string field with quotes
			css += "\"" + font.Weight + "\""
		} else if font.WeightNumber != 0 {
			// Convert int to string properly
			weightStr := ""
			num := font.WeightNumber
			for num > 0 {
				weightStr = string(rune(num%10+'0')) + weightStr
				num /= 10
			}
			css += weightStr
		} else {
			css += "400"
		}
		css += ";\n"
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
	src          []string
	localSrc     []string
	Weight       string
	WeightNumber int
	Style        string
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

type Style struct {
	Name       string
	Family     *Family
	Size       float32
	LineHeight float32
	Tracking   float32
	Weight     string
	Style      string
}
