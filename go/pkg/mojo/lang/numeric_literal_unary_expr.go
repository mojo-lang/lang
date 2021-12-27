package lang

func (m *NumericLiteralUnaryExpr) GetIntegerLiteralExpr() *IntegerLiteralExpr {
	return m.GetArgument().GetIntegerLiteralExpr()
}

func (m *NumericLiteralUnaryExpr) GetFloatLiteralExpr() *FloatLiteralExpr {
	return m.GetArgument().GetFloatLiteralExpr()
}
