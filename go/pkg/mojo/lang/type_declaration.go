package lang

func NewStructTypeDeclaration(decl *StructDecl) *TypeDeclaration {
	return &TypeDeclaration{
		TypeDeclaration: &TypeDeclaration_StructDecl{
			StructDecl: decl,
		},
	}
}

func NewTypeDeclarationFromDeclaration(decl *Declaration) *TypeDeclaration {
	switch decl.Declaration.(type) {
	case *Declaration_StructDecl:
		return &TypeDeclaration{
			TypeDeclaration: &TypeDeclaration_StructDecl{
				StructDecl: decl.GetStructDecl(),
			},
		}
	case *Declaration_InterfaceDecl:
		return &TypeDeclaration{
			TypeDeclaration: &TypeDeclaration_InterfaceDecl{
				InterfaceDecl: decl.GetInterfaceDecl(),
			},
		}
	case *Declaration_EnumDecl:
		return &TypeDeclaration{
			TypeDeclaration: &TypeDeclaration_EnumDecl{
				EnumDecl: decl.GetEnumDecl(),
			},
		}
	case *Declaration_TypeAliasDecl:
		return &TypeDeclaration{
			TypeDeclaration: &TypeDeclaration_TypeAliasDecl{
				TypeAliasDecl: decl.GetTypeAliasDecl(),
			},
		}
	case *Declaration_GenericParameter:
		return &TypeDeclaration{
			TypeDeclaration: &TypeDeclaration_GenericParameter{
				GenericParameter: decl.GetGenericParameter(),
			},
		}
	}

	return nil
}

func (m *TypeDeclaration) GetSourceFileName () string {
	if m == nil {
		return ""
	}

	switch m.TypeDeclaration.(type) {
	case *TypeDeclaration_StructDecl:
		return m.GetStructDecl().SourceFileName
	case *TypeDeclaration_EnumDecl:
		return m.GetEnumDecl().SourceFileName
	case *TypeDeclaration_InterfaceDecl:
		return m.GetInterfaceDecl().SourceFileName
	case *TypeDeclaration_TypeAliasDecl:
		return m.GetTypeAliasDecl().SourceFileName
	default:
		return ""
	}
}

func (m *TypeDeclaration) GetEnclosingType() *NominalType {
	if m == nil {
		return nil
	}

	switch m.TypeDeclaration.(type) {
	case *TypeDeclaration_StructDecl:
		return m.GetStructDecl().EnclosingType
	case *TypeDeclaration_InterfaceDecl:
		return m.GetInterfaceDecl().EnclosingType
	case *TypeDeclaration_EnumDecl:
		return m.GetEnumDecl().EnclosingType
	case *TypeDeclaration_TypeAliasDecl:
		return m.GetTypeAliasDecl().EnclosingType
	case *TypeDeclaration_GenericParameter:
		return m.GetGenericParameter().EnclosingType
	default:
		return nil
	}
}

func (m *TypeDeclaration) Merge(dependencies ...[]*Identifier) {
	var merged [][]*Identifier
	for _, dep := range dependencies {
		merged = append(merged, dep)
	}

	switch m.TypeDeclaration.(type) {
	case *TypeDeclaration_StructDecl:
		decl := m.GetStructDecl()
		merged = append(merged, decl.ResolvedIdentifiers)
		decl.ResolvedIdentifiers = MergeDependencies(merged...)
	case *TypeDeclaration_EnumDecl:
		decl := m.GetEnumDecl()
		merged = append(merged, decl.ResolvedIdentifiers)
		decl.ResolvedIdentifiers = MergeDependencies(merged...)
	case *TypeDeclaration_InterfaceDecl:
		decl := m.GetInterfaceDecl()
		merged = append(merged, decl.ResolvedIdentifiers)
		decl.ResolvedIdentifiers = MergeDependencies(merged...)
	case *TypeDeclaration_TypeAliasDecl:
		decl := m.GetTypeAliasDecl()
		merged = append(merged, decl.ResolvedIdentifiers)
		decl.ResolvedIdentifiers = MergeDependencies(merged...)
	}
}

