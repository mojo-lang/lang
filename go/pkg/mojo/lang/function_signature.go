package lang

import "errors"

func (m *FunctionSignature) GetParameters() []*ValueDecl {
    return m.GetParameter().GetDecls()
}

func (m *FunctionSignature) ParameterDecl(index int) *ValueDecl {
    if decls := m.GetParameters(); len(decls) > 0 && index >= 0 && index < len(decls) {
        return decls[index]
    }
    return nil
}

func (m *FunctionSignature) ParameterDeclByName(name string) *ValueDecl {
    if decls := m.GetParameters(); len(decls) > 0 && len(name) > 0 {
        for _, decl := range decls {
            if decl.Name == name {
                return decl
            }
        }
    }
    return nil
}

func (m *FunctionSignature) GetResultType() *NominalType {
    return m.GetResult().GetType()
}

func (m *FunctionSignature) SetStartPosition(position *Position) {
    if m != nil {
        m.StartPosition = PatchPosition(m.StartPosition, position)
    }
}

func (m *FunctionSignature) SetEndPosition(position *Position) {
    if m != nil {
        m.EndPosition = PatchPosition(m.EndPosition, position)
    }
}

func (m *FunctionSignature) MergeComment(comment *Comment) (bool, error) {
    if m != nil {
        return MergeCommentToTerms(comment, m.GetTerms())
    }

    return false, errors.New("nil FunctionSignature")
}

func (m *FunctionSignature) GetTerms() []interface{} {
    if m != nil {
        var terms []interface{}

        if m.Parameter != nil {
            includingParen := m.GetStartPosition().Compare(m.Parameter.GetStartPosition()) == 0
            if includingParen {
                terms = append(terms,
                    &Term{StartPosition: m.StartPosition,
                        EndPosition: m.StartPosition,
                        Type:        TermTypeStart,
                        Value:       "(",
                    })
            }

            for _, parameter := range m.GetParameters() {
                terms = append(terms, parameter)
            }

            if includingParen {
                terms = append(terms,
                    &Term{StartPosition: m.StartPosition,
                        EndPosition: m.StartPosition,
                        Type:        TermTypeEnd,
                        Value:       ")",
                    })
            }
        }

        if m.Result != nil && !m.Result.Implicit {
            terms = append(terms,
                &Term{StartPosition: m.Result.GetStartPosition(),
                    EndPosition: m.Result.GetStartPosition(),
                    Type:        TermTypeSymbol,
                    Value:       "->",
                })

            terms = append(terms, m.GetResultType())
        }

        return terms
    }

    return nil
}

func (m *FunctionSignature_Parameter) HasFollowingDocument() bool {
    for _, param := range m.Decls {
        if param.Document.Following {
            return true
        }
    }
    return false
}

func (m *FunctionSignature_Result) HasFollowingDocument() bool {
    return m != nil && m.Type != nil && m.Type.Document.Following
}
