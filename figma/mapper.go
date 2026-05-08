package figma

import (
	"fmt"

	"github.com/dansimco/go-design-tokens/theme"
)

func ThemeToVariables(t *theme.Theme) PublishVariablesBody {
	var body PublishVariablesBody

	baseUnit := float64(t.BaseSpacingUnit)
	if baseUnit == 0 {
		baseUnit = 16
	}
	gridDiv := t.GridDivision
	if gridDiv == 0 {
		gridDiv = 0.25
	}

	mapColors(&body, t)
	mapSpacing(&body, t, baseUnit, gridDiv)
	mapRadius(&body, t, baseUnit)
	mapTypography(&body, t, baseUnit)

	return body
}

func mapColors(body *PublishVariablesBody, t *theme.Theme) {
	for _, mode := range t.ColorModes {
		collID := "$collection:colors-" + mode.Name
		lightModeID := collID + "-light"
		darkModeID := collID + "-dark"

		body.VariableCollections = append(body.VariableCollections, VariableCollectionChange{
			Action:        "CREATE",
			ID:            collID,
			Name:          "colors/" + mode.Name,
			InitialModeID: lightModeID,
		})
		body.VariableModes = append(body.VariableModes,
			VariableModeChange{Action: "UPDATE", ID: lightModeID, Name: "Light"},
			VariableModeChange{Action: "CREATE", ID: darkModeID, Name: "Dark", VariableCollectionID: collID},
		)

		for _, role := range mode.Roles {
			for _, state := range role.States {
				varID := fmt.Sprintf("$var:colors-%s-%s-%s", mode.Name, role.Name, state.Name)
				body.Variables = append(body.Variables, VariableChange{
					Action:               "CREATE",
					ID:                   varID,
					Name:                 role.Name + "/" + state.Name,
					VariableCollectionID: collID,
					ResolvedType:         "COLOR",
				})
				body.VariableModeValues = append(body.VariableModeValues,
					VariableModeValue{VariableID: varID, ModeID: lightModeID, Value: labToRGBA(state.Light)},
					VariableModeValue{VariableID: varID, ModeID: darkModeID, Value: labToRGBA(state.Dark)},
				)
			}
		}
	}
}

func mapSpacing(body *PublishVariablesBody, t *theme.Theme, baseUnit, gridDiv float64) {
	collID := "$collection:spacing"
	modeID := "$mode:spacing-default"

	body.VariableCollections = append(body.VariableCollections, VariableCollectionChange{
		Action:        "CREATE",
		ID:            collID,
		Name:          "spacing",
		InitialModeID: modeID,
	})
	body.VariableModes = append(body.VariableModes, VariableModeChange{
		Action: "UPDATE",
		ID:     modeID,
		Name:   "Default",
	})

	for _, token := range t.SpaceTokens {
		varID := "$var:spacing-" + token.Name
		body.Variables = append(body.Variables, VariableChange{
			Action:               "CREATE",
			ID:                   varID,
			Name:                 "named/" + token.Name,
			VariableCollectionID: collID,
			ResolvedType:         "FLOAT",
		})
		body.VariableModeValues = append(body.VariableModeValues, VariableModeValue{
			VariableID: varID,
			ModeID:     modeID,
			Value:      token.UnitMultiple * baseUnit,
		})
	}

	// Numeric scale sp-0 through sp-32
	body.Variables = append(body.Variables, VariableChange{
		Action: "CREATE", ID: "$var:spacing-scale-0",
		Name: "scale/sp-0", VariableCollectionID: collID, ResolvedType: "FLOAT",
	})
	body.VariableModeValues = append(body.VariableModeValues, VariableModeValue{
		VariableID: "$var:spacing-scale-0", ModeID: modeID, Value: 0.0,
	})
	for i := 1; i <= 32; i++ {
		varID := fmt.Sprintf("$var:spacing-scale-%d", i)
		body.Variables = append(body.Variables, VariableChange{
			Action:               "CREATE",
			ID:                   varID,
			Name:                 fmt.Sprintf("scale/sp-%d", i),
			VariableCollectionID: collID,
			ResolvedType:         "FLOAT",
		})
		body.VariableModeValues = append(body.VariableModeValues, VariableModeValue{
			VariableID: varID,
			ModeID:     modeID,
			Value:      float64(i) * gridDiv * baseUnit,
		})
	}
}

func mapRadius(body *PublishVariablesBody, t *theme.Theme, baseUnit float64) {
	if len(t.RadiusTokens) == 0 {
		return
	}
	collID := "$collection:radius"
	modeID := "$mode:radius-default"

	body.VariableCollections = append(body.VariableCollections, VariableCollectionChange{
		Action:        "CREATE",
		ID:            collID,
		Name:          "radius",
		InitialModeID: modeID,
	})
	body.VariableModes = append(body.VariableModes, VariableModeChange{
		Action: "UPDATE",
		ID:     modeID,
		Name:   "Default",
	})

	for _, token := range t.RadiusTokens {
		varID := "$var:radius-" + token.Name
		body.Variables = append(body.Variables, VariableChange{
			Action:               "CREATE",
			ID:                   varID,
			Name:                 token.Name,
			VariableCollectionID: collID,
			ResolvedType:         "FLOAT",
		})
		body.VariableModeValues = append(body.VariableModeValues, VariableModeValue{
			VariableID: varID,
			ModeID:     modeID,
			Value:      token.UnitMultiple * baseUnit,
		})
	}
}

func mapTypography(body *PublishVariablesBody, t *theme.Theme, baseUnit float64) {
	if len(t.TypeStyles) == 0 {
		return
	}
	collID := "$collection:typography"
	modeID := "$mode:typography-default"

	body.VariableCollections = append(body.VariableCollections, VariableCollectionChange{
		Action:        "CREATE",
		ID:            collID,
		Name:          "typography",
		InitialModeID: modeID,
	})
	body.VariableModes = append(body.VariableModes, VariableModeChange{
		Action: "UPDATE",
		ID:     modeID,
		Name:   "Default",
	})

	for _, style := range t.TypeStyles {
		prefix := style.Name + "/"

		if style.Family != nil && style.Family.Name != "" {
			varID := "$var:type-" + style.Name + "-family"
			body.Variables = append(body.Variables, VariableChange{
				Action: "CREATE", ID: varID,
				Name: prefix + "font-family", VariableCollectionID: collID, ResolvedType: "STRING",
			})
			body.VariableModeValues = append(body.VariableModeValues, VariableModeValue{
				VariableID: varID, ModeID: modeID, Value: style.Family.Name,
			})
		}

		if style.Size > 0 {
			varID := "$var:type-" + style.Name + "-size"
			body.Variables = append(body.Variables, VariableChange{
				Action: "CREATE", ID: varID,
				Name: prefix + "font-size", VariableCollectionID: collID, ResolvedType: "FLOAT",
			})
			body.VariableModeValues = append(body.VariableModeValues, VariableModeValue{
				VariableID: varID, ModeID: modeID, Value: float64(style.Size) * baseUnit,
			})
		}

		if style.LineHeight > 0 {
			varID := "$var:type-" + style.Name + "-lh"
			body.Variables = append(body.Variables, VariableChange{
				Action: "CREATE", ID: varID,
				Name: prefix + "line-height", VariableCollectionID: collID, ResolvedType: "FLOAT",
			})
			body.VariableModeValues = append(body.VariableModeValues, VariableModeValue{
				VariableID: varID, ModeID: modeID, Value: float64(style.LineHeight) * baseUnit,
			})
		}

		if style.Tracking != 0 {
			varID := "$var:type-" + style.Name + "-tracking"
			body.Variables = append(body.Variables, VariableChange{
				Action: "CREATE", ID: varID,
				Name: prefix + "letter-spacing", VariableCollectionID: collID, ResolvedType: "FLOAT",
			})
			body.VariableModeValues = append(body.VariableModeValues, VariableModeValue{
				VariableID: varID, ModeID: modeID, Value: float64(style.Tracking) * baseUnit,
			})
		}

		if style.UseNumberedWeight && style.WeightNumber != 0 {
			varID := "$var:type-" + style.Name + "-weight"
			body.Variables = append(body.Variables, VariableChange{
				Action: "CREATE", ID: varID,
				Name: prefix + "font-weight", VariableCollectionID: collID, ResolvedType: "FLOAT",
			})
			body.VariableModeValues = append(body.VariableModeValues, VariableModeValue{
				VariableID: varID, ModeID: modeID, Value: float64(style.WeightNumber),
			})
		} else if !style.UseNumberedWeight && style.Weight != "" {
			varID := "$var:type-" + style.Name + "-weight"
			body.Variables = append(body.Variables, VariableChange{
				Action: "CREATE", ID: varID,
				Name: prefix + "font-weight", VariableCollectionID: collID, ResolvedType: "STRING",
			})
			body.VariableModeValues = append(body.VariableModeValues, VariableModeValue{
				VariableID: varID, ModeID: modeID, Value: style.Weight,
			})
		}
	}
}

