package lang

import "strings"

func ParseIdentifierName(fullName string) (string, string) {
	index := strings.LastIndex(fullName, ".")
	if index > 0 {
		return fullName[0:index], fullName[index+1:]
	} else {
		return "", fullName
	}
}

func FindIdentifier(identifiers []*Identifier, fullName string) *Identifier {
	for _, id := range identifiers {
		if id.FullName == fullName {
			return id
		}
	}
	return nil
}

func (m *Identifier) IsGenericInstantiated() bool {
	return strings.Contains(m.GetName(), "<")
}

func (m *Identifier) ToNominalType() *NominalType {
	if m != nil {
		switch m.Kind {
		case Identifier_KIND_ENUM, Identifier_KIND_STRUCT, Identifier_KIND_TYPE_ALIAS, Identifier_KIND_INTERFACE:
			return &NominalType{
				StartPosition:    m.StartPosition,
				EndPosition:      m.EndPosition,
				PackageName:      m.PackageName,
				Name:             m.Name,
				TypeDeclaration:  NewTypeDeclarationFromDeclaration(m.Declaration),
				GenericArguments: nil,
				Attributes:       nil,
				EnclosingType:    nil,
			}
		}
	}
	return nil
}
