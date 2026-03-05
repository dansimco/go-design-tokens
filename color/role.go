package color

type Role struct {
	Name   string
	States []UIColor
}

func NewRole(name string) *Role {
	return &Role{
		Name:   name,
		States: []UIColor{},
	}
}

func (r *Role) AddState(state string, color UIColor) {
	color.Name = state
	r.States = append(r.States, color)
}

func (r *Role) toCSS(prepend string) string {
	css := ""
	if prepend == "" {
		prepend = "--"
	} else {
		prepend = prepend + "-"
	}

	for _, state := range r.States {
		if state.Name == "default" {
			css += prepend + r.Name + ": " + state.ToCSSLightDark() + ";\n"
		} else {
			css += prepend + r.Name + "-" + state.Name + ": " + state.ToCSSLightDark() + ";\n"
		}
	}

	return css
}
