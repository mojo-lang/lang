package lang

import "strings"

func (m *NominalType) GetFullName() string {
	if m != nil {
		if strings.Contains(m.Name, ".") {
			return GetFullName(m.Package, nil, m.Name)
		} else {
			return GetFullName(m.Package, m.GetEnclosingNames(), m.Name)
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

func (m *NominalType) IsArrayType() bool {
	return m.isType("mojo.core.Array")
}

func (m *NominalType) IsDictionaryType() bool {
	return m.isType("mojo.core.Dictionary")
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
