/// A declaration of a generic type parameter.
///
/// A generic type parameter introduces a new, named type parameter along
/// with some set of requirements on any type argument used to realize this
/// type parameter. The requirements involve conformances to specific
/// protocols or inheritance from a specific class type.
///
/// In the following example, 'T' is a generic type parameter with the
/// requirement that the type argument conform to the 'Comparable' protocol.
///
/// \code
/// func min<T : Comparable>(x : T, y : T) -> T { ... }
/// \endcode

/// A type declaration that can have generic parameters attached to it.  Because
/// it has these generic parameters, it is always a DeclContext.
type GenericParameter : TypeDecl @skip_fields() {
    ///
    name: String @10

    ///
    constraint: NominalType @11
}