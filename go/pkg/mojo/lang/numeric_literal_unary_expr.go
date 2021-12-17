package lang

func (m *NumericLiteralUnaryExpr) GetIntegerLiteralExpr() *IntegerLiteralExpr {
	return m.GetExpression().GetIntegerLiteralExpr()
}

func (m *NumericLiteralUnaryExpr) GetFloatLiteralExpr() *FloatLiteralExpr {
	return m.GetExpression().GetFloatLiteralExpr()
}
