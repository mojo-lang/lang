package lang

func (m *BoolLiteralExpr) EvalValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}
