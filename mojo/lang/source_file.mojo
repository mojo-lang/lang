///
///
///
type SourceFile {
    name:      String @1
    full_name: String @2
    package: Package  @3

    // statements
    statements: [Statement] @5

    // decls' index
    import_decls:    [ImportDecl] @10 @linked

    const_decls:     [ConstDecl] @11 @linked
    attribute_decls: [AttributeDecl] @12 @linked

    data_decls:      [DataDecl] @13 @linked
    struct_decls:    [StructDecl] @14 @linked
    enum_decls:      [EnumDecl] @15 @linked
    interface_decls: [InterfaceDecl] @16 @linked

    func_decls:      [FuncDecl] @17 @linked
}