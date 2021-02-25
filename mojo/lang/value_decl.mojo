/// VarDecl - 'var' and 'const' declarations.
type ValueDecl : Decl {
    ///
    name: String @10

    ///
    type: NominalType @12

    ///
    initial_value: Expression @13

    ///
    group: GroupDecl @14
}
