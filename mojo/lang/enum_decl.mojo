
/// \brief This represents a single case of an 'enum' declaration.
///
/// For example, the X, Y, and Z in this enum:
///
/// \code
///   enum V {
///     case X(Int), Y(Int)
///     case Z
///   }
/// \endcode
///
/// The type of an EnumElementDecl is always the EnumType for the containing
/// enum. EnumElementDecls are represented in the AST as members of their
/// parent EnumDecl, although syntactically they are subordinate to the
/// EnumCaseDecl.
class EnumElementDecl : public ValueDecl {
}

///
type EnumDecl : TypeDecl @disable_fields(generic_params) {
    ///
    type: EnumType @13
}



/// \brief This is the declaration of an enum.
///
/// For example:
///
/// \code
/// \endcode
///
/// The type of the decl itself is a MetatypeType; use getDeclaredType()
/// to get the declared type ("Bool" or "Optional" in the above example).
class EnumDecl final : public NominalTypeDecl {
}