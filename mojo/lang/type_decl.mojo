///
///
type TypeDecl : Decl {
    ///
    name: String @10
    
    ///
    generic_params: [GenericParam] @11

    ///
    attributes: [Attribute] @12
}
