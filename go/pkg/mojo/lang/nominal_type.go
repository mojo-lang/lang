package lang

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/mojo-lang/core/go/pkg/mojo/core"
)

func (m *NominalType) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *NominalType) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *NominalType) MergeComment(comment *Comment) (bool, error) {
	if m != nil {
		return MergeCommentToTerms(comment, m.GetTerms())
	}

	return false, errors.New("nil NominalType")
}

func (m *NominalType) GetTerms() []interface{} {
	if m != nil {
		var terms []interface{}

		//terms = append(terms, NewSymbolTerm(m.StartPosition, TermTypeName, m.Name))

		for _, attribute := range m.Attributes {
			terms = append(terms, attribute)
		}

		if m.Document != nil && m.Document.Following {
			terms = append(terms, m.Document)
		}

		return terms
	}
	return nil
}

func (m *NominalType) NewIdentifier() *Identifier {
	if m != nil {
		identifier := &Identifier{
			StartPosition:      m.StartPosition,
			EndPosition:        m.EndPosition,
			Kind:               Identifier_KIND_UNSPECIFIED,
			PackageName:        m.PackageName,
			SourceFileName:     m.TypeDeclaration.GetSourceFileName(),
			EnclosingTypeNames: GetEnclosingNames(m.EnclosingType),
			Name:               m.Name,
			FullName:           m.GetFullName(),
			Declaration:        NewDeclarationFromTypeDeclaration(m.TypeDeclaration),
			Implicit:           m.TypeDeclaration.Implicit(),
		}

		switch m.TypeDeclaration.TypeDeclaration.(type) {
		case *TypeDeclaration_EnumDecl:
			identifier.Kind = Identifier_KIND_ENUM
		case *TypeDeclaration_StructDecl:
			identifier.Kind = Identifier_KIND_STRUCT
		case *TypeDeclaration_TypeAliasDecl:
			identifier.Kind = Identifier_KIND_TYPE_ALIAS
		case *TypeDeclaration_InterfaceDecl:
			identifier.Kind = Identifier_KIND_INTERFACE
		}
		return identifier
	}
	return nil
}

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

func (m *NominalType) GetGenericFullName() string {
	if m != nil {
		if strings.Contains(m.Name, ".") {
			return GetFullName(m.PackageName, nil, m.GetGenericName())
		} else {
			return GetFullName(m.PackageName, m.GetEnclosingGenericNames(), m.GetGenericName())
		}
	}
	return ""
}

func (m *NominalType) GetGenericName() string {
	if m != nil {
		name := m.GetName()
		buf := bytes.NewBufferString(name)
		if len(m.GenericArguments) > 0 {
			buf.WriteByte('<')
			for i, argument := range m.GenericArguments {
				if i > 0 {
					buf.WriteByte(',')
				}
				buf.WriteString(argument.GetGenericFullName())
			}
			buf.WriteByte('>')
		}
		return buf.String()
	}
	return ""
}

func (m *NominalType) GetEnclosingNames() []string {
	if m == nil {
		return []string{}
	}
	return GetEnclosingNames(m.EnclosingType)
}

func (m *NominalType) GetEnclosingGenericNames() []string {
	if m == nil {
		return []string{}
	}
	return GetEnclosingGenericNames(m.EnclosingType)
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
			case core.BoolTypeFullName, core.Int8TypeFullName, core.Int16TypeFullName, core.Int32TypeFullName, core.Int64TypeFullName,
				core.UInt8TypeFullName, core.UInt16TypeFullName, core.UInt32TypeFullName, core.UInt64TypeFullName, core.Float32TypeFullName,
				core.Float64TypeFullName, core.StringTypeFullName, core.BytesTypeFullName:
				return true
			default:
				return false
			}
		} else if alias := m.TypeDeclaration.GetTypeAliasDecl(); alias != nil {
			switch alias.GetFullName() {
			case core.IntTypeFullName, core.UIntTypeFullName, core.FloatTypeFullName, core.DoubleTypeFullName:
				return true
			default:
				return false
			}
		}
	}
	return false
}

func (m *NominalType) IsArrayType() bool {
	return m.isType(core.ArrayTypeFullName)
}

func (m *NominalType) IsMapType() bool {
	return m.isType(core.MapTypeFullName)
}

func (m *NominalType) IsTupleType() bool {
	return m.isType(core.TupleTypeFullName)
}

func (m *NominalType) IsIntersectionType() bool {
	return m.isType(core.IntersectionTypeFullName)
}

func (m *NominalType) IsUnionType() bool {
	return m.isType(core.UnionTypeFullName)
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

// support 'mojo.alias.Foo<mojo.alias.Bar>'
// support 'mojo.alias.Foo'
func ParseNominalTypeFullName(name string) (*NominalType, error) {
	nominalType := &NominalType{}
	return nominalType, nominalType.ParseFullName(name)
}

func (m *NominalType) ParseFullName(fullName string) error {
	_, err := m.parseFullGenericName(fullName)
	return err
}

func (m *NominalType) parseFullGenericName(fullName string) (string, error) {
	if m != nil {
		var err error
		originalFullName := fullName
		fullName = strings.TrimSpace(fullName)
		fullName, err = m.parseFullName(fullName)
		if err != nil {
			return "", err
		}

		fullName = strings.TrimSpace(fullName)
		if len(fullName) > 0 {
			if strings.HasPrefix(fullName, "<") {
				fullName = fullName[1:]
				fullName = strings.TrimSpace(fullName)
				for !strings.HasPrefix(fullName, ">") {
					argument := &NominalType{}
					fullName, err = argument.parseFullGenericName(fullName)
					if err != nil {
						return "", err
					}
					m.GenericArguments = append(m.GenericArguments, argument)
					fullName = strings.TrimSpace(fullName)

					if len(fullName) == 0 {
						return "", fmt.Errorf("malformed generic full name (%s)", originalFullName)
					}
				}
				fullName = fullName[1:]
				fullName = strings.TrimSpace(fullName)
				if strings.HasPrefix(fullName, ",") {
					fullName = fullName[1:]
					fullName = strings.TrimSpace(fullName)
				}
				return fullName, nil
			} else if strings.HasPrefix(fullName, ",") {
				fullName = strings.TrimPrefix(fullName, ",")
				fullName = strings.TrimSpace(fullName)
				return fullName, nil
			} else if strings.HasPrefix(fullName, ">") {
				return fullName, nil
			} else {
				return "", fmt.Errorf("malformed generic full name (%s)", originalFullName)
			}
		}
		return "", nil
	}
	return "", errors.New("nil NominalType")
}

func (m *NominalType) parseFullName(fullName string) (string, error) {
	if m != nil {
		if pkgReg, err := regexp.Compile(`^([a-z][a-z_0-9]*\.)+`); err != nil {
			return "", err
		} else {
			pkg := pkgReg.FindString(fullName)
			if len(pkg) > 0 {
				m.PackageName = strings.TrimSuffix(pkg, ".")
				fullName = fullName[len(pkg):]
			}
		}

		if enclosingReg, err := regexp.Compile(`^([A-Z][a-zA-Z0-9_]*\.)+`); err != nil {
			return "", err
		} else {
			enclosing := enclosingReg.FindString(fullName)
			fullName = fullName[len(enclosing):]
			enclosing = strings.TrimSuffix(enclosing, ".")
			if len(enclosing) > 0 {
				segments := strings.Split(enclosing, ".")
				var enclosingType *NominalType
				for _, segment := range segments {
					if enclosingType == nil {
						enclosingType = &NominalType{
							PackageName: m.PackageName,
							Name:        segment,
						}
					} else {
						enclosingType = &NominalType{
							PackageName:   m.PackageName,
							Name:          segment,
							EnclosingType: enclosingType,
						}
					}
				}
				m.EnclosingType = enclosingType
			}
		}

		if nameReg, err := regexp.Compile(`^[A-Z][a-zA-Z0-9_]*`); err != nil {
			return "", err
		} else {
			name := nameReg.FindString(fullName)
			if len(name) > 0 {
				m.Name = name
				fullName = fullName[len(name):]
			}
		}
	}

	return fullName, nil
}
