/// Semantic Versioning 2.0.0
///
/// Given a version number MAJOR.MINOR.PATCH, increment the:
///
/// MAJOR version when you make incompatible API changes,
/// MINOR version when you add functionality in a backwards compatible manner, and
/// PATCH version when you make backwards compatible bug fixes.
/// Additional labels for pre-release and build metadata are available as extensions to the MAJOR.MINOR.PATCH format.
///
type Version {
    /// major version
    major: String @1

    /// minor version
    minor: String @2

    /// patch version
    patch: String @3

    /// pre-release version
    pre_release: [String] @4

    /// build metadata
    build: [String] @5
}