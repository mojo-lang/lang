package lang

func (m *InterfaceDecl) GetFullName() string {
	if m != nil {
		return GetFullName(m.PackageName, nil, m.Name)
	}
	return ""
}

func (m *InterfaceDecl) SetScope(scope *Scope) {
	if m != nil {
		m.Scope = scope
	}
}