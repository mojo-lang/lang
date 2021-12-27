package lang

func (m *StringLiteralUnaryExpr) GetStringLiteralExpr() *StringLiteralExpr {
	return m.GetArgument().GetStringLiteralExpr()
}
