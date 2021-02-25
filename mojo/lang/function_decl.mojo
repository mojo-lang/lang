///
///
type FunctionDecl : Decl {
    ///
    type Signature {
        ///
        parameters: [ValueDecl] @1

        ///
        result: NominalType @2
    }

    ///
    name: String @10
    
    ///
    generic_parameters: [GenericParameter] @11
    
    ///
    signature: Signature @12

    ///
    attributes: [Attribute] @13

    ///
    body: BlockStmt @14
}
