///
///
type Decl {
    /// position of first character belonging to the Expr
    start_pos: Pos @1

    /// position of first character immediately after the Expr
    end_pos  : Pos @2

    ///
    document : String @3 | Document @4

    ///
    package: Package @5 @reference
}