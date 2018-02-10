/// \brief A collection literal expression.
///
/// The subexpression is represented as a TupleExpr or ParenExpr and
/// passed on to the appropriate semantics-providing conversion
/// operation.
class CollectionExpr : public Expr
