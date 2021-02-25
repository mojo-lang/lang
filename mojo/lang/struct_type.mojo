///
///
///
type StructType {
    /// position of first character belonging to the Expr
    start_position: Position @1

    /// position of first character immediately after the Expr
    end_position: Position @2

    /// A Field represents a Field declaration list in a struct type,
    /// a method list in an interface type, or a parameter/result declaration
    /// in a signature.
    ///
    fields: [ValueDecl] @10

    ///
    inherits: [NominalType] @11

    ///
    groups: [ValueGroupDecl] @12
}
