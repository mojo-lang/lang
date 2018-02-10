type Expression = NoneLiteralExpr |
                  IntegerLiteralExpr |
                  FloatLiteralExpr |
                  BooleanLiteralExpr |
                  StringLiteralExpr |
                  ObjectLiteralExpr |
                  ArrayExpr |
                  DictionaryExpr

/// DiscardAssignmentExpr - A '_' in the left-hand side of an assignment, which
/// discards the corresponding tuple element on the right-hand side.
type DiscardAssignmentExpr : Expr

/// MemberRefExpr - This represents 'a.b' where we are referring to a member
/// of a type, such as a property or variable.
///
/// Note that methods found via 'dot' syntax are expressed as DotSyntaxCallExpr
/// nodes, because 'a.f' is actually an application of 'a' (the implicit object
/// argument) to the function 'f'.
type MemberExpr

/// Common base for expressions that involve dynamic lookup, which
/// determines at runtime whether a particular method, property, or
/// subscript is available.
type DynamicLookupExpr : Expr

/// UnresolvedMemberExpr - This represents '.foo', an unresolved reference to a
/// member, which is to be resolved with context sensitive type information into
/// bar.foo.  These always have unresolved type.
class UnresolvedMemberExpr

/// An expression node that does not affect the evaluation of its subexpression.
class IdentityExpr : public Expr

/// Subscripting expression that applies a keypath to a base.
class KeyPathApplicationExpr : public Expr

/// TupleElementExpr - Refer to an element of a tuple,
/// e.g. "(1,field:2).field".
class TupleElementExpr : public Expr 

/// ImplicitConversionExpr - An abstract class for expressions which
/// implicitly convert the value of an expression in some way.
class ImplicitConversionExpr : public Expr

type KeyPathExpr

/// CaptureListExpr - This expression represents the capture list on an explicit
/// closure.  Because the capture list is evaluated outside of the closure, this
/// CaptureList wraps the ClosureExpr.  The dynamic semantics are that evaluates
/// the variable bindings from the capture list, then evaluates the
/// subexpression (the closure itself) and returns the result.
class CaptureListExpr : public Expr

/// DynamicTypeExpr - "type(of: base)" - Produces a metatype value.
///
/// The metatype value comes from evaluating an expression then retrieving the
/// metatype of the result.
class DynamicTypeExpr : public Expr 

/// An expression referring to an opaque object of a fixed type.
///
/// Opaque value expressions occur when a particular value within the AST
/// needs to be re-used without being re-evaluated or for a value that is
/// a placeholder. OpaqueValueExpr nodes are introduced by some other AST
/// node (say, a \c DynamicMemberRefExpr) and can only be used within the
/// subexpressions of that AST node.
class OpaqueValueExpr

/// CallExpr - Application of an argument to a function, which occurs
/// syntactically through juxtaposition with a TupleExpr whose
/// leading '(' is un-spaced.
class CallExpr : ApplyExpr

/// PrefixUnaryExpr - Prefix unary expressions like '!y'.
class PrefixUnaryExpr : public ApplyExpr

/// PostfixUnaryExpr - Prefix unary expressions like '!y'.
class PostfixUnaryExpr : public ApplyExpr

/// BinaryExpr - Infix binary expressions like 'x+y'.  The argument is always
/// an implicit tuple expression of the type expected by the function.
class BinaryExpr : public ApplyExpr {
}

/// DotSyntaxCallExpr - Refer to a method of a type, e.g. P.x.  'x'
/// is modeled as a DeclRefExpr or OverloadSetRefExpr on the method.
class DotSyntaxCallExpr : public SelfApplyExpr

/// ConstructorRefCallExpr - Refer to a constructor for a type P.  The
/// actual reference to function which returns the constructor is modeled
/// as a DeclRefExpr.
class ConstructorRefCallExpr : public SelfApplyExpr

/// DotSyntaxBaseIgnoredExpr - When a.b resolves to something that does not need
/// the actual value of the base (e.g. when applied to a metatype, module, or
/// the base of a 'static' function) this expression node is created.  The
/// semantics are that its base is evaluated and discarded, then 'b' is
/// evaluated and returned as the result of the expression.
class DotSyntaxBaseIgnoredExpr : public Expr

/// \brief A pattern production that has been parsed but hasn't been resolved
/// into a complete pattern. Name binding converts these into standalone pattern
/// nodes or raises an error if a pattern production appears in an invalid
/// position.
class UnresolvedPatternExpr : public Expr


/// Represents the unusual behavior of a . in a \ keypath expression, such as
/// \.[0] and \Foo.?.
class KeyPathDotExpr : public Expr

/// ErrorExpr - Represents a semantically erroneous subexpression in the AST,
/// typically this will have an ErrorType.
type ErrorExpr : Expr 
