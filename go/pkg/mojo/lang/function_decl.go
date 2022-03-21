package lang

import (
    "errors"
    "github.com/golang/protobuf/proto"
)

func (m *FunctionDecl) SetStartPosition(position *Position) {
    if m != nil {
        m.StartPosition = PatchPosition(m.StartPosition, position)
    }
}

func (m *FunctionDecl) SetEndPosition(position *Position) {
    if m != nil {
        m.EndPosition = PatchPosition(m.EndPosition, position)
    }
}

func (m *FunctionDecl) MergeComment(comment *Comment) (bool, error) {
    if m != nil {
        return MergeCommentToTerms(comment, m.GetTerms())
    }

    return false, errors.New("nil StructDecl")
}

func (m *FunctionDecl) GetTerms() []interface{} {
    if m != nil {
        var terms []interface{}

        if m.Document != nil {
            terms = append(terms, m.Document)
        }
        for _, attribute := range m.Attributes {
            terms = append(terms, attribute)
        }

        if m.Signature != nil {
            terms = append(terms, m.Signature)
        }

        if m.Body != nil {
            terms = append(terms, m.Body)
        }

        return terms
    }
    return nil
}

func (m *FunctionDecl) GetAttributeArguments(name string) ([]*Argument, error) {
    if m != nil {
        return GetAttributeArguments(m.Attributes, name)
    }
    return nil, errors.New("FunctionDecl is nil")
}

func (m *FunctionDecl) GetAttributeArgument(name string) (*Argument, error) {
    if m != nil {
        return GetAttributeArgument(m.Attributes, name)
    }
    return nil, errors.New("FunctionDecl is nil")
}

func (m *FunctionDecl) HasAttribute(name string) bool {
    if m != nil {
        return HasAttribute(m.Attributes, name)
    }
    return false
}

func (m *FunctionDecl) GetAttribute(name string) *Attribute {
    if m != nil {
        return GetAttribute(m.Attributes, name)
    }
    return nil
}

func (m *FunctionDecl) GetBoolAttribute(name string) (bool, error) {
    argument, err := m.GetAttributeArgument(name)
    if err != nil {
        return false, err
    }

    return argument.GetBool()
}

func (m *FunctionDecl) SetBoolAttribute(name string, value bool) *Attribute {
    if m != nil {
        m.Attributes = SetBoolAttribute(m.Attributes, name, value)
        return m.Attributes[len(m.Attributes)-1]
    }
    return nil
}

func (m *FunctionDecl) SetImplicitBoolAttribute(name string, value bool) *Attribute {
    if m != nil {
        return m.SetBoolAttribute(name, value).SetImplicit(true)
    }
    return nil
}

func (m *FunctionDecl) GetIntegerAttribute(name string) (int64, error) {
    if argument, err := m.GetAttributeArgument(name); err != nil {
        return 0, err
    } else {
        return argument.GetInteger()
    }
}

func (m *FunctionDecl) SetIntegerAttribute(name string, value int64) *Attribute {
    if m != nil {
        m.Attributes = SetIntegerAttribute(m.Attributes, name, value)
        return m.Attributes[len(m.Attributes)-1]
    }
    return nil
}

func (m *FunctionDecl) SetImplicitIntegerAttribute(name string, value int64) *Attribute {
    if m != nil {
        return m.SetIntegerAttribute(name, value).SetImplicit(true)
    }
    return nil
}

func (m *FunctionDecl) GetStringAttribute(name string) (string, error) {
    if argument, err := m.GetAttributeArgument(name); err != nil {
        return "", err
    } else {
        return argument.GetString()
    }
}

func (m *FunctionDecl) SetStringAttribute(name string, value string) *Attribute {
    if m != nil {
        m.Attributes = SetStringAttribute(m.Attributes, name, value)
        return m.Attributes[len(m.Attributes)-1]
    }
    return nil
}

func (m *FunctionDecl) SetImplicitStringAttribute(name string, value string) *Attribute {
    if m != nil {
        return m.SetStringAttribute(name, value).SetImplicit(true)
    }
    return nil
}

func (m *FunctionDecl) RemoveAttribute(name string) {
    if m != nil {
        m.Attributes = RemoveAttribute(m.Attributes, name)
    }
}

func (m *FunctionDecl) GetResource() string {
    if m != nil {
        //if m.GetStringAttribute()
    }
    return ""
}

// Copy shallow copy for the FunctionDecl
// deepFields will using deep copy
func (m *FunctionDecl) Copy(deepFields ...string) *FunctionDecl {
    decl := &FunctionDecl{
        StartPosition:     m.StartPosition,
        EndPosition:       m.EndPosition,
        Implicit:          m.Implicit,
        Document:          m.Document,
        PackageName:       m.PackageName,
        SourceFileName:    m.SourceFileName,
        KeywordPosition:   m.KeywordPosition,
        Name:              m.Name,
        FullName:          m.FullName,
        Attributes:        m.Attributes,
        GenericParameters: m.GenericParameters,
        EnclosingType:     m.EnclosingType,
        NamePosition:      m.NamePosition,
        Signature:         m.Signature,
        Body:              m.Body,
        Scope:             m.Scope,
    }

    index := make(map[string]bool)
    for _, field := range deepFields {
        index[field] = true
    }

    if index["attributes"] {
        var attributes []*Attribute
        for _, attribute := range m.Attributes {
            attributes = append(attributes, proto.Clone(attribute).(*Attribute))
        }
        decl.Attributes = attributes
    }

    if index["signature"] {
        decl.Signature = proto.Clone(m.Signature).(*FunctionSignature)
    }

    if index["body"] {
        decl.Body = proto.Clone(m.Body).(*BlockStmt)
    }

    return decl
}
