package lang

func NewBoolLiteralArgument(expr *BoolLiteralExpr) *Argument {
	return &Argument{
		Value: NewBoolLiteralExpr(expr),
	}
}

func NewIntegerLiteralArgument(expr *IntegerLiteralExpr) *Argument {
	return &Argument{
		Value: NewIntegerLiteralExpr(expr),
	}
}

func NewFloatLiteralArgument(expr *FloatLiteralExpr) *Argument {
	return &Argument{
		Value: NewFloatLiteralExpr(expr),
	}
}

func NewStringLiteralArgument(expr *StringLiteralExpr) *Argument {
	return &Argument{
		Value: NewStringLiteralExpr(expr),
	}
}

func NewArrayLiteralArgument(expr *ArrayLiteralExpr) *Argument {
	return &Argument{
		Value: NewArrayLiteralExpr(expr),
	}
}

func NewDictionaryLiteralArgument(expr *DictionaryLiteralExpr) *Argument {
	return &Argument{
		Value: NewDictionaryLiteralExpr(expr),
	}
}

func NewObjectLiteralArgument(expr *ObjectLiteralExpr) *Argument {
	return &Argument{
		Value: NewObjectLiteralExpr(expr),
	}
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

func (m *Argument) GetString() *string {
	if m != nil && m.Value != nil {
		if expr := m.Value.GetStringLiteralExpr(); expr != nil {
			return &expr.Value
		}
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

func (m *Argument) GetDictionaryLiteralExpr() *DictionaryLiteralExpr {
	if m != nil && m.Value != nil {
		return m.Value.GetDictionaryLiteralExpr()
	}
	return nil
}

func (m *Argument) GetIdentifierExpr() *IdentifierExpr {
	if m != nil && m.Value != nil {
		return m.Value.GetIdentifierExpr()
	}
	return nil
}
