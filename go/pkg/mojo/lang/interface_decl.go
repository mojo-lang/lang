package lang

import "errors"

func (m *InterfaceDecl) SetStartPosition(position *Position) {
    if m != nil {
        m.StartPosition = PatchPosition(m.StartPosition, position)
    }
}

func (m *InterfaceDecl) SetEndPosition(position *Position) {
    if m != nil {
        m.EndPosition = PatchPosition(m.EndPosition, position)
    }
}

func (m *InterfaceDecl) MergeComment(comment *Comment) (bool, error) {
    if m != nil {
        return MergeCommentToTerms(comment, m.GetTerms())
    }

    return false, errors.New("nil InterfaceDecl")
}

func (m *InterfaceDecl) GetTerms() []interface{} {
    if m != nil {
        var terms []interface{}

        if m.Document != nil {
            terms = append(terms, m.Document)
        }
        for _, attribute := range m.Attributes {
            terms = append(terms, attribute)
        }

        terms = append(terms,
            NewSymbolTerm(m.KeywordPosition, TermTypeKeyword, KeywordInterface),
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

            for _, method := range m.Type.Methods {
                terms = append(terms, method)
            }

            terms = append(terms, NewSymbolTerm(m.Type.EndPosition, TermTypeEnd, ""))
        }

        return terms
    }
    return nil
}

func (m *InterfaceDecl) GetFullName() string {
    if m != nil {
        return GetFullName(m.PackageName, nil, m.Name)
    }
    return ""
}

func (m *InterfaceDecl) SetScope(scope *Scope) {
    if m != nil {
        m.Scope = scope
    }
}

func (m *InterfaceDecl) GetInheritMethods() []*FunctionDecl {
    if m != nil && m.Type != nil {
        return m.Type.GetInheritMethods()
    }
    return nil
}

func (m *InterfaceDecl) GetMethodGroups() map[string][]*FunctionDecl {
    if m != nil && m.Type != nil {
        return m.Type.GetMethodGroups()
    }
    return nil
}

func (m *InterfaceDecl) IsBodyEmpty() bool {
    if m != nil {
        return len(m.TypeAliasDecls) == 0 && m.MethodCount() == 0
    }
    return true
}

func (m *InterfaceDecl) MethodCount() int {
    return len(m.GetType().GetMethods())
}
