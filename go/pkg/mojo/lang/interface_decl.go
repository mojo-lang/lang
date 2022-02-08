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

func (m *InterfaceDecl) GetMethodGroups() map[string][]*FunctionDecl {
	if m != nil && m.Type != nil {
		return m.Type.GetMethodGroups()
	}
	return nil
}
