package figma

import (
	"encoding/json"
	"fmt"

	"github.com/dansimco/go-design-tokens/color"
	"github.com/dansimco/go-design-tokens/theme"
)

type pluginScriptData struct {
	ColorCollections []pluginColorColl `json:"colorCollections,omitempty"`
	SpacingVars      []pluginFloatVar  `json:"spacingVars,omitempty"`
	RadiusVars       []pluginFloatVar  `json:"radiusVars,omitempty"`
	TypoVars         []pluginTypoVar   `json:"typoVars,omitempty"`
}

type pluginColorColl struct {
	Name string           `json:"name"`
	Vars []pluginColorVar `json:"vars"`
}

type pluginColorVar struct {
	Name  string    `json:"name"`
	Light RGBAValue `json:"light"`
	Dark  RGBAValue `json:"dark"`
}

type pluginFloatVar struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

// pluginTypoVar holds a single typography variable; Value is float64 or string.
type pluginTypoVar struct {
	Name    string `json:"name"`
	VarType string `json:"type"`
	Value   any    `json:"value"`
}

// ThemeToPluginScript generates JavaScript for the Figma Plugin API that creates
// all variables when executed via the use_figma MCP tool.
func ThemeToPluginScript(t *theme.Theme) string {
	baseUnit := float64(t.BaseSpacingUnit)
	if baseUnit == 0 {
		baseUnit = 16
	}
	gridDiv := t.GridDivision
	if gridDiv == 0 {
		gridDiv = 0.25
	}

	data := pluginScriptData{}
	buildColorCollections(&data, t)
	buildSpacingVars(&data, t, baseUnit, gridDiv)
	buildRadiusVars(&data, t, baseUnit)
	buildTypoVars(&data, t, baseUnit)

	dataJSON, _ := json.MarshalIndent(data, "", "  ")

	return fmt.Sprintf(`const data = %s;

for (const coll of (data.colorCollections || [])) {
  const c = figma.variables.createVariableCollection(coll.name);
  c.renameMode(c.modes[0].modeId, "Light");
  const darkModeId = c.addMode("Dark");
  for (const v of coll.vars) {
    const variable = figma.variables.createVariable(v.name, c, "COLOR");
    variable.setValueForMode(c.modes[0].modeId, v.light);
    variable.setValueForMode(darkModeId, v.dark);
  }
}

if ((data.spacingVars || []).length > 0) {
  const spColl = figma.variables.createVariableCollection("spacing");
  spColl.renameMode(spColl.modes[0].modeId, "Default");
  for (const v of data.spacingVars) {
    const variable = figma.variables.createVariable(v.name, spColl, "FLOAT");
    variable.setValueForMode(spColl.modes[0].modeId, v.value);
  }
}

if ((data.radiusVars || []).length > 0) {
  const rColl = figma.variables.createVariableCollection("radius");
  rColl.renameMode(rColl.modes[0].modeId, "Default");
  for (const v of data.radiusVars) {
    const variable = figma.variables.createVariable(v.name, rColl, "FLOAT");
    variable.setValueForMode(rColl.modes[0].modeId, v.value);
  }
}

if ((data.typoVars || []).length > 0) {
  const tColl = figma.variables.createVariableCollection("typography");
  tColl.renameMode(tColl.modes[0].modeId, "Default");
  for (const v of data.typoVars) {
    const variable = figma.variables.createVariable(v.name, tColl, v.type);
    variable.setValueForMode(tColl.modes[0].modeId, v.value);
  }
}

console.log("Done");
`, string(dataJSON))
}

func buildColorCollections(data *pluginScriptData, t *theme.Theme) {
	for _, mode := range t.ColorModes {
		coll := pluginColorColl{Name: "colors/" + mode.Name}
		for _, role := range mode.Roles {
			for _, state := range role.States {
				coll.Vars = append(coll.Vars, pluginColorVar{
					Name:  role.Name + "/" + state.Name,
					Light: labToRGBA(state.Light),
					Dark:  labToRGBA(state.Dark),
				})
			}
		}
		data.ColorCollections = append(data.ColorCollections, coll)
	}
}

func buildSpacingVars(data *pluginScriptData, t *theme.Theme, baseUnit, gridDiv float64) {
	for _, token := range t.SpaceTokens {
		data.SpacingVars = append(data.SpacingVars, pluginFloatVar{
			Name:  "named/" + token.Name,
			Value: token.UnitMultiple * baseUnit,
		})
	}
	data.SpacingVars = append(data.SpacingVars, pluginFloatVar{Name: "scale/sp-0", Value: 0})
	for i := 1; i <= 32; i++ {
		data.SpacingVars = append(data.SpacingVars, pluginFloatVar{
			Name:  fmt.Sprintf("scale/sp-%d", i),
			Value: float64(i) * gridDiv * baseUnit,
		})
	}
}

func buildRadiusVars(data *pluginScriptData, t *theme.Theme, baseUnit float64) {
	for _, token := range t.RadiusTokens {
		data.RadiusVars = append(data.RadiusVars, pluginFloatVar{
			Name:  token.Name,
			Value: token.UnitMultiple * baseUnit,
		})
	}
}

func buildTypoVars(data *pluginScriptData, t *theme.Theme, baseUnit float64) {
	for _, style := range t.TypeStyles {
		prefix := style.Name + "/"

		if style.Family != nil && style.Family.Name != "" {
			data.TypoVars = append(data.TypoVars, pluginTypoVar{
				Name: prefix + "font-family", VarType: "STRING", Value: style.Family.Name,
			})
		}
		if style.Size > 0 {
			data.TypoVars = append(data.TypoVars, pluginTypoVar{
				Name: prefix + "font-size", VarType: "FLOAT", Value: float64(style.Size) * baseUnit,
			})
		}
		if style.LineHeight > 0 {
			data.TypoVars = append(data.TypoVars, pluginTypoVar{
				Name: prefix + "line-height", VarType: "FLOAT", Value: float64(style.LineHeight) * baseUnit,
			})
		}
		if style.Tracking != 0 {
			data.TypoVars = append(data.TypoVars, pluginTypoVar{
				Name: prefix + "letter-spacing", VarType: "FLOAT", Value: float64(style.Tracking) * baseUnit,
			})
		}
		if style.UseNumberedWeight && style.WeightNumber != 0 {
			data.TypoVars = append(data.TypoVars, pluginTypoVar{
				Name: prefix + "font-weight", VarType: "FLOAT", Value: float64(style.WeightNumber),
			})
		} else if !style.UseNumberedWeight && style.Weight != "" {
			data.TypoVars = append(data.TypoVars, pluginTypoVar{
				Name: prefix + "font-weight", VarType: "STRING", Value: style.Weight,
			})
		}
	}
}

// labToRGBA is also used by mapper.go — kept here as it's used by both script and REST approaches.
func labToRGBA(c color.LABColor) RGBAValue {
	rgb := c.ToRGB()
	return RGBAValue{R: rgb[0] / 255.0, G: rgb[1] / 255.0, B: rgb[2] / 255.0, A: 1.0}
}
