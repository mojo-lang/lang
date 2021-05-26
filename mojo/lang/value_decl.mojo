/// VarDecl - 'var' and 'const' declarations.
type ValueDecl : Decl {
    ///
    name: String @10

    ///
    attributes: [Attribute] @12

    ///
    group: GroupDecl @15

    ///
    type: NominalType @20

    ///
    initial_value: Expression @22
}
