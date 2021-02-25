/// A Package represents a set of source files
/// collectively building a Mojo package.
type Package {
    ///
    type RequiredVersion = Cached<VersionRange>

    ///
    ///
    type Requirement {
        /// ^1.2.3
        /// ~1.2.3
        /// >= 1.2, < 1.5
        version: RequiredVersion @1

        ///
        registry: String @3
        
        ///
        path: String @4
        
        ///
        git: Url @5
        
        ///
        branch: String @6
    }

    /// package name
    name : String @1

    ///
    full_name: String @2

    ///
    version: Version @3

    ///
    parent: Package @4 @reference('full_name')

    ///
    children: [Package] @5

    //
    authors: [String] @6

    /// A description of the package.
    description: String @10

    /// 
    document: Cached<Document> @11

    /// URL of the package source repository.
    repository: Url @12

    /// The package license.
    license: String @13

    ///
    imports: [Import] @14

    ///
    exports: [string] @15

    ///
    source_files: [SourceFile] @16

    /// package scope across all files
    scope: Scope @20

    ///
    dependencies: {String: Requirement} @30

    /// 
    resolved_dependencies: {String, Package} @31
}

func to<T:String>(package: Package) -> T {
    return package.full_name
}

