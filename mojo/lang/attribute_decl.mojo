///
///
type AttributeDecl : Decl {
    ///
    name: String @10

    full_name: String @11

    ///
    attributes: [Attribute] @12

    ///
    generic_parameters: [GenericParameter] @13

    ///
    group: GroupDecl @15

    ///
    resolved_identifiers: [Identifier] @16

    ///
    unresolved_identifiers:  [Identifier] @17 //< unresolved identifiers in this file

    ///
    type: NominalType @20 | StructType @21

    ///
    initial_value: Expression @22
}