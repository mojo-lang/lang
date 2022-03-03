
/// BinaryExpr - Infix binary expressions like 'x+y'.  The argument is always
/// an implicit tuple expression of the type expected by the function.
type BinaryExpr : ApplyExpr {
    operator: Operator @20

    left_hand_argument: Expression @21
    right_hand_argument: Expression @22
}