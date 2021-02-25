/// TypeAliasDecl - This is a declaration of a type alias, for example:
///
/// ```
///    type Foo = Int
/// ```
type TypeAliasDecl : TypeDecl {
    ///
    type: NominalType @15
}