package figma

import (
	"testing"

	"github.com/dansimco/go-design-tokens/color"
	"github.com/dansimco/go-design-tokens/theme"
)

func TestThemeToVariables_Colors(t *testing.T) {
	th := theme.New()
	mode := th.AddColorMode("neutral")
	content := mode.AddRole("content")
	content.AddState("default", color.UIColor{
		Light: color.FromHex("#FFFFFF"),
		Dark:  color.FromHex("#000000"),
	})
	content.AddState("hover", color.UIColor{
		Light: color.FromHex("#EEEEEE"),
		Dark:  color.FromHex("#111111"),
	})

	body := ThemeToVariables(&th)

	// 1 color collection + 1 spacing collection
	if len(body.VariableCollections) != 2 {
		t.Errorf("expected 2 collections, got %d", len(body.VariableCollections))
	}

	// Find the color collection
	var colorCollFound bool
	for _, c := range body.VariableCollections {
		if c.Name == "colors/neutral" && c.Action == "CREATE" {
			colorCollFound = true
		}
	}
	if !colorCollFound {
		t.Error("colors/neutral collection not found")
	}

	// color modes: 1 UPDATE (light) + 1 CREATE (dark) = 2; spacing: 1 UPDATE = 1; total 3
	if len(body.VariableModes) != 3 {
		t.Errorf("expected 3 mode changes, got %d", len(body.VariableModes))
	}

	// 2 color vars + 33 spacing scale vars (sp-0..sp-32)
	if len(body.Variables) != 35 {
		t.Errorf("expected 35 variables, got %d", len(body.Variables))
	}

	// COLOR variable names use role/state grouping
	var colorVars []VariableChange
	for _, v := range body.Variables {
		if v.ResolvedType == "COLOR" {
			colorVars = append(colorVars, v)
		}
	}
	if len(colorVars) != 2 {
		t.Fatalf("expected 2 COLOR variables, got %d", len(colorVars))
	}
	if colorVars[0].Name != "content/default" {
		t.Errorf("expected content/default, got %s", colorVars[0].Name)
	}
	if colorVars[1].Name != "content/hover" {
		t.Errorf("expected content/hover, got %s", colorVars[1].Name)
	}

	// 2 color vars × 2 modes + 33 spacing vars × 1 mode = 37 mode values
	if len(body.VariableModeValues) != 37 {
		t.Errorf("expected 37 mode values, got %d", len(body.VariableModeValues))
	}

	// White (#FFFFFF) light value should be RGBA {1, 1, 1, 1}
	defaultVarID := colorVars[0].ID
	lightModeID := "$collection:colors-neutral-light"
	for _, mv := range body.VariableModeValues {
		if mv.VariableID == defaultVarID && mv.ModeID == lightModeID {
			rgba, ok := mv.Value.(RGBAValue)
			if !ok {
				t.Fatalf("expected RGBAValue, got %T", mv.Value)
			}
			if rgba.R < 0.99 || rgba.G < 0.99 || rgba.B < 0.99 || rgba.A != 1.0 {
				t.Errorf("expected white RGBA, got %+v", rgba)
			}
			break
		}
	}
}

func TestThemeToVariables_Spacing(t *testing.T) {
	th := theme.New()
	th.AddSpaceToken("xs", 0.5)
	th.AddSpaceToken("sm", 1.0)

	body := ThemeToVariables(&th)

	// Only spacing collection (no colors/radius/typography)
	if len(body.VariableCollections) != 1 {
		t.Errorf("expected 1 collection, got %d", len(body.VariableCollections))
	}

	// 2 named + 33 scale
	if len(body.Variables) != 35 {
		t.Errorf("expected 35 variables, got %d", len(body.Variables))
	}

	// Named token values: xs = 0.5 * 16 = 8, sm = 1.0 * 16 = 16
	modeID := "$mode:spacing-default"
	xsVarID := "$var:spacing-xs"
	for _, mv := range body.VariableModeValues {
		if mv.VariableID == xsVarID && mv.ModeID == modeID {
			val, ok := mv.Value.(float64)
			if !ok {
				t.Fatalf("expected float64, got %T", mv.Value)
			}
			if val != 8.0 {
				t.Errorf("expected xs = 8px, got %v", val)
			}
		}
	}

	// Scale: sp-4 = 4 * 0.25 * 16 = 16px
	sp4VarID := "$var:spacing-scale-4"
	for _, mv := range body.VariableModeValues {
		if mv.VariableID == sp4VarID && mv.ModeID == modeID {
			val, ok := mv.Value.(float64)
			if !ok {
				t.Fatalf("expected float64, got %T", mv.Value)
			}
			if val != 16.0 {
				t.Errorf("expected sp-4 = 16px, got %v", val)
			}
		}
	}
}

func TestThemeToVariables_RadiusAndTypographySkippedWhenEmpty(t *testing.T) {
	th := theme.New()
	body := ThemeToVariables(&th)

	for _, c := range body.VariableCollections {
		if c.Name == "radius" {
			t.Error("radius collection should not be created when there are no radius tokens")
		}
		if c.Name == "typography" {
			t.Error("typography collection should not be created when there are no type styles")
		}
	}
}
