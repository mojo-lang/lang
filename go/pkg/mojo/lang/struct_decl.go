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

func (m *StructDecl) IsBoxedType() bool {
	if m != nil && m.Type != nil && len(m.Type.Inherits) == 1 && len(m.Type.Fields) == 0 {
		inherit := m.Type.Inherits[0]
		structDecl := inherit.TypeDeclaration.GetStructDecl()
		return structDecl != nil && structDecl.Type == nil
	}
	return false
}
