/// \brief Integer literal with a '+' or '-' sign, like '+4' or '- 2'.
///
/// After semantic analysis assigns types, this is guaranteed to only have
/// a BuiltinIntegerType.
type IntegerLiteralExpr : NumericLiteralExpr {
    value : Int64 @16
}