package typography

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
	return &font
}

func (f *Family) ToCSS() string {
	css := ""

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
