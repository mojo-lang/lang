package lang

import (
	"errors"
	"github.com/mojo-lang/core/go/pkg/mojo/core"
	"reflect"
)

func NewNullLiteralExpression(expr *NullLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_NullLiteralExpr{NullLiteralExpr: expr},
	}
}

func NewBoolLiteralExpression(expr *BoolLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_BoolLiteralExpr{BoolLiteralExpr: expr},
	}
}

func NewIntegerLiteralExpressionFrom(value int64) *Expression {
	return NewIntegerLiteralExpression(NewIntegerLiteralExpr(value))
}

func NewIntegerLiteralExpression(expr *IntegerLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_IntegerLiteralExpr{IntegerLiteralExpr: expr},
	}
}

func NewFloatLiteralExpressionFrom(value float64) *Expression {
	return NewFloatLiteralExpression(NewFloatLiteralExpr(value))
}

func NewFloatLiteralExpression(expr *FloatLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_FloatLiteralExpr{FloatLiteralExpr: expr},
	}
}

func NewStringLiteralExpressionFrom(value string) *Expression {
	return NewStringLiteralExpression(NewStringLiteralExpr(value))
}

func NewStringLiteralExpression(expr *StringLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_StringLiteralExpr{StringLiteralExpr: expr},
	}
}

func NewArrayLiteralExpression(expr *ArrayLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_ArrayLiteralExpr{ArrayLiteralExpr: expr},
	}
}

func NewMapLiteralExpression(expr *MapLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_MapLiteralExpr{MapLiteralExpr: expr},
	}
}

func NewObjectLiteralExpression(expr *ObjectLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_ObjectLiteralExpr{ObjectLiteralExpr: expr},
	}
}

func NewStructLiteralExpression(expr *StructLiteralExpr) *Expression {
	return &Expression{
		Expression: &Expression_StructLiteralExpr{StructLiteralExpr: expr},
	}
}

func NewIdentifierExpression(expr *IdentifierExpr) *Expression {
	return &Expression{
		Expression: &Expression_IdentifierExpr{IdentifierExpr: expr},
	}
}

func NewStringLiteralUnaryExpression(expr *StringLiteralUnaryExpr) *Expression {
	return &Expression{
		Expression: &Expression_StringLiteralUnaryExpr{StringLiteralUnaryExpr: expr},
	}
}

func NewNumericLiteralUnaryExpression(expr *NumericLiteralUnaryExpr) *Expression {
	return &Expression{
		Expression: &Expression_NumericLiteralUnaryExpr{NumericLiteralUnaryExpr: expr},
	}
}

func NewPrefixUnaryExpression(expr *PrefixUnaryExpr) *Expression {
	return &Expression{
		Expression: &Expression_PrefixUnaryExpr{PrefixUnaryExpr: expr},
	}
}

func NewPostfixUnaryExpression(expr *PostfixUnaryExpr) *Expression {
	return &Expression{
		Expression: &Expression_PostfixUnaryExpr{PostfixUnaryExpr: expr},
	}
}

func NewFunctionCallExpression(expr *FunctionCallExpr) *Expression {
	return &Expression{
		Expression: &Expression_FunctionCallExpr{FunctionCallExpr: expr},
	}
}

func NewExplicitMemberExpression(expr *ExplicitMemberExpr) *Expression {
	return &Expression{
		Expression: &Expression_ExplicitMemberExpr{ExplicitMemberExpr: expr},
	}
}

func NewSubscriptExpression(expr *SubscriptExpr) *Expression {
	return &Expression{
		Expression: &Expression_SubscriptExpr{SubscriptExpr: expr},
	}
}

func NewConditionalExpression(expr *ConditionalExpr) *Expression {
	return &Expression{
		Expression: &Expression_ConditionalExpr{ConditionalExpr: expr},
	}
}

func NewAssignmentExpression(expr *AssignmentExpr) *Expression {
	return &Expression{
		Expression: &Expression_AssignmentExpr{AssignmentExpr: expr},
	}
}

func NewBinaryExpression(expr *BinaryExpr) *Expression {
	return &Expression{
		Expression: &Expression_BinaryExpr{BinaryExpr: expr},
	}
}

func NewWildcardExpression(expr *WildcardExpr) *Expression {
	return &Expression{
		Expression: &Expression_WildcardExpr{WildcardExpr: expr},
	}
}

func NewParenthesizedExpression(expr *ParenthesizedExpr) *Expression {
	return &Expression{
		Expression: &Expression_ParenthesizedExpr{ParenthesizedExpr: expr},
	}
}

func NewClosureExpression(expr *ClosureExpr) *Expression {
	return &Expression{
		Expression: &Expression_ClosureExpr{ClosureExpr: expr},
	}
}

func NewTupleExpression(expr *TupleExpr) *Expression {
	return &Expression{
		Expression: &Expression_TupleExpr{TupleExpr: expr},
	}
}

func (x *Expression) IsUnion() {
}

func (x *Expression) GetStartPosition() *Position {
	return GetUnionPosition(x, StartPositionFieldName)
}

func (x *Expression) GetEndPosition() *Position {
	return GetUnionPosition(x, EndPositionFieldName)
}

func (x *Expression) SetStartPosition(position *Position) {
	SetUnionPosition(x, StartPositionFieldName, position)
}

func (x *Expression) SetEndPosition(position *Position) {
	SetUnionPosition(x, EndPositionFieldName, position)
}

func (x *Expression) EvalBoolLiteral() (bool, error) {
	valueExpr := x.GetBoolLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Value, nil
	}
	return false, errors.New("expression can NOT be evaluated to the BoolLiteralExpr")
}

func (x *Expression) EvalIntegerLiteral() (int64, error) {
	valueExpr := x.GetIntegerLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalValue(), nil
	}
	return 0, errors.New("expression can NOT be evaluated to the Int64LiteralExpr")
}

func (x *Expression) EvalFloatLiteral() (float64, error) {
	valueExpr := x.GetFloatLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Value, nil
	}
	return 0, errors.New("expression can NOT be evaluated to the FloatLiteralExpr")
}

func (x *Expression) EvalStringLiteral() (string, error) {
	valueExpr := x.GetStringLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Value, nil
	}
	return "", errors.New("expression can NOT be evaluated to the StringLiteralExpr")
}

func (x *Expression) EvalArrayLiteral() ([]*Expression, error) {
	valueExpr := x.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.Elements, nil
	}
	return nil, errors.New("expression can NOT be evaluated to the ArrayLiteralExpr")
}

func (x *Expression) EvalIntegerArrayLiteral() ([]int64, error) {
	valueExpr := x.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalIntegerArray()
	}
	return nil, errors.New("expression can NOT be evaluated to the ArrayLiteralExpr")
}

func (x *Expression) EvalFloatArrayLiteral() ([]float64, error) {
	valueExpr := x.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalFloatArray()
	}
	return nil, errors.New("expression can NOT be evaluated to the ArrayLiteralExpr")
}

func (x *Expression) EvalStringArrayLiteral() ([]string, error) {
	valueExpr := x.GetArrayLiteralExpr()
	if valueExpr != nil {
		return valueExpr.EvalStringArray()
	}
	return nil, errors.New("expression can NOT be evaluated to the ArrayLiteralExpr")
}

func (x *Expression) EvalStringMapLiteral(iterator func(key string, value *Expression) error) error {
	valueExpr := x.GetMapLiteralExpr()
	if valueExpr != nil {
		for _, entry := range valueExpr.Entries {
			err := iterator(entry.Key, entry.Value)
			if err != nil {
				if core.IsBreakError(err) {
					return nil
				}
				return err
			}
		}
		return nil
	}
	return errors.New("expression can NOT be evaluated to the MapLiteralExpr")
}

func (x *Expression) EvalIdentifier() (string, error) {
	valueExpr := x.GetIdentifierExpr()
	if valueExpr != nil {
		return valueExpr.Identifier.Name, nil
	}
	return "", errors.New("expression can NOT be evaluated to the IdentifierExpr")
}

func (x *Expression) MergeComment(comment *Comment) (bool, error) {
	if x != nil {
		value := reflect.ValueOf(x.Expression)
		value = reflect.Indirect(value).Field(0)
		if merger, ok := value.Interface().(CommentMerger); ok {
			return merger.MergeComment(comment)
		} else {
			// error
		}
	}

	return false, errors.New("nil Expression")
}
