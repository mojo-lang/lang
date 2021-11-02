package lang

import "errors"

func NewNullLiteralExpr(expr *NullLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_NullLiteralExpr{
			NullLiteralExpr: expr,
		},
	}
}

func NewBoolLiteralExpr(expr *BoolLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_BoolLiteralExpr{
			BoolLiteralExpr: expr,
		},
	}
}

func NewIntegerLiteralExpr(expr *IntegerLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_IntegerLiteralExpr{
			IntegerLiteralExpr: expr,
		},
	}
}

func NewFloatLiteralExpr(expr *FloatLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_FloatLiteralExpr{
			FloatLiteralExpr: expr,
		},
	}
}

func NewStringLiteralExpr(expr *StringLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_StringLiteralExpr{
			StringLiteralExpr: expr,
		},
	}
}

func NewArrayLiteralExpr(expr *ArrayLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_ArrayLiteralExpr{
			ArrayLiteralExpr: expr,
		},
	}
}

func NewDictionaryLiteralExpr(expr *DictionaryLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_DictionaryLiteralExpr{
			DictionaryLiteralExpr: expr,
		},
	}
}

func NewObjectLiteralExpr(expr *ObjectLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_ObjectLiteralExpr{
			ObjectLiteralExpr: expr,
		},
	}
}

func NewIdentifierExpr(expr *IdentifierExpr) *Expression {
	return &Expression{
		Expression: &Expression_IdentifierExpr{
			IdentifierExpr: expr,
		},
	}
}

func (m *Expression) EvalBoolLiteral() (bool, error) {
	valueExpr := m.GetBoolLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Value, nil
	}
	return false, errors.New("Expression can NOT be evaluated to the BoolLiteralExpr")
}

func (m *Expression) EvalIntegerLiteral() (int64, error) {
	valueExpr := m.GetIntegerLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalValue(), nil
	}
	return 0, errors.New("Expression can NOT be evaluated to the Int64LiteralExpr")
}

func (m *Expression) EvalFloatLiteral() (float64, error) {
	valueExpr := m.GetFloatLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Value, nil
	}
	return 0, errors.New("Expression can NOT be evaluated to the FloatLiteralExpr")
}

func (m *Expression) EvalStringLiteral() (string, error) {
	valueExpr := m.GetStringLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Value, nil
	}
	return "", errors.New("Expression can NOT be evaluated to the StringLiteralExpr")
}

func (m *Expression) EvalArrayLiteral() ([]*Expression, error) {
	valueExpr := m.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Elements, nil
	}
	return nil, errors.New("Expression can NOT be evaluated to the BoolLiteralExpr")
}

func (m *Expression) EvalIntegerArrayLiteral() ([]int64, error) {
	valueExpr := m.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalIntegerArray()
	}
	return nil, errors.New("Expression can NOT be evaluated to the BoolLiteralExpr")
}

func (m *Expression) EvalFloatArrayLiteral() ([]float64, error) {
	valueExpr := m.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalFloatArray()
	}
	return nil, errors.New("Expression can NOT be evaluated to the BoolLiteralExpr")
}

func (m *Expression) EvalStringArrayLiteral() ([]string, error) {
	valueExpr := m.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalStringArray()
	}
	return nil, errors.New("Expression can NOT be evaluated to the BoolLiteralExpr")
}

func (m *Expression) EvalDictionaryLiteral(iterator func(key *Expression, value *Expression) error) error {
	valueExpr := m.GetDictionaryLiteralExpr()
	if valueExpr != nil {
		for _, entry := range valueExpr.Entries {
			err := iterator(entry.Key, entry.Value)
			if err != nil {
				break
			}
		}
		return nil
	}
	return errors.New("Expression can NOT be evaluated to the DictionaryLiteralExpr")
}

func (m *Expression) EvalStringDictionaryLiteral(iterator func(key string, value *Expression) error) error {
	valueExpr := m.GetDictionaryLiteralExpr()
	if valueExpr != nil {
		for _, entry := range valueExpr.Entries {
			key, err := entry.Key.EvalStringLiteral()
			if err != nil {
				return err
			}

			err = iterator(key, entry.Value)
			if err != nil {
				break
			}
		}
		return nil
	}
	return errors.New("Expression can NOT be evaluated to the DictionaryLiteralExpr")
}

func (m *Expression) EvalIdentifier() (string, error) {
	valueExpr := m.GetIdentifierExpr()
	if valueExpr != nil {
		return valueExpr.Identifier.Name, nil
	}
	return "", errors.New("Expression can NOT be evaluated to the IdentifierExpr")
}
