package typography

import (
	"strings"
	"testing"

	"github.com/dansimco/go-design-tokens/css_util"
)

type typographyTestInputData struct {
	Family Family
	Styles []Style
}

func testSetup() typographyTestInputData {

	helvetica_now := NewFontFamily("Helvetica Now")

	helvetica_now_regular := (&helvetica_now).AddFont()
	helvetica_now_regular.AddLocalSrc("Helvetica Now")
	helvetica_now_regular.AddSrc("/assets/fonts/helvetica_now_regular.woff2")
	helvetica_now_regular.SetWeight("regular")

	helvetica_now_bold := (&helvetica_now).AddFont()
	helvetica_now_bold.AddLocalSrc("Helvetica Now")
	helvetica_now_bold.AddSrc("/assets/fonts/helvetica_now_bold")
	helvetica_now_bold.SetWeightNumber(600)

	(&helvetica_now).AddFallbackFont("arial")

	// styles
	//
	t_body := NewTypeStyle("body")
	t_body.SetFamily(&helvetica_now)
	t_body.SetSize(0.9)
	t_body.SetWeight("regular")
	t_body.SetLineHeight(1.5)

	t_caption := NewTypeStyle("caption")
	t_caption.SetFamily(&helvetica_now)
	t_caption.SetSize(0.875)
	t_caption.SetWeight("regular")
	t_caption.SetLineHeight(1.5)
	t_caption.SetTracking(0.03)

	return typographyTestInputData{
		Family: helvetica_now,
		Styles: []Style{t_body, t_caption},
	}
}

func TestCSSFontFaceGeneration(t *testing.T) {

	family := testSetup().Family

	generated_css := css_util.Format(family.ToCSS())

	expected_css := `
	@font-face {
	  font-family: "Helvetica Now";
	  src: local("Helvetica Now"), url("/assets/fonts/helvetica_now_regular.woff2") format("woff2");
	  font-weight: 400;
	  font-style: normal;
	  font-display: swap;
	}

	@font-face {
	  font-family: "Helvetica Now";
	  src: local("Helvetica Now"), url("/assets/fonts/helvetica_now_bold.woff2") format("woff2");
	  font-weight: 600;
	  font-style: normal;
	  font-display: swap;
	}
	`

	expected_css = css_util.Format(expected_css)

	if generated_css != expected_css {
		t.Errorf("Expected generated_css to equal \n %s \n got \n %s", expected_css, generated_css)
	}

}

func TestCSSStyleGeneration(t *testing.T) {

	expected_css := `
		.body {
		    font-family: "Helvetica Now", arial;
		    font-size: 0.9rem;
		    line-height: 1.5rem;
		    font-weight: 400;
		}

		.caption {
		    font-family: "Helvetica Now", arial;
		    font-size: 0.875rem;
		    line-height: 1.5rem;
			letter-spacing: 0.03rem;
		    font-weight: 400;
		}
	`
	expected_css = css_util.Format(expected_css)

	generated_css := ""
	styles := testSetup().Styles

	for _, style := range styles {
		generated_css += style.ToCSS()
	}

	generated_css = css_util.Format(generated_css)

	if generated_css != expected_css {
		t.Errorf("Expected generated_css to equal \n %s \n got \n %s", expected_css, generated_css)
	}

}

func TestCSSVarsGeneration(t *testing.T) {

	style := testSetup().Styles[1]
	vars := style.ToCSSVars("t")

	expected := `--t-caption-font-family: "Helvetica Now", arial;
--t-caption-font-size: 0.875rem;
--t-caption-line-height: 1.5rem;
--t-caption-letter-spacing: 0.03rem;
--t-caption-font-weight: 400;
`

	if vars != expected {
		t.Errorf("Expected vars to equal \n %s \n got \n %s", expected, vars)
	}
}

func TestResolve(t *testing.T) {

	style := testSetup().Styles[0]
	r := style.Resolve()

	if r.Name != "body" || r.ClassName != "body" {
		t.Errorf("expected name and class name 'body', got %q / %q", r.Name, r.ClassName)
	}
	if r.FamilyName != "Helvetica Now" {
		t.Errorf("expected family 'Helvetica Now', got %q", r.FamilyName)
	}
	if len(r.FallbackFonts) != 1 || r.FallbackFonts[0] != "arial" {
		t.Errorf("expected fallbacks [arial], got %v", r.FallbackFonts)
	}
	if r.SizeRem != 0.9 || r.LineHeightRem != 1.5 || r.TrackingRem != 0 {
		t.Errorf("unexpected metrics: size %g, line-height %g, tracking %g", r.SizeRem, r.LineHeightRem, r.TrackingRem)
	}
	if r.Weight != WeightRegular {
		t.Errorf("expected weight %d, got %d", WeightRegular, r.Weight)
	}
}

func TestWeightFromName(t *testing.T) {
	cases := map[string]int{
		"regular":    WeightRegular,
		"Regular":    WeightRegular,
		"semi-bold":  WeightSemiBold,
		"Semi Bold":  WeightSemiBold,
		"bold":       WeightBold,
		"black":      WeightBlack,
		"not-a-name": WeightRegular,
	}
	for name, want := range cases {
		if got := WeightFromName(name); got != want {
			t.Errorf("WeightFromName(%q) = %d, want %d", name, got, want)
		}
	}
}

func TestStyleNameSanitization(t *testing.T) {

	style := NewTypeStyle("Body Large!")
	if style.Name != "body-large" {
		t.Errorf("expected sanitized name 'body-large', got %q", style.Name)
	}

	css := style.ToCSS()
	if !strings.Contains(css, ".body-large {") {
		t.Errorf("expected selector '.body-large', got %s", css)
	}
}

func TestGenericFallbacksUnquoted(t *testing.T) {

	family := NewFontFamily("Helvetica Now")
	(&family).AddFallbackFont("sans-serif")

	style := NewTypeStyle("body")
	style.SetFamily(&family)

	css := style.ToCSS()
	if !strings.Contains(css, `font-family: "Helvetica Now", sans-serif;`) {
		t.Errorf("expected unquoted generic fallback, got %s", css)
	}
}

func TestTextCaseFromName(t *testing.T) {
	cases := map[string]string{
		"uppercase":  TextCaseUppercase,
		"Upper":      TextCaseUppercase,
		"All Caps":   TextCaseUppercase,
		"lower-case": TextCaseLowercase,
		"title":      TextCaseCapitalize,
		"none":       TextCaseOriginal,
		"not-a-name": TextCaseOriginal,
	}
	for name, want := range cases {
		if got := TextCaseFromName(name); got != want {
			t.Errorf("TextCaseFromName(%q) = %q, want %q", name, got, want)
		}
	}
}

func TestUppercaseStyleCSS(t *testing.T) {

	style := NewTypeStyle("overline")
	style.SetSize(0.75)
	style.SetUppercase()

	if style.Resolve().TextCase != TextCaseUppercase {
		t.Errorf("expected resolved text case %q, got %q", TextCaseUppercase, style.Resolve().TextCase)
	}

	css := style.ToCSS()
	if !strings.Contains(css, "text-transform: uppercase;") {
		t.Errorf("expected 'text-transform: uppercase', got %s", css)
	}

	vars := style.ToCSSVars("t")
	if !strings.Contains(vars, "--t-overline-text-transform: uppercase;") {
		t.Errorf("expected text-transform custom property, got %s", vars)
	}
}

func TestTextCaseOriginalOmittedFromCSS(t *testing.T) {

	style := NewTypeStyle("body")
	style.SetSize(1)
	style.SetTextCase("original")

	if css := style.ToCSS(); strings.Contains(css, "text-transform") {
		t.Errorf("expected no text-transform for original case, got %s", css)
	}
}

func TestFontFormatFromName(t *testing.T) {
	cases := map[string]string{
		"woff2":             FormatWOFF2,
		".woff":             FormatWOFF,
		"TTF":               FormatTrueType,
		"truetype":          FormatTrueType,
		"otf":               FormatOpenType,
		"ttc":               FormatCollection,
		"embedded-opentype": FormatEOT,
		"eot":               FormatEOT,
		"not-a-format":      "",
	}
	for name, want := range cases {
		if got := FormatFromName(name); got != want {
			t.Errorf("FormatFromName(%q) = %q, want %q", name, got, want)
		}
	}
}

func TestFontFaceFormatsFromExtension(t *testing.T) {

	family := NewFontFamily("Inter")
	font := (&family).AddFont()
	font.AddSrc("/assets/fonts/inter.woff2")
	font.AddSrc("/assets/fonts/inter.woff")
	font.AddSrc("/assets/fonts/inter.ttf")
	font.AddSrc("/assets/fonts/inter.otf")
	font.AddSrc("/assets/fonts/inter.eot?#iefix")

	css := family.ToCSS()

	expected := `src: url("/assets/fonts/inter.woff2") format("woff2"), ` +
		`url("/assets/fonts/inter.woff") format("woff"), ` +
		`url("/assets/fonts/inter.ttf") format("truetype"), ` +
		`url("/assets/fonts/inter.otf") format("opentype"), ` +
		`url("/assets/fonts/inter.eot?#iefix") format("embedded-opentype");`

	if !strings.Contains(css, expected) {
		t.Errorf("expected src \n %s \n got \n %s", expected, css)
	}
}

func TestFontFaceExplicitFormat(t *testing.T) {

	family := NewFontFamily("Inter")
	font := (&family).AddFont()
	font.AddSrcFormat("https://cdn.example.com/fonts?id=inter-regular", "woff2")

	css := family.ToCSS()
	if !strings.Contains(css, `url("https://cdn.example.com/fonts?id=inter-regular") format("woff2");`) {
		t.Errorf("expected explicit format on an extensionless URL, got %s", css)
	}
}

func TestFontFaceUnknownExtensionOmitsFormat(t *testing.T) {

	family := NewFontFamily("Inter")
	font := (&family).AddFont()
	font.AddSrc("/assets/fonts/inter.dfont")

	css := family.ToCSS()
	if !strings.Contains(css, `url("/assets/fonts/inter.dfont");`) {
		t.Errorf("expected src without a format hint, got %s", css)
	}
	if strings.Contains(css, "format(") {
		t.Errorf("expected no format() for an unknown extension, got %s", css)
	}
}

func TestFontFaceExtensionlessSrcDefaultsToWOFF2(t *testing.T) {

	family := NewFontFamily("Inter")
	font := (&family).AddFont()
	font.AddSrc("/assets/fonts/inter-regular")

	css := family.ToCSS()
	if !strings.Contains(css, `url("/assets/fonts/inter-regular.woff2") format("woff2");`) {
		t.Errorf("expected .woff2 to be appended, got %s", css)
	}
}
