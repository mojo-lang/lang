///
///
///
type StructType {
    ///
    package: String @4 | Package @5 @reference

    /// A Field represents a Field declaration list in a struct type,
    /// a method list in an interface type, or a parameter/result declaration
    /// in a signature.
    ///
    fields: [ValueDecl] @10

    ///
    inherits: [NominalType] @11
}
