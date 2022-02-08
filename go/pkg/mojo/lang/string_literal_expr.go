package lang

func NewStringLiteral(value string) *StringLiteralExpr {
	return &StringLiteralExpr{
		Kind:  0,
		Value: value,
	}
}

func NewImplicitStringLiteral(value string) *StringLiteralExpr {
	return &StringLiteralExpr{
		Kind:     0,
		Implicit: true,
		Value:    value,
	}
}

func (m *StringLiteralExpr) EvalValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}