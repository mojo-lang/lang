///
///
///
type Package {
    parent: Package

    version: Version
    name : String
    full_name: String

    summary: String

    repository: Uri

    license: String

    imports: [Import]
    exports: [String]

    source_directories: [String]
    source_files: [SourceFile]

    //?dependencies: {String, Range<Version>}

    // auto generate
    attributes: {String, Attribute}

    variables: {String, Variable}
    constants: {String, Constant}
    functions: {String, Function}

    value_types: {String, DataType}
    struct_types: {String, StructType}
    enum_types: {String, EnumType}
    service_types: {String, ServiceType}
}
