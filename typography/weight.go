package typography

import "strings"

// Font weights on the standard CSS / OpenType numeric scale.
const (
	WeightThin       = 100
	WeightExtraLight = 200
	WeightLight      = 300
	WeightRegular    = 400
	WeightMedium     = 500
	WeightSemiBold   = 600
	WeightBold       = 700
	WeightExtraBold  = 800
	WeightBlack      = 900
)

var weightNames = map[string]int{
	"thin":       WeightThin,
	"hairline":   WeightThin,
	"extralight": WeightExtraLight,
	"ultralight": WeightExtraLight,
	"light":      WeightLight,
	"normal":     WeightRegular,
	"regular":    WeightRegular,
	"book":       WeightRegular,
	"medium":     WeightMedium,
	"semibold":   WeightSemiBold,
	"demibold":   WeightSemiBold,
	"bold":       WeightBold,
	"extrabold":  WeightExtraBold,
	"ultrabold":  WeightExtraBold,
	"black":      WeightBlack,
	"heavy":      WeightBlack,
}

// WeightFromName maps a common weight name ("regular", "semi-bold", ...) to
// its numeric value. Names are matched case-insensitively, ignoring spaces
// and hyphens. Unrecognised names fall back to WeightRegular so the output
// is always valid CSS.
func WeightFromName(name string) int {
	key := strings.ToLower(name)
	key = strings.ReplaceAll(key, "-", "")
	key = strings.ReplaceAll(key, " ", "")
	if w, ok := weightNames[key]; ok {
		return w
	}
	return WeightRegular
}
