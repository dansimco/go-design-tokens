package color

type Mode struct {
	Name  string
	Roles []*Role
}

func NewMode(name string) Mode {
	mode := Mode{
		Name: name,
	}
	return mode
}

func (m *Mode) AddRole(name string) *Role {
	role := NewRole(name)
	m.Roles = append(m.Roles, role)
	return role
}

func (m *Mode) ToCSS() string {
	css := ""
	for _, role := range m.Roles {
		css += role.toCSS("  --" + m.Name)
	}
	return css
}
