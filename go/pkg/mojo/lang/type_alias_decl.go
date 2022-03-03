package lang

const OriginalTypeAliasName = "original_type_alias_name"

func (m *TypeAliasDecl) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *TypeAliasDecl) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *TypeAliasDecl) GetFullName() string {
	if m != nil {
		return GetFullName(m.PackageName, GetEnclosingNames(m.EnclosingType), m.Name)
	}
	return ""
}

func (m *TypeAliasDecl) SetScope(scope *Scope) {
	if m != nil {
		m.Scope = scope
	}
}

func (m *TypeAliasDecl) IsGeneric() bool {
	return m != nil && len(m.GenericParameters) > 0
}
