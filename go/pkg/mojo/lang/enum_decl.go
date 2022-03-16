package lang

import "errors"

func (m *EnumDecl) GetFullName() string {
	if m != nil {
		return GetFullName(m.PackageName, GetEnclosingNames(m.EnclosingType), m.Name)
	}
	return ""
}

func (m *EnumDecl) SetScope(scope *Scope) {
	if m != nil {
		m.Scope = scope
	}
}

func (m *EnumDecl) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *EnumDecl) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *EnumDecl) MergeComment(comment *Comment) (bool, error) {
	if m != nil {
		return MergeCommentToTerms(comment, m.GetTerms())
	}

	return false, errors.New("nil EnumDecl")
}

func (m *EnumDecl) GetTerms() []interface{} {
	if m != nil {
		var terms []interface{}

		for _, attribute := range m.Attributes {
			terms = append(terms, attribute)
		}

		terms = append(terms,
			NewSymbolTerm(m.KeywordPosition, TermTypeKeyword, KeywordEnum),
			NewSymbolTerm(m.NamePosition, TermTypeName, m.Name))

		for _, parameter := range m.GenericParameters {
			terms = append(terms, parameter)
		}

		if m.Type != nil {
			if m.Type.UnderlyingType != nil {
				terms = append(terms,
					NewSymbolTerm(m.Type.UnderlyingTypePosition, TermTypeSymbol, ":"),
					m.Type.UnderlyingType)
			}

			terms = append(terms, NewSymbolTerm(m.Type.StartPosition, TermTypeStart, ""))

			for _, enumerator := range m.Type.Enumerators {
				terms = append(terms, enumerator)
			}

			terms = append(terms, NewSymbolTerm(m.Type.EndPosition, TermTypeEnd, ""))
		}
	}

	return nil
}
