package typography

type Font struct {
	Name         string
	Family       []string
	WeightNumber int
	Weight       string
	Style        string
}

func NewFont(name string) *Font {
	f := Font{
		Name: name,
	}
	return &f
}

func (f *Font) AddFamilyFont(fontName string) {
	f.Family = append(f.Family, fontName)
}
