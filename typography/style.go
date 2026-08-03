package typography

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dansimco/go-design-tokens/css_util"
)

type Style struct {
	Name       string
	Family     *Family
	Size       float64
	LineHeight float64
	Tracking   float64
	Weight     int // 0 = unset; use the Weight* constants
	Style      string
	cssClass   string
}

// Resolved is a platform-neutral snapshot of a Style with every value
// reduced to plain data, so renderers (CSS, Figma, iOS, Android) can
// format tokens without re-deriving any logic. Rem values are relative
// to the theme's base unit (1rem = BaseSpacingUnit px).
type Resolved struct {
	Name          string
	ClassName     string
	FamilyName    string
	FallbackFonts []string
	SizeRem       float64
	LineHeightRem float64
	TrackingRem   float64
	Weight        int // 0 = unset
	FontStyle     string
}

var invalidNameChars = regexp.MustCompile(`[^a-z0-9_-]+`)

// sanitizeName makes a style name safe for use as a CSS class name,
// custom property name, and cross-platform token identifier.
func sanitizeName(name string) string {
	name = strings.ToLower(strings.TrimSpace(name))
	name = invalidNameChars.ReplaceAllString(name, "-")
	return strings.Trim(name, "-")
}

func NewTypeStyle(name string) Style {
	return Style{
		Name: sanitizeName(name),
	}
}

func (s *Style) SetCSSClass(className string) {
	s.cssClass = sanitizeName(className)
}

func (s *Style) SetFamily(f *Family) {
	s.Family = f
}

func (s *Style) SetSize(size float64) {
	s.Size = size
}

func (s *Style) SetLineHeight(lineHeight float64) {
	s.LineHeight = lineHeight
}

func (s *Style) SetTracking(tracking float64) {
	s.Tracking = tracking
}

func (s *Style) SetWeight(weight string) {
	s.Weight = WeightFromName(weight)
}

func (s *Style) SetWeightNumber(weight int) {
	s.Weight = weight
}

func (s *Style) SetStyle(style string) {
	s.Style = style
}

func (s *Style) Resolve() Resolved {
	r := Resolved{
		Name:          s.Name,
		ClassName:     s.Name,
		SizeRem:       s.Size,
		LineHeightRem: s.LineHeight,
		TrackingRem:   s.Tracking,
		Weight:        s.Weight,
		FontStyle:     s.Style,
	}
	if s.cssClass != "" {
		r.ClassName = s.cssClass
	}
	if s.Family != nil {
		r.FamilyName = s.Family.Name
		r.FallbackFonts = s.Family.FallbackFonts
	}
	return r
}

// cssProperties returns the style's set properties as ordered
// name/value pairs, shared by the class and custom-property renderers.
func (r Resolved) cssProperties() [][2]string {
	var props [][2]string
	if r.FamilyName != "" {
		props = append(props, [2]string{"font-family", fontFamilyList(r.FamilyName, r.FallbackFonts)})
	}
	if r.SizeRem > 0 {
		props = append(props, [2]string{"font-size", fmt.Sprintf("%grem", r.SizeRem)})
	}
	if r.LineHeightRem > 0 {
		props = append(props, [2]string{"line-height", fmt.Sprintf("%grem", r.LineHeightRem)})
	}
	if r.TrackingRem != 0 {
		props = append(props, [2]string{"letter-spacing", fmt.Sprintf("%grem", r.TrackingRem)})
	}
	if r.Weight != 0 {
		props = append(props, [2]string{"font-weight", strconv.Itoa(r.Weight)})
	}
	if r.FontStyle != "" {
		props = append(props, [2]string{"font-style", r.FontStyle})
	}
	return props
}

func (s *Style) ToCSS() string {
	r := s.Resolve()

	var b strings.Builder
	b.WriteString("." + r.ClassName + " {\n")
	for _, p := range r.cssProperties() {
		b.WriteString(p[0] + ": " + p[1] + ";\n")
	}
	b.WriteString("}")

	return css_util.Format(b.String())
}

func (s *Style) ToCSSVars(prefix string) string {
	r := s.Resolve()

	varPrefix := "--"
	if prefix != "" {
		varPrefix += prefix + "-"
	}
	varPrefix += r.Name

	var b strings.Builder
	for _, p := range r.cssProperties() {
		b.WriteString(varPrefix + "-" + p[0] + ": " + p[1] + ";\n")
	}
	return b.String()
}
