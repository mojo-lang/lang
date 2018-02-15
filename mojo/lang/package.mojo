///
///
///
type Package {
    ///
    name : String

    ///
    full_name: String

    ///
    version: Version

    ///
    parent: Package

    ///
    children: [Package]

    summary: String

    repository: Uri

    license: String

    imports: [Import]

    exports: [String]

    source_directories: [String]

    source_files: [SourceFile]

    //?dependencies: {String, Range<Version>}

    // auto generated
    attribute_decls: {String, AttributeDecl} @20 @linked
    const_decls:     {String, ConstDecl} @21 @linked

    data_decls:      {String, DataDecl} @22 @linked
    struct_decls:    {String, StructDecl} @23 @linked
    enum_decls:      {String, EnumDecl} @24 @linked
    interface_decls: {String, InterfaceDecl} @25 @linked
    type_alias_decls:{String, TypeAliasDecl} @26 @linked

    func_decls:      {String, FunctionDecl} @27 @linked
}

func to<T:String>(package: Package) -> T {
    package.full_name
}
