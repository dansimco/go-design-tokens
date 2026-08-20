package typography

import "strings"

// Font file formats, named after their CSS @font-face format() keywords.
const (
	FormatWOFF2    = "woff2"
	FormatWOFF     = "woff"
	FormatTrueType = "truetype"
	FormatOpenType = "opentype"
	FormatEOT      = "embedded-opentype"
	FormatSVG      = "svg"
	// FormatCollection covers .ttc/.otc bundles of several faces.
	FormatCollection = "collection"
)

// formatExtensions maps a file extension to its format() keyword.
var formatExtensions = map[string]string{
	"woff2": FormatWOFF2,
	"woff":  FormatWOFF,
	"ttf":   FormatTrueType,
	"otf":   FormatOpenType,
	"ttc":   FormatCollection,
	"otc":   FormatCollection,
	"eot":   FormatEOT,
	"svg":   FormatSVG,
	"svgz":  FormatSVG,
}

var formatNames = map[string]string{
	"woff2":            FormatWOFF2,
	"woff":             FormatWOFF,
	"ttf":              FormatTrueType,
	"truetype":         FormatTrueType,
	"otf":              FormatOpenType,
	"opentype":         FormatOpenType,
	"ttc":              FormatCollection,
	"otc":              FormatCollection,
	"collection":       FormatCollection,
	"eot":              FormatEOT,
	"embeddedopentype": FormatEOT,
	"svg":              FormatSVG,
	"svgz":             FormatSVG,
}

// FormatFromName maps a font format name or file extension ("woff2", "ttf",
// "truetype", ...) to its CSS format() keyword. Names are matched
// case-insensitively, ignoring leading dots, spaces and hyphens.
// Unrecognised names return "" so the src is written without a format hint
// rather than with a wrong one.
func FormatFromName(name string) string {
	key := strings.ToLower(strings.TrimPrefix(strings.TrimSpace(name), "."))
	key = strings.ReplaceAll(key, "-", "")
	key = strings.ReplaceAll(key, " ", "")
	return formatNames[key]
}

// srcExtension returns the lowercased file extension of a font URL, ignoring
// any query string or fragment (".../font.woff2?v=2#iefix" -> "woff2").
// It returns "" when the URL's last path segment has no extension.
func srcExtension(src string) string {
	path := src
	if i := strings.IndexAny(path, "?#"); i != -1 {
		path = path[:i]
	}
	slash := strings.LastIndex(path, "/")
	dot := strings.LastIndex(path, ".")
	if dot == -1 || dot < slash {
		return ""
	}
	return strings.ToLower(path[dot+1:])
}
