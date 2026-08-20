package typography

import "strings"

// Text cases, named after their CSS text-transform keywords so a Style's
// value can be rendered to CSS directly. Other renderers (Figma, iOS,
// Android) map these to their own platform enums.
const (
	TextCaseOriginal   = "none"
	TextCaseUppercase  = "uppercase"
	TextCaseLowercase  = "lowercase"
	TextCaseCapitalize = "capitalize"
)

var textCaseNames = map[string]string{
	"none":       TextCaseOriginal,
	"original":   TextCaseOriginal,
	"normal":     TextCaseOriginal,
	"uppercase":  TextCaseUppercase,
	"upper":      TextCaseUppercase,
	"caps":       TextCaseUppercase,
	"allcaps":    TextCaseUppercase,
	"lowercase":  TextCaseLowercase,
	"lower":      TextCaseLowercase,
	"capitalize": TextCaseCapitalize,
	"title":      TextCaseCapitalize,
	"titlecase":  TextCaseCapitalize,
}

// TextCaseFromName maps a common text case name ("uppercase", "all caps",
// "title", ...) to its CSS text-transform keyword. Names are matched
// case-insensitively, ignoring spaces and hyphens. Unrecognised names fall
// back to TextCaseOriginal so the output is always valid CSS.
func TextCaseFromName(name string) string {
	key := strings.ToLower(name)
	key = strings.ReplaceAll(key, "-", "")
	key = strings.ReplaceAll(key, " ", "")
	if c, ok := textCaseNames[key]; ok {
		return c
	}
	return TextCaseOriginal
}
