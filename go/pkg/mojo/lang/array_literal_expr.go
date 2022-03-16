package lang

import "errors"

func (m *ArrayLiteralExpr) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *ArrayLiteralExpr) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *ArrayLiteralExpr) MergeComment(comment *Comment) (bool, error) {
	if m != nil {
		return MergeCommentToTerms(comment, m.GetTerms())
	}

	return false, errors.New("nil ArrayLiteralExpr")
}

func (m *ArrayLiteralExpr) GetTerms() []interface{} {
	if m != nil {
		var terms []interface{}
		for _, element := range m.Elements {
			terms = append(terms, element)
		}

		return terms
	}
	return nil
}

func (m *ArrayLiteralExpr) EvalIntegerArray() ([]int64, error) {
	if m != nil {
		var values []int64
		for _, elem := range m.Elements {
			value, err := elem.EvalIntegerLiteral()
			if err != nil {
				return nil, err
			}
			values = append(values, value)
		}
		return values, nil
	}
	return nil, errors.New("ArrayLiteralExpr is nil")
}

func (m *ArrayLiteralExpr) EvalFloatArray() ([]float64, error) {
	if m != nil {
		var values []float64
		for _, elem := range m.Elements {
			value, err := elem.EvalFloatLiteral()
			if err != nil {
				return nil, err
			}
			values = append(values, value)
		}
		return values, nil
	}
	return nil, errors.New("ArrayLiteralExpr is nil")
}

func (m *ArrayLiteralExpr) EvalStringArray() ([]string, error) {
	if m != nil {
		var values []string
		for _, elem := range m.Elements {
			value, err := elem.EvalStringLiteral()
			if err != nil {
				return nil, err
			}
			values = append(values, value)
		}
		return values, nil
	}
	return nil, errors.New("ArrayLiteralExpr is nil")
}
