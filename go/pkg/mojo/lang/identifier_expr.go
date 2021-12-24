package lang

func NewTypeIdentifierExpr(nominal *NominalType) *IdentifierExpr {
	if nominal != nil {
		return &IdentifierExpr{
			StartPosition: nominal.StartPosition,
			EndPosition:   nominal.EndPosition,
			Kind:          0,
			Implicit:      false,
			Identifier: &Identifier{
				StartPosition:      nil,
				EndPosition:        nil,
				Kind:               Identifier_KIND_STRUCT,
				PackageName:        nominal.PackageName,
				Name:               nominal.Name,
				Declaration:        NewDeclarationFromTypeDeclaration(nominal.TypeDeclaration),
				FullName:           nominal.GetFullName(),
				EnclosingTypeNames: nominal.GetEnclosingNames(),
			},
			GenericArguments: nominal.GenericArguments,
		}
	}
	return nil
}


func (m *IdentifierExpr) GetName() string {
	if m != nil && m.Identifier != nil {
		return m.Identifier.Name
	}
	return ""
}

func (m *IdentifierExpr) GetFullName() string {
	if m != nil && m.Identifier != nil {
		return m.Identifier.FullName
	}
	return ""
}

func (m *IdentifierExpr) GetPackageName() string {
	if m != nil && m.Identifier != nil {
		return m.Identifier.PackageName
	}
	return ""
}
