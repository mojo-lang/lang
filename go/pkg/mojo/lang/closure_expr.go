package lang

import "errors"

func (m *ClosureExpr) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *ClosureExpr) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *ClosureExpr) MergeComment(comment *Comment) (bool, error) {
	if m != nil {
		return MergeCommentToTerms(comment, m.GetTerms())
	}

	return false, errors.New("nil StructDecl")
}

func (m *ClosureExpr) GetTerms() []interface{} {
	if m != nil {
		var terms []interface{}

		if m.Signature != nil {
			if m.Signature.Parameter != nil {
				terms = append(terms, m.Signature.Parameter)
			}

			if m.Signature.Result != nil /*&& !m.Signature.Result.Implicit*/ {
				terms = append(terms, m.Signature.Result)
			}
		}

		if m.Body != nil {
			terms = append(terms, m.Body)
		}

		return terms
	}
	return nil
}
