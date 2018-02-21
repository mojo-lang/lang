///
///
type DeclGroup {
    /// position of first character belonging to the Group
    start_pos: Pos @1

    /// position of first character immediately after the Group
    end_pos  : Pos @2

    ///
    document: Document

    ///
    attributes: [Attribute]
}