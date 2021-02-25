///
///
type Decl {
    /// position of first character belonging to the Expr
    start_position: Position @1

    /// position of first character immediately after the Expr
    end_position:   Position @2

    ///
    document : Cached<Document> @4

    ///
    package: Package @5 @reference('full_name')
}