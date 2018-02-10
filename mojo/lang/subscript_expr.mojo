/// Subscripting expressions like a[i] that refer to an element within a
/// container.
///
/// There is no built-in subscripting in the language. Rather, a fully
/// type-checked and well-formed subscript expression refers to a subscript
/// declaration, which provides a getter and (optionally) a setter that will
/// be used to perform reads/writes.
class SubscriptExpr final : public Expr
