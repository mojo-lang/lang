/// FloatLiteralExpr - Floating point literal, like '4.0'.  After semantic
/// analysis assigns types, this is guaranteed to only have a
/// BuiltinFloatingPointType.
type FloatLiteralExpr : NumberLiteralExpr {
    value: Float64
}