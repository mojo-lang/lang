package lang

import (
    "errors"
    "github.com/mojo-lang/core/go/pkg/mojo/core"
)

func NewBoolAttribute(pkg string, name string) *Attribute {
    return newBoolAttribute(pkg, name, false)
}

func NewStringAttribute(pkg string, name string, value string) *Attribute {
    return newStringAttribute(pkg, name, value, false)
}

func NewImplicitStringAttribute(pkg string, name string, value string) *Attribute {
    return newStringAttribute(pkg, name, value, false)
}

func NewImplicitBoolAttribute(pkg string, name string) *Attribute {
    return newBoolAttribute(pkg, name, true)
}

// IsNumber return true if the attribute is number attribute
func (m *Attribute) IsNumber() bool {
    return m != nil && m.Name == core.NumberAttributeName
}

func (m *Attribute) IsRequired() bool {
    return m != nil && m.Name == core.RequiredAttributeName
}

func (m *Attribute) IsOptional() bool {
    return m != nil && m.Name == core.OptionalAttributeName
}

func (m *Attribute) SetStartPosition(position *Position) {
    if m != nil {
        m.StartPosition = PatchPosition(m.StartPosition, position)
    }
}

func (m *Attribute) SetEndPosition(position *Position) {
    if m != nil {
        m.EndPosition = PatchPosition(m.EndPosition, position)
    }
}

func (m *Attribute) MergeComment(comment *Comment) (bool, error) {
    if m != nil {
        return MergeCommentToTerms(comment, m.GetTerms())
    }

    return false, errors.New("nil ArrayLiteralExpr")
}

func (m *Attribute) GetTerms() []interface{} {
    if m != nil {
        var terms []interface{}
        terms = append(terms, &Term{StartPosition: m.NamePosition,
            EndPosition: &Position{
                Line:   m.NamePosition.Line,
                Column: m.NamePosition.Column + int64(len(m.Name)),
            },
            Type:  "Name",
            Value: m.Name,
        })

        for _, argument := range m.GenericArguments {
            terms = append(terms, argument)
        }
        for _, argument := range m.Arguments {
            terms = append(terms, argument)
        }
        return terms
    }
    return nil
}

func (m *Attribute) IsSameName(fullName string) bool {
    if m == nil {
        return false
    }

    pkg, name := ParseIdentifierName(fullName)
    if len(pkg) > 0 {
        return m.PackageName == pkg && m.Name == name
    } else {
        return m.Name == name
    }
}

func (m *Attribute) GetFullName() string {
    fullName := ""
    if m != nil {
        if len(m.PackageName) > 0 {
            fullName = m.PackageName + "."
        }
        fullName = fullName + m.Name
    }
    return fullName
}

func (m *Attribute) Repeatable() bool {
    if m != nil && m.Declaration != nil {
        if v, err := GetBoolAttribute(m.Declaration.Attributes, "repeatable"); err == nil {
            return v
        }
    }
    return false
}

func (m *Attribute) GetArrayLiteralArgument() *ArrayLiteralExpr {
    array := &ArrayLiteralExpr{
        Implicit: true,
    }

    if len(m.Arguments) == 1 {
        if arrayLiteral := m.Arguments[0].Value.GetArrayLiteralExpr(); arrayLiteral != nil {
            return arrayLiteral
        }
    }

    for _, argument := range m.Arguments {
        if len(argument.Label) > 0 {
            return nil
        }

        array.Elements = append(array.Elements, argument.Value)
    }
    return array
}

func (m *Attribute) GetObjectLiteralArgument() *ObjectLiteralExpr {
    if m != nil {
        object := &ObjectLiteralExpr{
            Implicit: true,
        }
        if len(m.Arguments) == 1 {
            if objectLiteral := m.Arguments[0].Value.GetObjectLiteralExpr(); objectLiteral != nil {
                return objectLiteral
            }
        }

        object, _ = Arguments(m.Arguments).ToObjectLiteralExpr()
        return object
    }
    return nil
}

func (m *Attribute) GetMapLiteralArgument() *MapLiteralExpr {
    if len(m.Arguments) == 1 {
        if mapLiteral := m.Arguments[0].Value.GetMapLiteralExpr(); mapLiteral != nil {
            return mapLiteral
        }
    }

    object := m.GetObjectLiteralArgument()
    if object != nil {
        mapLiteral := &MapLiteralExpr{
            Implicit: true,
        }
        for _, field := range object.Fields {
            mapLiteral.Entries = append(mapLiteral.Entries, &MapLiteralExpr_Entry{
                Key:   field.Name,
                Value: field.Value,
            })
        }
        return mapLiteral
    }
    return nil
}

func (m *Attribute) GetObjectArgument(object interface{}) error {
    o := m.GetObjectLiteralArgument()
    if o != nil {
        return o.To(object)
    }
    return errors.New("not an object literal argument")
}

func (m *Attribute) SetImplicit(value bool) *Attribute {
    if m != nil {
        m.Implicit = value
    }
    return m
}

// GetBool eval the argument first, then use the default value of the attribute declaration if no argument
func (m *Attribute) GetBool() bool {
    if m != nil {
        if len(m.Arguments) == 0 {
            if m.Declaration != nil && m.Declaration.DefaultValue != nil {
                val, _ := m.Declaration.DefaultValue.EvalBoolLiteral()
                return val
            }
            return true
        }
        val, _ := m.Arguments[0].GetValue().EvalBoolLiteral()
        return val
    }
    return false
}

func (m *Attribute) GetInteger() (int64, error) {
    if m != nil {
        if len(m.Arguments) > 0 {
            return m.Arguments[0].GetInteger()
        } else {
            return m.GetValue().EvalIntegerLiteral()
        }
    }
    return 0, errors.New("")
}

func newBoolAttribute(pkg string, name string, implicit bool) *Attribute {
    return &Attribute{
        PackageName: pkg,
        Name:        name,
        Arguments: []*Argument{{
            Value: NewBoolLiteralExpression(&BoolLiteralExpr{
                Kind:     0,
                Implicit: implicit,
                Value:    true,
            }),
        },
        },
    }
}

func newStringAttribute(pkg string, name string, value string, implicit bool) *Attribute {
    return &Attribute{
        PackageName: pkg,
        Name:        name,
        Arguments: []*Argument{{
            Value: NewStringLiteralExpression(&StringLiteralExpr{
                Kind:     0,
                Implicit: implicit,
                Value:    value,
            }),
        },
        },
    }
}
