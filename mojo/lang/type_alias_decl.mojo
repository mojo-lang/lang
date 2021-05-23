/// TypeAliasDecl - This is a declaration of a type alias, for example:
///
/// ```
///    type Foo = Int
/// ```
type TypeAliasDecl : TypeDecl {
    ///
    type: NominalType @20

    /// the 'generic-parameter' type identifier will be in here
    scope:       Scope @30
}
