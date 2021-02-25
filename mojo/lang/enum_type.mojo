///
///
///
type EnumType {
    ///
    package: String @4 | Package @5 @Reference

    //
    enumerators: [ValueDecl] @10

    //
    underlying_type: NominalType @11
}
