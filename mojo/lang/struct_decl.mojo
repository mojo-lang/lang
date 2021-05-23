/// StructDecl - This is the declaration of a struct, for example:
///
/// ``` mojo
/// struct Complex { r : Double, i : Double }
/// ```
///
/// The type of the decl itself is a MetaType; use `Complex.type`
/// to get the `Complex`'s declared type.
///
type StructDecl : TypeDecl {
    ///
    type: StructType @20

    ///
    type_alias_decls: [TypeAliasDecl] @21

    ///
    enum_decls: [EnumDecl] @22

    ///
    struct_decls: [StructDecl] @23

    scope:       Scope @30
}