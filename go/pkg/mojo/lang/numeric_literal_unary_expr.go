package lang

func (x *NumericLiteralUnaryExpr) SetStartPosition(position *Position) {
    if x != nil {
        x.StartPosition = PatchPosition(x.StartPosition, position)
    }
}

func (x *NumericLiteralUnaryExpr) SetEndPosition(position *Position) {
    if x != nil {
        x.EndPosition = PatchPosition(x.EndPosition, position)
    }
}

func (x *NumericLiteralUnaryExpr) GetIntegerLiteralExpr() *IntegerLiteralExpr {
    return x.GetArgument().GetIntegerLiteralExpr()
}

func (x *NumericLiteralUnaryExpr) GetFloatLiteralExpr() *FloatLiteralExpr {
    return x.GetArgument().GetFloatLiteralExpr()
}
