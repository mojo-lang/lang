package lang

import "strings"

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
			case "mojo.core.Bool", "mojo.core.Int8", "mojo.core.Int16", "mojo.core.Int32", "mojo.core.Int64",
				"mojo.core.UInt8", "mojo.core.UInt16", "mojo.core.UInt32", "mojo.core.UInt64", "mojo.core.Float32",
				"mojo.core.Float64", "mojo.core.String", "mojo.core.Bytes":
				return true
			default:
				return false
			}
		} else if alias := m.TypeDeclaration.GetTypeAliasDecl(); alias != nil {
			switch alias.GetFullName() {
			case "mojo.core.Int", "mojo.core.UInt", "mojo.core.Float", "mojo.core.Double":
				return true
			default:
				return false
			}
		}
	}
	return false
}

func (m *NominalType) IsArrayType() bool {
	return m.isType("mojo.core.Array")
}

func (m *NominalType) IsMapType() bool {
	return m.isType("mojo.core.Map")
}

func (m *NominalType) IsTupleType() bool {
	return m.isType("mojo.core.Tuple")
}

func (m *NominalType) IsIntersectionType() bool {
	return m.isType("mojo.core.Intersection")
}

func (m *NominalType) IsUnionType() bool {
	return m.isType("mojo.core.Union")
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
