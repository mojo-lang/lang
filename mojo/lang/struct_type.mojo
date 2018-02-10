///
///
///
type StructType : DataType {
    // A Field represents a Field declaration list in a struct type,
    // a method list in an interface type, or a parameter/result declaration
    // in a signature.
    //
    fields: [ValueDecl] @10
}
