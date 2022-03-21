package lang

func (m *BoolLiteralExpr) SetStartPosition(position *Position) {
    if m != nil {
        m.StartPosition = PatchPosition(m.StartPosition, position)
    }
}

func (m *BoolLiteralExpr) SetEndPosition(position *Position) {
    if m != nil {
        m.EndPosition = PatchPosition(m.EndPosition, position)
    }
}

func (m *BoolLiteralExpr) EvalValue() bool {
    if m != nil {
        return m.Value
    }
    return false
}
