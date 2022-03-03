package lang

func (m *StringLiteralUnaryExpr) SetStartPosition(position *Position) {
	if m != nil {
		m.StartPosition = PatchPosition(m.StartPosition, position)
	}
}

func (m *StringLiteralUnaryExpr) SetEndPosition(position *Position) {
	if m != nil {
		m.EndPosition = PatchPosition(m.EndPosition, position)
	}
}

func (m *StringLiteralUnaryExpr) GetStringLiteralExpr() *StringLiteralExpr {
	return m.GetArgument().GetStringLiteralExpr()
}
