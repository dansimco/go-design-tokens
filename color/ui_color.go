package color

type UIColor struct {
	Name  string
	Light LABColor
	Dark  LABColor
}

func NewUIColor(light LABColor, dark LABColor) UIColor {
	return UIColor{
		Light: light,
		Dark:  dark,
	}
}

func (c *UIColor) ToCSSLightDark() string {
	return "light-dark(" + c.Light.ToHex() + ", " + c.Dark.ToHex() + ")"
}
