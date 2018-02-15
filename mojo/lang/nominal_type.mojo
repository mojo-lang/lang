///
///
type NominalType {
    ///
    name: String @1
    
    ///
    type_declaration: TypeDeclaration

    ///
    generic_arguments: [GenericArgument]

    ///
    attributes: [Attribute]
}
