/// VarDecl - 'var' and 'const' declarations.
type ValueDecl : Decl {
    ///
    name: String @10

    ///
    type: NominalType @12

    ///
    default_value: Expression @13

    ///
    group: DeclGroup @14
}