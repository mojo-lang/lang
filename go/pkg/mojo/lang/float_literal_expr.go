package lang

func NewFloatLiteral(value float64) *FloatLiteralExpr {
	return &FloatLiteralExpr{
		Kind:  0,
		Value: value,
	}
}

func (m *FloatLiteralExpr) SetNegative() *FloatLiteralExpr {
	if m != nil {
		m.IsNegative = true
	}
	return m
}
