
/// BinaryExpr - Infix binary expressions like 'x+y'.  The argument is always
/// an implicit tuple expression of the type expected by the function.
type BinaryExpr : ApplyExpr {
    operator: Operator @15

    left_hand_argument: Expression @16
    right_hand_argument: Expression @17
}