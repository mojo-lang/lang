package lang

func (m *MapLiteralExpr) SetStartPosition(position *Position) {
    if m != nil {
        m.StartPosition = PatchPosition(m.StartPosition, position)
    }
}

func (m *MapLiteralExpr) SetEndPosition(position *Position) {
    if m != nil {
        m.EndPosition = PatchPosition(m.EndPosition, position)
    }
}

func (m *MapLiteralExpr) RangeStringKeys(iter func(key string, value *Expression) error) {
}
