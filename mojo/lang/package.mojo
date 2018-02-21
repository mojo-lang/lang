/// A Package represents a set of source files
/// collectively building a Mojo package.
type Package {
    /// package name
    name : String @1

    ///
    full_name: String @2

    ///
    version: Version @3

    ///
    parent: Package @4

    ///
    children: [Package] @5

    ///
    summary: String @10

    ///
    repository: Uri @11

    ///
    license: String @12

    ///
    imports: [Import]

    ///
    exports: [String]

    ///
    //dependencies: {String: Range<Version>}

    ///
    source_directories: [String]

    ///
    source_files: [SourceFile]

    ///
    scope: Scope
}

func to<T:String>(package: Package) -> T {
    package.full_name
}
