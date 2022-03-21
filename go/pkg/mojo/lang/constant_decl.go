package lang

func (m *ConstantDecl) SetStartPosition(position *Position) {
    if m != nil {
        m.StartPosition = PatchPosition(m.StartPosition, position)
    }
}

func (m *ConstantDecl) SetEndPosition(position *Position) {
    if m != nil {
        m.EndPosition = PatchPosition(m.EndPosition, position)
    }
}

func (m *ConstantDecl) GetInitialValue() *Expression {
    if m != nil {
        return m.GetInitializer().GetValue()
    }
    return nil
}
