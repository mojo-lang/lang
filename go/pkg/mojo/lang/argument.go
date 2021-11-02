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

func (m *Argument) IterateDictionary(iterator func(key *Expression, value *Expression) error) error {
	return m.GetValue().EvalDictionaryLiteral(iterator)
}

func (m *Argument) IterateStringDictionary(iterator func(key string, value *Expression) error) error {
	return m.GetValue().EvalStringDictionaryLiteral(iterator)
}
