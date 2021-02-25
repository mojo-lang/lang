type ImportDecl : Decl {

    /// import package name
    package:       String @10

    /// alias for the package name
    package_alias: String @11

    /// imported identifiers
    identifiers: [Identifier] @13
}