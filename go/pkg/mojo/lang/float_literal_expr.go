package lang

func NewFloatLiteral(value float64) *FloatLiteralExpr {
	return &FloatLiteralExpr{
		Kind:  0,
		Value: value,
	}
}

func (m *FloatLiteralExpr) EvalValue() float64 {
	if m != nil {
		if m.IsNegative {
			return -m.Value
		}
		return m.Value
	}
	return 0
}

func (m *FloatLiteralExpr) SetNegative() *FloatLiteralExpr {
	if m != nil {
		m.IsNegative = true
	}
	return m
}