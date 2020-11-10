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
    repository: Url @11

    ///
    license: String @12

    ///
    imports: [Import] @13

    ///
    exports: [String] @14

    ///
    //dependencies: {String: Range<Version>}

    ///
    source_directories: [String] @16

    ///
    source_files: [SourceFile] @17

    ///
    scope: Scope @20
}

func to<T:String>(package: Package) -> T {
    package.full_name
}
