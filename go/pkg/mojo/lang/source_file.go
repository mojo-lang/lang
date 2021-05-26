package lang

func (m *SourceFile) SetScope(scope *Scope) {
	if m != nil {
		m.Scope = scope
	}
}


