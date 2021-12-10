package lang

func (m *StructDecl) GetFullName() string {
	if m != nil {
		return GetFullName(m.PackageName, GetEnclosingNames(m.EnclosingType), m.Name)
	}
	return ""
}

func (m *StructDecl) SetScope(scope *Scope) {
	if m != nil {
		m.Scope = scope
	}
}

func (m *StructDecl) IsBoxed() bool {
	if m != nil && m.Type != nil && len(m.Type.Inherits) == 1 && len(m.Type.Fields) == 0 {
		inherit := m.Type.Inherits[0]
		structDecl := inherit.TypeDeclaration.GetStructDecl()
		return structDecl != nil && structDecl.Type == nil
	}
	return false
}

func (m *StructDecl) IsOpacity() bool {
	return m != nil && m.Type == nil
}

func (m *StructDecl) IsGeneric() bool {
	return m != nil && len(m.GenericParameters) > 0
}

func (m *StructDecl) FieldCount() int {
	return len(m.GetType().GetFields())
}

func (m *StructDecl) FieldNames() []string {
	if m != nil {
		return m.Type.FieldNames()
	}
	return nil
}

func (m *StructDecl) EnclosingTypeDecl() interface{} {
	return m
}
