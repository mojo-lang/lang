package lang

func (m *IntegerLiteralExpr) EvalValue() int64 {
	if m != nil {
		if m.IsNegative {
			return -m.Value
		}
		return m.Value
	} else {
		return 0
	}
}