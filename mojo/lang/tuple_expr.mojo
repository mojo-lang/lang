/// TupleExpr - Parenthesized expressions like `(a: x+x)` and '(x, y, 4)'.  Also
/// used to represent the operands to a binary operator.  Note that
/// expressions like '(4)' are represented with a ParenExpr.
type TupleExpr : Expr {
    /// Whether this tuple has any labels.
    has_element_names : Bool

    /// Whether this tuple has label locations.
    has_element_locations : Bool
}