/// TypeAliasDecl - This is a declaration of a type alias, for example:
///
/// ```
///    type Foo = Int
/// ```
type TypeAliasDecl : TypeDecl {
    ///
    type: DataType @15 | NominalType @16
}