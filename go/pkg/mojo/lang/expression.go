package lang

import "errors"

func NewNullLiteralExpression(expr *NullLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_NullLiteralExpr{
			NullLiteralExpr: expr,
		},
	}
}

func NewBoolLiteralExpression(expr *BoolLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_BoolLiteralExpr{
			BoolLiteralExpr: expr,
		},
	}
}

func NewIntegerLiteralExpression(expr *IntegerLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_IntegerLiteralExpr{
			IntegerLiteralExpr: expr,
		},
	}
}

func NewFloatLiteralExpression(expr *FloatLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_FloatLiteralExpr{
			FloatLiteralExpr: expr,
		},
	}
}

func NewStringLiteralExpression(expr *StringLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_StringLiteralExpr{
			StringLiteralExpr: expr,
		},
	}
}

func NewArrayLiteralExpression(expr *ArrayLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_ArrayLiteralExpr{
			ArrayLiteralExpr: expr,
		},
	}
}

func NewMapLiteralExpression(expr *MapLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_MapLiteralExpr{
			MapLiteralExpr: expr,
		},
	}
}

func NewObjectLiteralExpression(expr *ObjectLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_ObjectLiteralExpr{
			ObjectLiteralExpr: expr,
		},
	}
}

func NewIdentifierExpression(expr *IdentifierExpr) *Expression {
	return &Expression{
		Expression: &Expression_IdentifierExpr{
			IdentifierExpr: expr,
		},
	}
}

func NewPrefixUnaryExpression(expr *PrefixUnaryExpr) *Expression {
	return &Expression{
		Expression: &Expression_PrefixUnaryExpr{
			PrefixUnaryExpr: expr,
		},
	}
}

func NewPostfixUnaryExpression(expr *PostfixUnaryExpr) *Expression {
	return &Expression{
		Expression: &Expression_PostfixUnaryExpr{
			PostfixUnaryExpr: expr,
		},
	}
}

func NewBinaryExpression(expr *BinaryExpr) *Expression {
	return &Expression{
		Expression: &Expression_BinaryExpr{
			BinaryExpr: expr,
		},
	}
}

func NewWildcardExpression(expr *WildcardExpr) *Expression {
	return &Expression{
		Expression: &Expression_WildcardExpr{
			WildcardExpr: expr,
		},
	}
}

func NewClosureExpression(expr *ClosureExpr) *Expression {
	return &Expression{
		Expression: &Expression_ClosureExpr{
			ClosureExpr: expr,
		},
	}
}

func NewTupleExpression(expr *TupleExpr) *Expression {
	return &Expression{
		Expression: &Expression_TupleExpr{
			TupleExpr: expr,
		},
	}
}

func (m *Expression) EvalBoolLiteral() (bool, error) {
	valueExpr := m.GetBoolLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Value, nil
	}
	return false, errors.New("expression can NOT be evaluated to the BoolLiteralExpr")
}

func (m *Expression) EvalIntegerLiteral() (int64, error) {
	valueExpr := m.GetIntegerLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalValue(), nil
	}
	return 0, errors.New("expression can NOT be evaluated to the Int64LiteralExpr")
}

func (m *Expression) EvalFloatLiteral() (float64, error) {
	valueExpr := m.GetFloatLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Value, nil
	}
	return 0, errors.New("expression can NOT be evaluated to the FloatLiteralExpr")
}

func (m *Expression) EvalStringLiteral() (string, error) {
	valueExpr := m.GetStringLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Value, nil
	}
	return "", errors.New("expression can NOT be evaluated to the StringLiteralExpr")
}

func (m *Expression) EvalArrayLiteral() ([]*Expression, error) {
	valueExpr := m.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Elements, nil
	}
	return nil, errors.New("expression can NOT be evaluated to the BoolLiteralExpr")
}

func (m *Expression) EvalIntegerArrayLiteral() ([]int64, error) {
	valueExpr := m.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalIntegerArray()
	}
	return nil, errors.New("expression can NOT be evaluated to the BoolLiteralExpr")
}

func (m *Expression) EvalFloatArrayLiteral() ([]float64, error) {
	valueExpr := m.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalFloatArray()
	}
	return nil, errors.New("expression can NOT be evaluated to the BoolLiteralExpr")
}

func (m *Expression) EvalStringArrayLiteral() ([]string, error) {
	valueExpr := m.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalStringArray()
	}
	return nil, errors.New("expression can NOT be evaluated to the BoolLiteralExpr")
}

func (m *Expression) EvalStringMapLiteral(iterator func(key string, value *Expression) error) error {
	valueExpr := m.GetMapLiteralExpr()
	if valueExpr != nil {
		for _, entry := range valueExpr.Entries {
			err := iterator(entry.Key, entry.Value)
			if err != nil {
				break
			}
		}
		return nil
	}
	return errors.New("expression can NOT be evaluated to the MapLiteralExpr")
}

func (m *Expression) EvalIdentifier() (string, error) {
	valueExpr := m.GetIdentifierExpr()
	if valueExpr != nil {
		return valueExpr.Identifier.Name, nil
	}
	return "", errors.New("expression can NOT be evaluated to the IdentifierExpr")
}
