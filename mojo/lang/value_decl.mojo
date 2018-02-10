/// VarDecl - 'var' and 'const' declarations.
type ValueDecl : Decl {
    ///
    name: Identifier @10

    ///
    type: NominalType @12 @link

    ///
    attributes: [Attribute] @13

    ///
    value: Expression @14
}
