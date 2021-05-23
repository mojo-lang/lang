/// InterpolatedStringLiteral - An interpolated string literal.
///
/// An interpolated string literal mixes expressions (which are evaluated and
/// converted into string form) within a string literal.
///
/// \code
/// "[\(min)..\(max)]"
/// \endcode
type InterpolatedStringLiteralExpr : LiteralExpr {
    segments: [Expr] @20
    semantic_expr: Expr @21
}