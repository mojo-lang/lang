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
    type: StructType @13

    ///
    enclosing_type : NominalType @14

    ///
    type_alias_decls: [TypeAliasDecl] @15

    ///
    struct_decls: [StructDecl] @17

    ///
    enum_decls: [EnumDecl] @18
}