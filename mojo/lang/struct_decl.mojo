/// StructDecl - This is the declaration of a struct, for example:
///
///    struct Complex { R : Double, I : Double }
///
/// The type of the decl itself is a MetatypeType; use getDeclaredType()
/// to get the declared type ("Complex" in the above example).
type StructDecl : TypeDecl {
    ///
    type: StructType @13

    ///
    ///
    enclosing_type : NominalType @15

    ///
    type_alias_decls: [TypeAliasDecl] @16
    
    ///
    struct_decls: [StructDecl] @17
    
    ///
    enum_decls: [EnumDecl] @18
}