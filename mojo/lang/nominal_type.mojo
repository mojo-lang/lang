///
///
type NominalType {
    /// position of first character belonging to the Expr
    start_position: Position @1

    /// position of first character immediately after the Expr
    end_position: Position @2

    ///
    document: Document @4

    /// user writing package name
    /// var foo : package.name.Bar
    /// ==> mojo.package.name.Bar
    package: String @5

    ///
    name: String @10

    ///
    type_declaration: TypeDeclaration @11 @reference

    ///
    generic_arguments: [GenericArgument] @12

    ///
    attributes: [Attribute] @13

    ///
    enclosing_type: NominalType @14
}

//func enclosing_types(NominalType type) -> [NominalType] {
//    var enclosing_types
//    while enclosing = type.enclosing_type {
//        enclosing_types.append(enclosing)
//        type = enclosing
//    }
//    return enclosing_types.reverse
//}