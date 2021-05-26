///
///
type TypeDecl : Decl {
    ///
    name: String @10

    // mojo.core.Cached
    //full_name: String @11
    
    ///
    attributes: [Attribute] @12

    ///
    generic_parameters: [GenericParameter] @13

    ///
    enclosing_type : NominalType @14

    ///
    group: GroupDecl @15

    ///
    resolved_identifiers: [Identifier] @16

    ///
    unresolved_identifiers:  [Identifier] @17 //< unresolved identifiers in this file
}
