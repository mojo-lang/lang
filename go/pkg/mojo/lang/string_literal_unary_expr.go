package lang

func (x *StringLiteralUnaryExpr) SetStartPosition(position *Position) {
    if x != nil {
        x.StartPosition = PatchPosition(x.StartPosition, position)
    }
}

func (x *StringLiteralUnaryExpr) SetEndPosition(position *Position) {
    if x != nil {
        x.EndPosition = PatchPosition(x.EndPosition, position)
    }
}

func (x *StringLiteralUnaryExpr) GetStringLiteralExpr() *StringLiteralExpr {
    return x.GetArgument().GetStringLiteralExpr()
}
