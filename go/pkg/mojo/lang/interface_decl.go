package lang

func (m *InterfaceDecl) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *InterfaceDecl) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

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
