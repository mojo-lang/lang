package lang

func NewStringLiteral(value string) *StringLiteralExpr {
    return &StringLiteralExpr{
        Kind:  0,
        Value: value,
    }
}

func NewImplicitStringLiteral(value string) *StringLiteralExpr {
    return &StringLiteralExpr{
        Kind:     0,
        Implicit: true,
        Value:    value,
    }
}

func (m *StringLiteralExpr) SetStartPosition(position *Position) {
    if m != nil {
        m.StartPosition = PatchPosition(m.StartPosition, position)
    }
}

func (m *StringLiteralExpr) SetEndPosition(position *Position) {
    if m != nil {
        m.EndPosition = PatchPosition(m.EndPosition, position)
    }
}

func (m *StringLiteralExpr) EvalValue() string {
    if m != nil {
        return m.Value
    }
    return ""
}
