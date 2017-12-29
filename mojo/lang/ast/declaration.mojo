type Declaration =


type Decl;

/// ValueDecl - All named decls that are values in the language.  These can
/// have a type, etc.
class ValueDecl : public Decl

/// This is a common base class for declarations which declare a type.
class TypeDecl : public ValueDecl

/// A type declaration that can have generic parameters attached to it.  Because
/// it has these generic parameters, it is always a DeclContext.
class GenericTypeDecl : TypeDecl, GenericContext {
}

/// TypeAliasDecl - This is a declaration of a typealias, for example:
///
///    typealias Foo = Int
///
/// TypeAliasDecl's always have 'MetatypeType' type.
///
class TypeAliasDecl : public GenericTypeDecl {

/// Abstract class describing generic type parameters and associated types,
/// whose common purpose is to anchor the abstract type parameter and specify
/// requirements for any corresponding type argument.
class AbstractTypeParamDecl : public TypeDecl {
}

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
class GenericTypeParamDecl : public AbstractTypeParamDecl {

}

/// NominalTypeDecl - a declaration of a nominal type, like a struct.
class NominalTypeDecl : public GenericTypeDecl, public IterableDeclContext

/// \brief This is the declaration of an enum.
///
/// For example:
///
/// \code
///    enum Bool {
///      case false
///      case true
///    }
///
///    enum Optional<T> {
///      case none
///      case some(T)
///    }
/// \endcode
///
/// The type of the decl itself is a MetatypeType; use getDeclaredType()
/// to get the declared type ("Bool" or "Optional" in the above example).
class EnumDecl final : public NominalTypeDecl {
}

/// StructDecl - This is the declaration of a struct, for example:
///
///    struct Complex { var R : Double, I : Double }
///
/// The type of the decl itself is a MetatypeType; use getDeclaredType()
/// to get the declared type ("Complex" in the above example).
class StructDecl final : public NominalTypeDecl


/// ProtocolDecl - A declaration of a protocol, for example:
///
///   service Drawable {
///     func draw()
///   }
type ServiceDecl : NominalTypeDecl

/// VarDecl - 'var' and 'let' declarations.
class VarDecl : public AbstractStorageDecl {

}

/// A function parameter declaration.
class ParamDecl : public VarDecl {
}

/// \brief Base class for function-like declarations.
class AbstractFunctionDecl : public ValueDecl, public GenericContext

/// FuncDecl - 'func' declaration.
class FuncDecl final : public AbstractFunctionDecl,

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

/// ConstructorDecl - Declares a constructor for a type.  For example:
///
/// \code
/// struct X {
///   var x : Int
///   init(i : Int) {
///      x = i
///   }
/// }
/// \endcode
class ConstructorDecl : public AbstractFunctionDecl {
}

/// Abstract base class of operator declarations.
class OperatorDecl : public Decl {

}

/// Declares the behavior of an infix operator. For example:
///
/// \code
/// infix operator /+/ : AdditivePrecedence
/// \endcode
class InfixOperatorDecl : public OperatorDecl {

}

/// Declares the behavior of a prefix operator. For example:
///
/// \code
/// prefix operator /+/ {}
/// \endcode
class PrefixOperatorDecl : public OperatorDecl {

}

/// Declares the behavior of a postfix operator. For example:
///
/// \code
/// postfix operator /+/ {}
/// \endcode
class PostfixOperatorDecl : public OperatorDecl {
}
