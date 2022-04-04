package lang

func (x *VariableDecl) SetStartPosition(position *Position) {
    if x != nil {
        x.StartPosition = PatchPosition(x.StartPosition, position)
    }
}

func (x *VariableDecl) SetEndPosition(position *Position) {
    if x != nil {
        x.EndPosition = PatchPosition(x.EndPosition, position)
    }
}

func (x *VariableDecl) GetInitialValue() *Expression {
    if x != nil {
        return x.GetInitializer().GetValue()
    }
    return nil
}
