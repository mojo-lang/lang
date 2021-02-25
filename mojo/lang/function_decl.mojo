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
    full_name: String @11

    ///
    attributes: [Attribute] @12
    
    ///
    generic_parameters: [GenericParameter] @13
    
    ///
    signature: Signature @20

    ///
    body: BlockStmt @21

    ///
    scope: Scope @30
}
