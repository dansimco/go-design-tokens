package color

type Ramp struct {
	Keys []RampKey
}

type RampKey struct {
	Color    LABColor
	Position float64
}

func NewRamp() Ramp {
	return Ramp{}
}

func (ramp *Ramp) AddKey(hex string, position float64) {
	key := RampKey{}
	key.Color = FromHex(hex)
	if position < 0 {
		key.Position = 0
	} else if position > 1 {
		key.Position = 1
	} else {
		key.Position = position
	}
	ramp.Keys = append(ramp.Keys, key)
}

func (ramp *Ramp) At(position float64) LABColor {
	// Handle edge cases
	if len(ramp.Keys) == 0 {
		return LABColor{}
	}
	if len(ramp.Keys) == 1 {
		return ramp.Keys[0].Color
	}

	// Clamp position between 0 and 1
	if position <= 0 {
		position = 0
	}
	if position >= 1 {
		position = 1
	}

	// Find the two keys to interpolate between
	var lowerKey, upperKey RampKey
	foundLower := false
	foundUpper := false

	for i := 0; i < len(ramp.Keys); i++ {
		if ramp.Keys[i].Position <= position {
			if !foundLower || ramp.Keys[i].Position > lowerKey.Position {
				lowerKey = ramp.Keys[i]
				foundLower = true
			}
		}
		if ramp.Keys[i].Position >= position {
			if !foundUpper || ramp.Keys[i].Position < upperKey.Position {
				upperKey = ramp.Keys[i]
				foundUpper = true
			}
		}
	}

	// If position exactly matches a key, return that color
	if lowerKey.Position == position {
		return lowerKey.Color
	}

	// Linear interpolation between the two keys in LAB space
	t := (position - lowerKey.Position) / (upperKey.Position - lowerKey.Position)

	return LABColor{
		L: lowerKey.Color.L + (upperKey.Color.L-lowerKey.Color.L)*t,
		A: lowerKey.Color.A + (upperKey.Color.A-lowerKey.Color.A)*t,
		B: lowerKey.Color.B + (upperKey.Color.B-lowerKey.Color.B)*t,
	}
}


type swatch struct {
	Color LABColor
	Position float64
}

func (r* Ramp) CreateSwatches(count int) []swatch {
	colors := []swatch{}

	for i := 0; i < count; i++ {
		position := float64(i) / float64(count-1)
		colors = append(colors, swatch{
			Color:    r.At(position),
			Position: position,
		})
	}
	return colors
}
