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
    type: StructType @15

    ///
    type_alias_decls: [TypeAliasDecl] @16

    ///
    nominal_decls: [NominalDecl] @17
    
    ///
    struct_decls: [StructDecl] @18

    ///
    enum_decls: [EnumDecl] @19
}