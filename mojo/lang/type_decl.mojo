///
///
type TypeDecl : Decl {
    ///
    name: String @10
    
    ///
    generic_parameters: [GenericParameter] @11

    ///
    attributes: [Attribute] @12
}
