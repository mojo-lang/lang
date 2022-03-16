package lang

import "errors"

const OriginalTypeAliasName = "original_type_alias_name"

func (m *TypeAliasDecl) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *TypeAliasDecl) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *TypeAliasDecl) MergeComment(comment *Comment) (bool, error) {
	if m != nil {
		return MergeCommentToTerms(comment, m.GetTerms())
	}

	return false, errors.New("nil NominalType")
}

func (m *TypeAliasDecl) GetTerms() []interface{} {
	if m != nil {
		var terms []interface{}

		for _, attribute := range m.Attributes {
			terms = append(terms, attribute)
		}

		terms = append(terms,
			NewSymbolTerm(m.KeywordPosition, TermTypeKeyword, KeywordType),
			NewSymbolTerm(m.NamePosition, TermTypeName, m.Name))

		for _, parameter := range m.GenericParameters {
			terms = append(terms, parameter)
		}

		terms = append(terms, m.Type)

		return terms
	}
	return nil
}

func (m *TypeAliasDecl) GetFullName() string {
	if m != nil {
		return GetFullName(m.PackageName, GetEnclosingNames(m.EnclosingType), m.Name)
	}
	return ""
}

func (m *TypeAliasDecl) SetScope(scope *Scope) {
	if m != nil {
		m.Scope = scope
	}
}

func (m *TypeAliasDecl) IsGeneric() bool {
	return m != nil && len(m.GenericParameters) > 0
}
