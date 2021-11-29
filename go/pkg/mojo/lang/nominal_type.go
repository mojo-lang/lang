package lang

import (
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"strings"
)

func (m *NominalType) GetFullName() string {
	if m != nil {
		if strings.Contains(m.Name, ".") {
			return GetFullName(m.PackageName, nil, m.Name)
		} else {
			return GetFullName(m.PackageName, m.GetEnclosingNames(), m.Name)
		}
	}
	return ""
}

func (m *NominalType) GetEnclosingNames() []string {
	if m == nil {
		return []string{}
	}
	return GetEnclosingNames(m.EnclosingType)
}

func (m *NominalType) IsTypeAlias() bool {
	if m == nil {
		return false
	}
	return m.TypeDeclaration != nil && m.TypeDeclaration.GetTypeAliasDecl() != nil
}

func (m *NominalType) IsInstantiatedGeneric() bool {
	if m.IsGeneric() {
		for _, argument := range m.GetGenericArguments() {
			if argument.IsGenericParameter() {
				return false
			}
		}
		return true
	}
	return false
}

func (m *NominalType) IsGeneric() bool {
	return len(m.GetGenericArguments()) > 0
}

func (m *NominalType) IsGenericParameter() bool {
	if m == nil {
		return false
	}
	return m.TypeDeclaration != nil && m.TypeDeclaration.GetGenericParameter() != nil
}

func (m *NominalType) IsStruct() bool {
	if m == nil {
		return false
	}
	return m.TypeDeclaration != nil && m.TypeDeclaration.GetStructDecl() != nil
}

func (m *NominalType) IsInterface() bool {
	if m == nil {
		return false
	}
	return m.TypeDeclaration != nil && m.TypeDeclaration.GetInterfaceDecl() != nil
}

func (m *NominalType) IsEnum() bool {
	if m == nil {
		return false
	}
	return m.TypeDeclaration != nil && m.TypeDeclaration.GetEnumDecl() != nil
}

func (m *NominalType) IsScalar() bool {
	if m == nil {
		return false
	}
	if m.TypeDeclaration != nil {
		if decl := m.TypeDeclaration.GetStructDecl(); decl != nil {
			switch decl.GetFullName() {
			case core.BoolTypeName, core.Int8TypeName, core.Int16TypeName, core.Int32TypeName, core.Int64TypeName,
				core.UInt8TypeName, core.UInt16TypeName, core.UInt32TypeName, core.UInt64TypeName, core.Float32TypeName,
				core.Float64TypeName, core.StringTypeName, core.BytesTypeName:
				return true
			default:
				return false
			}
		} else if alias := m.TypeDeclaration.GetTypeAliasDecl(); alias != nil {
			switch alias.GetFullName() {
			case core.IntTypeName, core.UIntTypeName, core.FloatTypeName, core.DoubleTypeName:
				return true
			default:
				return false
			}
		}
	}
	return false
}

func (m *NominalType) IsArrayType() bool {
	return m.isType(core.ArrayTypeName)
}

func (m *NominalType) IsMapType() bool {
	return m.isType(core.MapTypeName)
}

func (m *NominalType) IsTupleType() bool {
	return m.isType(core.TupleTypeName)
}

func (m *NominalType) IsIntersectionType() bool {
	return m.isType(core.IntersectionTypeName)
}

func (m *NominalType) IsUnionType() bool {
	return m.isType(core.UnionTypeName)
}

func (m *NominalType) isType(typeName string) bool {
	if m == nil {
		return false
	}
	if m.TypeDeclaration != nil {
		if decl := m.TypeDeclaration.GetStructDecl(); decl != nil {
			return decl.GetFullName() == typeName
		}
	}
	return false
}
