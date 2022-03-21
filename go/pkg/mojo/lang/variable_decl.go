package lang

func (m *VariableDecl) SetStartPosition(position *Position) {
    if m != nil {
        m.StartPosition = PatchPosition(m.StartPosition, position)
    }
}

func (m *VariableDecl) SetEndPosition(position *Position) {
    if m != nil {
        m.EndPosition = PatchPosition(m.EndPosition, position)
    }
}

func (m *VariableDecl) GetInitialValue() *Expression {
    if m != nil {
        return m.GetInitializer().GetValue()
    }
    return nil
}
