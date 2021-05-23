package lang

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

func (i IntegerLiteralExpr) EvalValue() int64 {
	if i.IsNegative {
		return -i.Value
	}
	return i.Value
}
