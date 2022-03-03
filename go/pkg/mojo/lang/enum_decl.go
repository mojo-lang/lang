package lang

func (m *EnumDecl) GetFullName() string {
	if m != nil {
		return GetFullName(m.PackageName, GetEnclosingNames(m.EnclosingType), m.Name)
	}
	return ""
}

func (m *EnumDecl) SetScope(scope *Scope) {
	if m != nil {
		m.Scope = scope
	}
}

func (m *EnumDecl) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *EnumDecl) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}
