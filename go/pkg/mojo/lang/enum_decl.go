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