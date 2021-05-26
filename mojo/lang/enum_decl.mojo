/// \brief This is the declaration of an enum.
///
/// For example:
///
/// \code
/// \endcode
///
/// The type of the decl itself is a MetatypeType; use getDeclaredType()
/// to get the declared type ("Bool" or "Optional" in the above example).
///
/// \brief This represents a single case of an 'enum' declaration.
///
/// For example, the X, Y, and Z in this enum:
///
/// \code
///   enum V : Int {
///     x, y, z
///   }
/// \endcode
///
/// The type of an EnumElementDecl is always the EnumType for the containing
/// enum. EnumElementDecls are represented in the AST as members of their
/// parent EnumDecl, although syntactically they are subordinate to the
/// EnumCaseDecl.
//class EnumElementDecl : public ValueDecl {
//}

///
type EnumDecl : TypeDecl @disable_fields("generic_parameters") {
    ///
    type: EnumType @20

    scope:   Scope @30
}
