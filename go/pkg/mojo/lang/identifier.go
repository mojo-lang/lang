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

func (x *Identifier) SetStartPosition(position *Position) {
    if x != nil {
        x.StartPosition = PatchPosition(x.StartPosition, position)
    }
}

func (x *Identifier) SetEndPosition(position *Position) {
    if x != nil {
        x.EndPosition = PatchPosition(x.EndPosition, position)
    }
}

func (x *Identifier) IsGenericInstantiated() bool {
    return strings.Contains(x.GetName(), "<")
}

func (x *Identifier) ToNominalType() *NominalType {
    if x != nil {
        switch x.Kind {
        case Identifier_KIND_ENUM, Identifier_KIND_STRUCT, Identifier_KIND_TYPE_ALIAS, Identifier_KIND_INTERFACE:
            return &NominalType{
                StartPosition:    x.StartPosition,
                EndPosition:      x.EndPosition,
                PackageName:      x.PackageName,
                Name:             x.Name,
                TypeDeclaration:  NewTypeDeclarationFromDeclaration(x.Declaration),
                GenericArguments: nil,
                Attributes:       nil,
                EnclosingType:    nil,
            }
        }
    }
    return nil
}
