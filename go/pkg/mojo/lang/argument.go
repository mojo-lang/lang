package lang

import "errors"

func NewBoolLiteralArgument(expr *BoolLiteralExpr) *Argument {
    return &Argument{
        Value: NewBoolLiteralExpression(expr),
    }
}

func NewIntegerLiteralArgument(expr *IntegerLiteralExpr) *Argument {
    return &Argument{
        Value: NewIntegerLiteralExpression(expr),
    }
}

func NewFloatLiteralArgument(expr *FloatLiteralExpr) *Argument {
    return &Argument{
        Value: NewFloatLiteralExpression(expr),
    }
}

func NewStringLiteralArgument(expr *StringLiteralExpr) *Argument {
    return &Argument{
        Value: NewStringLiteralExpression(expr),
    }
}

func NewStringArgument(val string) *Argument {
    return NewStringLiteralArgument(NewStringLiteral(val))
}

func NewArrayLiteralArgument(expr *ArrayLiteralExpr) *Argument {
    return &Argument{
        Value: NewArrayLiteralExpression(expr),
    }
}

func NewMapLiteralArgument(expr *MapLiteralExpr) *Argument {
    return &Argument{
        Value: NewMapLiteralExpression(expr),
    }
}

func NewObjectLiteralArgument(expr *ObjectLiteralExpr) *Argument {
    return &Argument{
        Value: NewObjectLiteralExpression(expr),
    }
}

func (m *Argument) SetStartPosition(position *Position) {
    if m != nil {
        m.StartPosition = PatchPosition(m.StartPosition, position)
    }
}

func (m *Argument) SetEndPosition(position *Position) {
    if m != nil {
        m.EndPosition = PatchPosition(m.EndPosition, position)
    }
}

func (m *Argument) MergeComment(comment *Comment) (bool, error) {
    if m != nil {
        return MergeCommentToTerms(comment, m.GetTerms())
    }

    return false, errors.New("nil Argument")
}

func (m *Argument) GetTerms() []interface{} {
    if m != nil {
        var terms []interface{}

        if len(m.Label) > 0 {
            terms = append(terms, &Term{
                StartPosition: m.StartPosition,
                EndPosition: &Position{
                    Line:   m.StartPosition.Line,
                    Column: m.StartPosition.Column + int64(len(m.Label)),
                },
                Type:  "Label",
                Value: m.Label,
            })
        }
        return append(terms, m.Value)
    }
    return nil
}

func (m *Argument) GetNullLiteralExpr() *NullLiteralExpr {
    if m != nil && m.Value != nil {
        return m.Value.GetNullLiteralExpr()
    }
    return nil
}

func (m *Argument) GetIntegerLiteralExpr() *IntegerLiteralExpr {
    if m != nil && m.Value != nil {
        return m.Value.GetIntegerLiteralExpr()
    }
    return nil
}

func (m *Argument) GetFloatLiteralExpr() *FloatLiteralExpr {
    if m != nil && m.Value != nil {
        return m.Value.GetFloatLiteralExpr()
    }
    return nil
}

func (m *Argument) GetBoolLiteralExpr() *BoolLiteralExpr {
    if m != nil && m.Value != nil {
        return m.Value.GetBoolLiteralExpr()
    }
    return nil
}

func (m *Argument) GetStringLiteralExpr() *StringLiteralExpr {
    if m != nil && m.Value != nil {
        return m.Value.GetStringLiteralExpr()
    }
    return nil
}

func (m *Argument) GetObjectLiteralExpr() *ObjectLiteralExpr {
    if m != nil && m.Value != nil {
        return m.Value.GetObjectLiteralExpr()
    }
    return nil
}

func (m *Argument) GetArrayLiteralExpr() *ArrayLiteralExpr {
    if m != nil && m.Value != nil {
        return m.Value.GetArrayLiteralExpr()
    }
    return nil
}

func (m *Argument) GetMapLiteralExpr() *MapLiteralExpr {
    if m != nil && m.Value != nil {
        return m.Value.GetMapLiteralExpr()
    }
    return nil
}

func (m *Argument) GetIdentifierExpr() *IdentifierExpr {
    if m != nil && m.Value != nil {
        return m.Value.GetIdentifierExpr()
    }
    return nil
}

func (m *Argument) GetBool() (bool, error) {
    return m.GetValue().EvalBoolLiteral()
}

func (m *Argument) GetInteger() (int64, error) {
    return m.GetValue().EvalIntegerLiteral()
}

func (m *Argument) GetFloat() (float64, error) {
    return m.GetValue().EvalFloatLiteral()
}

func (m *Argument) GetString() (string, error) {
    return m.GetValue().EvalStringLiteral()
}

func (m *Argument) GetArray() ([]*Expression, error) {
    return m.GetValue().EvalArrayLiteral()
}

func (m *Argument) GetIntegerArray() ([]int64, error) {
    return m.GetValue().EvalIntegerArrayLiteral()
}

func (m *Argument) GetFloatArray() ([]float64, error) {
    return m.GetValue().EvalFloatArrayLiteral()
}

func (m *Argument) GetStringArray() ([]string, error) {
    return m.GetValue().EvalStringArrayLiteral()
}

func (m *Argument) IterateStringMap(iterator func(key string, value *Expression) error) error {
    return m.GetValue().EvalStringMapLiteral(iterator)
}
