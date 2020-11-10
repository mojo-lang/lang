///
///
type NominalType {
    ///
    name: String @1

    ///
    package: String @4 | Package @5 @reference

    ///
    type_declaration: TypeDeclaration @10 @reference

    ///
    generic_arguments: [GenericArgument] @11

    ///
    attributes: [Attribute] @12

    ///
    enclosing_type : NominalType @14
}

func enclosing_types(NominalType type) -> [NominalType] {
    var enclosing_types
    while enclosing = type.enclosing_type {
        enclosing_types.append(enclosing)
        type = enclosing
    }
    return enclosing_types.reverse
}