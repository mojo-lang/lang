/// \brief Represents a runtime type check query, 'a is T', where 'T' is a type
/// and 'a' is a value of some related type. Evaluates to a Bool true if 'a' is
/// of the type and 'a as T' would succeed, false otherwise.
///
/// FIXME: We should support type queries with a runtime metatype value too.
type IsExpr : CheckedCastExpr 


type CheckedCastExpr : Expr