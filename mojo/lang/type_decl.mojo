///
///
type TypeDecl : Decl {
    ///
    name: String @10
    
    ///
    generic_parameters: [GenericParameter] @11

    ///
    attributes: [Attribute] @12

    ///
    enclosing_type : NominalType @13

    ///
    group: DeclGroup @14
}
