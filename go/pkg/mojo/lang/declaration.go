package lang

import (
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"strings"
)

const (
	EnumDeclName      = "EnumDecl"
	StructDeclName    = "StructDecl"
	InterfaceDeclName = "InterfaceDecl"
)

var declInfos map[reflect.Type]StructJsonInfo

func init() {
	declInfos = GetStructJsonInfos(jsoniter.Config{},
		PackageDecl{},
		ImportDecl{},
		ConstantDecl{},
		EnumDecl{},
		StructDecl{},
		InterfaceDecl{},
		FunctionDecl{},
	)
}

func NewEnumDeclaration(decl *EnumDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_EnumDecl{
			EnumDecl: decl,
		},
	}
}

func NewStructDeclaration(decl *StructDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_StructDecl{
			StructDecl: decl,
		},
	}
}

func NewInterfaceDeclaration(decl *InterfaceDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_InterfaceDecl{
			InterfaceDecl: decl,
		},
	}
}

func NewTypeAliasDeclaration(decl *TypeAliasDecl) *Declaration {
	return &Declaration{
		Declaration: &Declaration_TypeAliasDecl{
			TypeAliasDecl: decl,
		},
	}
}

func NewDeclarationFromTypeDeclaration(decl *TypeDeclaration) *Declaration {
	if decl != nil {
		switch decl.TypeDeclaration.(type) {
		case *TypeDeclaration_StructDecl:
			return &Declaration{
				Declaration: &Declaration_StructDecl{
					StructDecl: decl.GetStructDecl(),
				},
			}
		case *TypeDeclaration_InterfaceDecl:
			return &Declaration{
				Declaration: &Declaration_InterfaceDecl{
					InterfaceDecl: decl.GetInterfaceDecl(),
				},
			}
		case *TypeDeclaration_EnumDecl:
			return &Declaration{
				Declaration: &Declaration_EnumDecl{
					EnumDecl: decl.GetEnumDecl(),
				},
			}
		case *TypeDeclaration_TypeAliasDecl:
			return &Declaration{
				Declaration: &Declaration_TypeAliasDecl{
					TypeAliasDecl: decl.GetTypeAliasDecl(),
				},
			}
		case *TypeDeclaration_GenericParameter:
			return &Declaration{
				Declaration: &Declaration_GenericParameter{
					GenericParameter: decl.GetGenericParameter(),
				},
			}
		default:
			return nil
		}
	}
	return nil
}

func NewGenericParameterDeclaration(parameter *GenericParameter) *Declaration {
	return &Declaration{
		Declaration: &Declaration_GenericParameter{
			GenericParameter: parameter,
		},
	}
}

func (m *Declaration) GetName() string {
	if m != nil {
		switch m.Declaration.(type) {
		case *Declaration_StructDecl:
			return m.GetStructDecl().Name
		case *Declaration_InterfaceDecl:
			return m.GetInterfaceDecl().Name
		case *Declaration_EnumDecl:
			return m.GetEnumDecl().Name
		case *Declaration_TypeAliasDecl:
			return m.GetTypeAliasDecl().Name
		case *Declaration_GenericParameter:
			return m.GetGenericParameter().Name
		default:
			return ""
		}
	}
	return ""
}

func (m *Declaration) GetEnclosingType() *NominalType {
	if m != nil {
		switch m.Declaration.(type) {
		case *Declaration_StructDecl:
			return m.GetStructDecl().EnclosingType
		case *Declaration_InterfaceDecl:
			return m.GetInterfaceDecl().EnclosingType
		case *Declaration_EnumDecl:
			return m.GetEnumDecl().EnclosingType
		case *Declaration_TypeAliasDecl:
			return m.GetTypeAliasDecl().EnclosingType
		case *Declaration_GenericParameter:
			return m.GetGenericParameter().EnclosingType
		default:
			return nil
		}
	}
	return nil
}

func (m *Declaration) GetPackageName() string {
	if m != nil {
		switch m.Declaration.(type) {
		case *Declaration_StructDecl:
			return m.GetStructDecl().PackageName
		case *Declaration_InterfaceDecl:
			return m.GetInterfaceDecl().PackageName
		case *Declaration_EnumDecl:
			return m.GetEnumDecl().PackageName
		case *Declaration_TypeAliasDecl:
			return m.GetTypeAliasDecl().PackageName
		case *Declaration_GenericParameter:
			return m.GetGenericParameter().PackageName
		default:
			return ""
		}
	}
	return ""
}

func (m *Declaration) SetPackageName(name string) {
	if m != nil {
		switch m.Declaration.(type) {
		case *Declaration_StructDecl:
			m.GetStructDecl().PackageName = name
		case *Declaration_InterfaceDecl:
			m.GetInterfaceDecl().PackageName = name
		case *Declaration_EnumDecl:
			m.GetEnumDecl().PackageName = name
		case *Declaration_TypeAliasDecl:
			m.GetTypeAliasDecl().PackageName = name
		case *Declaration_GenericParameter:
			m.GetGenericParameter().PackageName = name
		}
	}
}

func (m *Declaration) IsGeneric() bool {
	switch decl := m.Declaration.(type) {
	case *Declaration_TypeAliasDecl:
		return decl.TypeAliasDecl.IsGeneric()
	case *Declaration_StructDecl:
		return decl.StructDecl.IsGeneric()
	}
	return false
}

func (m *Declaration) IsInstantiatedGeneric() bool {
	return strings.ContainsAny(m.GetStructDecl().GetName(), "<>")
}
