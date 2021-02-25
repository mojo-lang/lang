/// \brief An explicit unnamed function expression, which can optionally have
/// named arguments.
///
/// ```
///     { $0 + $1 }
///     { a, b -> Int in a + b }
///     { (a : Int, b : Int) -> a + b }
///     { a : Int -> a + c.getFoo() }
/// ```
type LambdaExpr : Expr {
}
