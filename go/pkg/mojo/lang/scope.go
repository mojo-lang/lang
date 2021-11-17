package lang

import (
	"strings"
)

func NewScope() *Scope {
	return &Scope{
		Identifiers: make(map[string]*Identifier),
	}
}

func (m *Scope) Declare(declaration *Declaration) *Identifier {
	if declaration == nil {
		return nil
	}

	switch declaration.Declaration.(type) {
	case *Declaration_StructDecl:
		decl := declaration.GetStructDecl()

		packageName := decl.PackageName
		enclosingNames := GetEnclosingNames(decl.EnclosingType)
		fullName := GetFullName(packageName, enclosingNames, decl.Name)

		identifier := &Identifier{
			PackageName:        packageName,
			EnclosingTypeNames: enclosingNames,
			Kind:               Identifier_KIND_STRUCT,
			Name:               decl.Name,
			FullName:           fullName,
			Declaration:        declaration,
		}
		m.Identifiers[identifier.Name] = identifier
		return identifier
	case *Declaration_EnumDecl:
		decl := declaration.GetEnumDecl()

		packageName := decl.PackageName
		enclosingNames := GetEnclosingNames(decl.EnclosingType)
		fullName := GetFullName(packageName, enclosingNames, decl.Name)

		identifier := &Identifier{
			PackageName:        packageName,
			EnclosingTypeNames: enclosingNames,
			Kind:               Identifier_KIND_ENUM,
			Name:               decl.Name,
			FullName:           fullName,
			Declaration:        declaration,
		}
		m.Identifiers[identifier.Name] = identifier
		return identifier
	case *Declaration_InterfaceDecl:
		decl := declaration.GetInterfaceDecl()

		packageName := decl.PackageName
		fullName := strings.Join([]string{packageName, decl.Name}, ".")

		identifier := &Identifier{
			PackageName:        packageName,
			EnclosingTypeNames: nil,
			Kind:               Identifier_KIND_INTERFACE,
			Name:               decl.Name,
			FullName:           fullName,
			Declaration:        declaration,
		}
		m.Identifiers[identifier.Name] = identifier
		return identifier
	case *Declaration_TypeAliasDecl:
		decl := declaration.GetTypeAliasDecl()

		packageName := decl.PackageName
		fullName := strings.Join([]string{packageName, decl.Name}, ".")

		identifier := &Identifier{
			PackageName:        packageName,
			EnclosingTypeNames: nil,
			Kind:               Identifier_KIND_TYPE_ALIAS,
			Name:               decl.Name,
			FullName:           fullName,
			Declaration:        declaration,
		}
		m.Identifiers[identifier.Name] = identifier
		return identifier
	case *Declaration_ImportDecl:
		//err := c.compileImport(c.Context, decl.GetImportDecl())
		//if err != nil {
		//	return err
		//}
		return nil
	case *Declaration_GenericParameter:
		parameter := declaration.GetGenericParameter()

		identifier := &Identifier{
			PackageName: parameter.PackageName,
			Kind:        Identifier_KIND_GENERIC_PARAMETER,
			Name:        parameter.Name,
			Declaration: declaration,
		}
		m.Identifiers[identifier.Name] = identifier
		return identifier
	}
	return nil
}

func (m *Scope) GetIdentifier(name string) *Identifier {
	if m != nil && m.Identifiers != nil {
		return m.Identifiers[GetTypeTypeName(name)]
	}
	return nil
}
