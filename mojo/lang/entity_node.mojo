
type EntityNode {
    /// full name of the type
    name: String @1 @key

    /// the type declaration
    type_declaration: TypeDeclaration @2 @reference


    key_field: ValueDecl @5
}
