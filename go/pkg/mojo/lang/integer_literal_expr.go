package lang

func NewIntegerLiteral(value int64) *IntegerLiteralExpr {
    isNegative := value < 0
    if isNegative {
        value = -value
    }
    return &IntegerLiteralExpr{
        Kind:       0,
        IsNegative: isNegative,
        Value:      uint64(value),
    }
}

func (m *IntegerLiteralExpr) SetStartPosition(position *Position) {
    if m != nil {
        m.StartPosition = PatchPosition(m.StartPosition, position)
    }
}

func (m *IntegerLiteralExpr) SetEndPosition(position *Position) {
    if m != nil {
        m.EndPosition = PatchPosition(m.EndPosition, position)
    }
}

func (m *IntegerLiteralExpr) EvalValue() int64 {
    if m != nil {
        if m.IsNegative {
            return -int64(m.Value)
        }
        return int64(m.Value)
    } else {
        return 0
    }
}

func (m *IntegerLiteralExpr) SetNegative() *IntegerLiteralExpr {
    if m != nil {
        m.IsNegative = true
    }
    return m
}
