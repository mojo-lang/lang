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
/// type X<T> {
///	    x : T
/// }
///
/// X<Int @argument>(i : Int)
///
/// X<T>(i: T) {
/// }
type ConstructorDecl : Decl {
    ///
    name: String @10

    ///
    full_name: String @11

    ///
    attributes: [Attribute] @12

    ///
    generic_parameters: [GenericParameter] @13

    ///
    signature: FunctionDecl.Signature @20

    ///
    body: BlockStmt @21

    ///
    scope: Scope @30
}
