package theme

import (
	"strconv"

	"github.com/dansimco/go-design-tokens/color"
	"github.com/dansimco/go-design-tokens/radius"
	"github.com/dansimco/go-design-tokens/spacing"
	"github.com/dansimco/go-design-tokens/typography"
)

type Theme struct {
	BaseSpacingUnit int
	GridDivision    float64
	SpacePrefix     string
	SpaceTokens     []spacing.NamedSpace
	RadiusPrefix    string
	RadiusTokens    []radius.RadiusToken
	ColorPrefix     string
	ColorModes      []color.Mode
	TypePrefix      string
	TypeFamilies    []typography.Family
	TypeStyles      []typography.Style
}

func New() Theme {
	t := Theme{
		SpacePrefix:     "sp-",
		ColorPrefix:     "c-",
		RadiusPrefix:    "radius-",
		TypePrefix:      "t-",
		BaseSpacingUnit: 16,
		GridDivision:    0.25,
	}
	return t
}

func (t *Theme) AddColorMode(name string) *color.Mode {
	mode := color.Mode{Name: name}
	t.ColorModes = append(t.ColorModes, mode)
	return &t.ColorModes[len(t.ColorModes)-1]
}

func (t *Theme) AddRadiusToken(name string, unitMultiple float64) {
	t.RadiusTokens = append(t.RadiusTokens, radius.RadiusToken{Name: name, UnitMultiple: unitMultiple})
}

func (t *Theme) AddSpaceToken(name string, unitMultiple float64) {
	t.SpaceTokens = append(t.SpaceTokens, spacing.NamedSpace{Name: name, UnitMultiple: unitMultiple})
}

func (t *Theme) AddTypeFamily(name string) *typography.Family {
	family := typography.Family{
		Name: name,
	}
	t.TypeFamilies = append(t.TypeFamilies, family)
	return &t.TypeFamilies[len(t.TypeFamilies)-1]
}

func (t *Theme) AddTypeStyle(name string) *typography.Style {
	style := typography.NewTypeStyle(name)
	t.TypeStyles = append(t.TypeStyles, style)
	return &t.TypeStyles[len(t.TypeStyles)-1]
}

func (t *Theme) ToCSS() string {
	css := ""

	for _, family := range t.TypeFamilies {
		css += family.ToCSS()
	}

	for _, style := range t.TypeStyles {
		css += style.ToCSS()
	}

	css += ":root {\n"

	// type style tokens
	typePrefix := ""
	if t.TypePrefix != "" {
		typePrefix = t.TypePrefix[:len(t.TypePrefix)-1]
	}
	for _, style := range t.TypeStyles {
		styleVars := style.ToCSSVars(typePrefix)
		if styleVars != "" {
			// Indent each line
			for i := 0; i < len(styleVars); i++ {
				if i == 0 || styleVars[i-1] == '\n' {
					css += "  "
				}
				css += string(styleVars[i])
			}
		}
	}

	// radius tokens
	for _, radiusToken := range t.RadiusTokens {
		css += `  --` + t.RadiusPrefix + radiusToken.Name + `: ` + strconv.FormatFloat(radiusToken.UnitMultiple, 'f', -1, 64) + `rem;` + "\n"
	}
	// spacing tokens
	for _, spaceToken := range t.SpaceTokens {
		css += `  --` + t.SpacePrefix + spaceToken.Name + `: ` + strconv.FormatFloat(spaceToken.UnitMultiple, 'f', -1, 64) + `rem;` + "\n"
	}
	if t.GridDivision == 0 {
		t.GridDivision = 0.25
	}
	css += `  --` + t.SpacePrefix + `0: 0;` + "\n"
	for i := 1; i <= 32; i++ {
		css += `  --` + t.SpacePrefix + strconv.Itoa(i) + `: ` + strconv.FormatFloat(float64(i)*t.GridDivision, 'f', -1, 64) + `rem;` + "\n"
	}
	// color tokens
	for _, mode := range t.ColorModes {
		css += mode.ToCSS(t.ColorPrefix[:len(t.ColorPrefix)-1])
	}
	css += "}\n\n"

	// color mode classes
	for _, mode := range t.ColorModes {
		css += mode.ToCSSClass(t.ColorPrefix[:len(t.ColorPrefix)-1]) + "\n"
	}

	return css
}

func (t *Theme) GetSpacingValues(n int) []float64 {
	if t.BaseSpacingUnit == 0 {
		t.BaseSpacingUnit = 16
	}
	values := []float64{}

	values = append(values, float64(t.BaseSpacingUnit/4))
	values = append(values, float64(t.BaseSpacingUnit/2))

	for i := 1; i <= n; i++ {
		values = append(values, float64(i*t.BaseSpacingUnit))
	}

	return values
}

func (t *Theme) GenerateHTMLPreview() string {
	html := `<html lang="en">
<head>
  <style>
`
	// Add CSS variables
	html += t.ToCSS()
	html += `
   body {
	 font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol";
     margin: 2rem;
     color-scheme: light dark;
     background-color: var(--neutral-surface-bg);
     color: light-dark(#222, #aaa);
   }

   main {
     max-width: 30rem;
   	 margin: 0 auto;
   }
   .color-role-swatches {
     display: flex;
     gap: 0.25rem;
     margin-bottom: 2rem;
   }
   .color-role-swatches .swatch {
     flex-grow: 2;
     aspect-ratio: 1;
     font-size: 0.7rem;
   }
   h2 {
     margin-top: 2rem;
     border-bottom: 2px solid #333;
     padding-bottom: 0.5rem;
   }
   h3 {
     margin-top: 1rem;
     margin-bottom: 0.5rem;
     font-size: 1rem;
   }

   strong {
   font-weight: 500;
   }

   .mode-preview {
   	font-size: 0.7rem;
    overflow: hidden;
    border-radius: 0.5rem;

    & .bg-container {
    	padding: 0.5rem;
    }

    & p {
    margin: 0;
    padding: 0 0 0.25rem 0;
    }
   }
  </style>
</head>
<body>
<main>
<section id="colors">
`
	// Generate color swatches for each mode
	for _, mode := range t.ColorModes {
		html += "<h2>" + mode.Name + "</h2>\n"
		html += `<div class="mode-preview" style="border: 0.5px solid var(--` + mode.Name + `-` + `trim); border-bottom: none;">`
		for _, role := range mode.Roles {
			if role.Context == "background" {
				html += `<div class="bg-container" style="background-color: var(--` + mode.Name + `-` + role.Name + `); border-bottom: 0.5px solid var(--` + mode.Name + `-` + `trim);">`

				for _, fg_role := range mode.Roles {
					if fg_role.Context == "foreground" {
						html += `<p style="color: var(--` + mode.Name + `-` + fg_role.Name + `);"><strong>` + fg_role.Name + `</strong> on ` + role.Name + `</p>`
					}
				}
				html += `</div>`
			}
		}
		html += `</div>`
		for _, role := range mode.Roles {
			html += "<h5>" + role.Name + "</h4>\n"
			html += `<div class="color-role-swatches">` + "\n"

			// Loop through all states in order
			for _, state := range role.States {
				varName := "--" + mode.Name + "-" + role.Name
				title := role.Name
				if state.Name != "default" {
					varName += "-" + state.Name
					title += " (" + state.Name + ")"
				}
				html += `  <div class="swatch" title="` + title + `" style="background-color: var(` + varName + `);"> ` + state.Name + `</div>` + "\n"
			}

			html += "</div>\n"
		}
	}

	html += `</section>
</main>
</body>
</html>`

	return html
}
