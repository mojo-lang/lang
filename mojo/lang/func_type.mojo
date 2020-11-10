///
type FuncType {
    ///
    package: String @4 | Package @5 @reference

    ///
    parameters: [ValueDecl] @10

    ///
    return_type: NominalType @11
}