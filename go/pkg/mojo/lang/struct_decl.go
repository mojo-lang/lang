package lang

import (
    "errors"
)

func (m *StructDecl) GetFullName() string {
    if m != nil {
        return GetFullName(m.PackageName, GetEnclosingNames(m.EnclosingType), m.Name)
    }
    return ""
}

func (m *StructDecl) SetScope(scope *Scope) {
    if m != nil {
        m.Scope = scope
    }
}

func (m *StructDecl) SetStartPosition(position *Position) {
    if m != nil {
        m.StartPosition = PatchPosition(m.StartPosition, position)
    }
}

func (m *StructDecl) SetEndPosition(position *Position) {
    if m != nil {
        m.EndPosition = PatchPosition(m.EndPosition, position)
    }
}

func (m *StructDecl) IsBoxed() bool {
    if m != nil && m.Type != nil && len(m.Type.Inherits) == 1 && len(m.Type.Fields) == 0 {
        inherit := m.Type.Inherits[0]
        structDecl := inherit.TypeDeclaration.GetStructDecl()
        return structDecl != nil && structDecl.Type == nil
    }
    return false
}

func (m *StructDecl) IsOpacity() bool {
    return m != nil && m.Type == nil
}

func (m *StructDecl) IsGeneric() bool {
    return m != nil && len(m.GenericParameters) > 0
}

func (m *StructDecl) IsBodyEmpty() bool {
    if m != nil {
        return len(m.TypeAliasDecls) == 0 && len(m.EnumDecls) == 0 && len(m.StructDecls) == 0 && m.FieldCount() == 0
    }
    return true
}

func (m *StructDecl) FieldCount() int {
    return len(m.GetType().GetFields())
}

func (m *StructDecl) FieldNames(option FieldNamOption) []string {
    if m != nil {
        return m.Type.FieldNames(option)
    }
    return nil
}

func (m *StructDecl) GetField(name string) *ValueDecl {
    if m != nil {
        return m.Type.GetField(name)
    }
    return nil
}

func (m *StructDecl) EachField(iter func(decl *ValueDecl) error) error {
    if m != nil {
        return m.Type.EachField(iter)
    }
    return nil
}

func (m *StructDecl) EnclosingTypeDecl() interface{} {
    return m
}

func (m *StructDecl) GetAttributeArguments(name string) ([]*Argument, error) {
    if m != nil {
        return GetAttributeArguments(m.Attributes, name)
    }
    return nil, errors.New("StructDecl is nil")
}

func (m *StructDecl) GetAttributeArgument(name string) (*Argument, error) {
    if m != nil {
        return GetAttributeArgument(m.Attributes, name)
    }
    return nil, errors.New("StructDecl is nil")
}

func (m *StructDecl) HasAttribute(name string) bool {
    if m != nil {
        return HasAttribute(m.Attributes, name)
    }
    return false
}

func (m *StructDecl) GetAttribute(name string) *Attribute {
    if m != nil {
        return GetAttribute(m.Attributes, name)
    }
    return nil
}

func (m *StructDecl) GetBoolAttribute(name string) (bool, error) {
    argument, err := m.GetAttributeArgument(name)
    if err != nil {
        return false, err
    }

    return argument.GetBool()
}

func (m *StructDecl) SetBoolAttribute(name string, value bool) *Attribute {
    if m != nil {
        m.Attributes = SetBoolAttribute(m.Attributes, name, value)
        return m.Attributes[len(m.Attributes)-1]
    }
    return nil
}

func (m *StructDecl) SetImplicitBoolAttribute(name string, value bool) *Attribute {
    if m != nil {
        return m.SetBoolAttribute(name, value).SetImplicit(true)
    }
    return nil
}

func (m *StructDecl) GetIntegerAttribute(name string) (int64, error) {
    if argument, err := m.GetAttributeArgument(name); err != nil {
        return 0, err
    } else {
        return argument.GetInteger()
    }
}

func (m *StructDecl) SetIntegerAttribute(name string, value int64) *Attribute {
    if m != nil {
        m.Attributes = SetIntegerAttribute(m.Attributes, name, value)
        return m.Attributes[len(m.Attributes)-1]
    }
    return nil
}

func (m *StructDecl) SetImplicitIntegerAttribute(name string, value int64) *Attribute {
    if m != nil {
        return m.SetIntegerAttribute(name, value).SetImplicit(true)
    }
    return nil
}

func (m *StructDecl) GetStringAttribute(name string) (string, error) {
    if argument, err := m.GetAttributeArgument(name); err != nil {
        return "", err
    } else {
        return argument.GetString()
    }
}

func (m *StructDecl) SetStringAttribute(name string, value string) *Attribute {
    if m != nil {
        m.Attributes = SetStringAttribute(m.Attributes, name, value)
        return m.Attributes[len(m.Attributes)-1]
    }
    return nil
}

func (m *StructDecl) SetImplicitStringAttribute(name string, value string) *Attribute {
    if m != nil {
        return m.SetStringAttribute(name, value).SetImplicit(true)
    }
    return nil
}

func (m *StructDecl) RemoveAttribute(name string) {
    if m != nil {
        m.Attributes = RemoveAttribute(m.Attributes, name)
    }
}

func (m *StructDecl) MergeComment(comment *Comment) (bool, error) {
    if m != nil {
        return MergeCommentToTerms(comment, m.GetTerms()...)
    }

    return false, errors.New("nil StructDecl")
}

func (m *StructDecl) GetTerms() []interface{} {
    if m != nil {
        var terms []interface{}

        if m.Document != nil {
            terms = append(terms, m.Document)
        }

        for _, attribute := range m.Attributes {
            terms = append(terms, attribute)
        }

        terms = append(terms,
            NewSymbolTerm(m.KeywordPosition, TermTypeKeyword, KeywordType),
            NewSymbolTerm(m.NamePosition, TermTypeName, m.Name))

        for _, parameter := range m.GenericParameters {
            terms = append(terms, parameter)
        }

        if m.Type != nil {
            if len(m.Type.Inherits) > 0 {
                terms = append(terms, NewSymbolTerm(m.Type.InheritePosition, TermTypeSymbol, ":"))

                for _, inherit := range m.Type.Inherits {
                    terms = append(terms, inherit)
                }
            }

            terms = append(terms, NewSymbolTerm(m.Type.StartPosition, TermTypeStart, ""))

            for _, alias := range m.TypeAliasDecls {
                terms = append(terms, alias)
            }

            for _, enum := range m.EnumDecls {
                terms = append(terms, enum)
            }

            for _, structDecl := range m.StructDecls {
                terms = append(terms, structDecl)
            }

            for _, field := range m.Type.Fields {
                terms = append(terms, field)
            }

            terms = append(terms, NewSymbolTerm(m.Type.EndPosition, TermTypeEnd, ""))
        }

        return terms
    }
    return nil
}
