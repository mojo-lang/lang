/// \brief An explicit unnamed function expression, which can optionally have
/// named arguments.
///
/// \code
///     { $0 + $1 }
///     { a, b -> Int in a + b }
///     { (a : Int, b : Int) -> Int in a + b }
///     { [weak c] (a : Int) -> Int in a + c!.getFoo() }
/// \endcode
type ClosureExpr