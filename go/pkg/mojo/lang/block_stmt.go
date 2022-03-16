package lang

import "errors"

func (m *BlockStmt) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *BlockStmt) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *BlockStmt) MergeComment(comment *Comment) (bool, error) {
	if m != nil {
		return MergeCommentToTerms(comment, m.GetTerms())
	}

	return false, errors.New("nil BlockStmt")
}

func (m *BlockStmt) GetTerms() []interface{} {
	if m != nil {
		var terms []interface{}

		terms = append(terms, NewSymbolTerm(m.StartPosition, TermTypeStart, ""))

		for _, statement := range m.Statements {
			terms = append(terms, statement)
		}

		terms = append(terms, NewSymbolTerm(m.EndPosition, TermTypeEnd, ""))
		return terms
	}
	return nil
}
