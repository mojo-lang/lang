/// ConstructorDecl - Declares a constructor for a type.  For example:
///
/// ```
/// type X {
///     x: Int
/// }
///
/// X(i : Int) {
///    x = i
/// }
/// ```
type ConstructorDecl : Decl {

    ///
    type_declaration: TypeDeclaration

    ///
    generic_arguments: [GenericArgument]

    ///
    attributes: [Attribute]

    ///
    body: BlockStmt @14
}
