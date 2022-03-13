package lang

import (
    "errors"
)

func (m *AttributeDecl) SetStartPosition(position *Position) {
    if m != nil {
        m.StartPosition = PatchPosition(m.StartPosition, position)
    }
}

func (m *AttributeDecl) SetEndPosition(position *Position) {
    if m != nil {
        m.EndPosition = PatchPosition(m.EndPosition, position)
    }
}

func (m *AttributeDecl) MergeComment(comment *Comment) (bool, error) {
    if m != nil {
        return MergeCommentToTerms(comment, m.GetTerms())
    }

    return false, errors.New("nil AttributeDecl")
}

func (m *AttributeDecl) GetTerms() []interface{} {
    if m != nil {
        var terms []interface{}

        for _, attribute := range m.Attributes {
            terms = append(terms, attribute)
        }

        if m.Group == nil {
            terms = append(terms,
                &Term{StartPosition: m.KeywordPosition,
                    EndPosition: &Position{
                        Line:   m.KeywordPosition.Line,
                        Column: m.KeywordPosition.Column + int64(len(KeywordAttribute)),
                    },
                    Type:  "Keyword",
                    Value: KeywordAttribute,
                },
            )
        }

        terms = append(terms,
            &Term{StartPosition: m.NamePosition,
                EndPosition: &Position{
                    Line:   m.NamePosition.Line,
                    Column: m.NamePosition.Column + int64(len(m.Name)),
                },
                Type:  "Name",
                Value: m.Name,
            })

        for _, parameter := range m.GenericParameters {
            terms = append(terms, parameter)
        }

        if nominalType := m.GetNominalType(); nominalType != nil {
            terms = append(terms, nominalType)
        } else if structType := m.GetStructType(); structType != nil {
            terms = append(terms,
                &Term{StartPosition: structType.StartPosition,
                    EndPosition: structType.StartPosition,
                    Type:        TermTypeStart,
                })

            for _, field := range structType.Fields {
                terms = append(terms, field)
            }

            terms = append(terms,
                &Term{StartPosition: structType.EndPosition,
                    EndPosition: structType.EndPosition,
                    Type:        TermTypeEnd,
                })
        }

        if m.DefaultValue != nil {
            terms = append(terms, m.DefaultValue)
        }

        return terms
    }
    return nil
}
