/// TypeAliasDecl - This is a declaration of a type alias, for example:
///
/// ```
///    type Foo = Int
/// ```
type TypeAliasDecl : TypeDecl @skip_fields(attributes) {
    ///
    type: NominalType @13
}