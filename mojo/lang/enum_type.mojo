///
///
///
type EnumType {
    /// position of first character belonging to the Expr
    start_position: Position @1

    /// position of first character immediately after the Expr
    end_position: Position @2

    //
    enumerators: [ValueDecl] @10

    //
    underlying_type: NominalType @11
}
