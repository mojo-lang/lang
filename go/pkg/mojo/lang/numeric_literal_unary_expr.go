package lang

func (m *NumericLiteralUnaryExpr) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *NumericLiteralUnaryExpr) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *NumericLiteralUnaryExpr) GetIntegerLiteralExpr() *IntegerLiteralExpr {
	return m.GetArgument().GetIntegerLiteralExpr()
}

func (m *NumericLiteralUnaryExpr) GetFloatLiteralExpr() *FloatLiteralExpr {
	return m.GetArgument().GetFloatLiteralExpr()
}
