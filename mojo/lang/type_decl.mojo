///
///
type TypeDecl : Decl {
    ///
    name: String @10

    // mojo.core.Cached
    full_name: String @11
    
    ///
    generic_parameters: [GenericParameter] @12

    ///
    attributes: [Attribute] @13

    ///
    enclosing_type : NominalType @14

    ///
    group: GroupDecl @15

    ///
    dependencies: [Identifier] @16
}
