/// A Scope maintains the set of named language entities declared
/// in the scope and a link to the immediately surrounding (enclosing)
/// scope.
///
type Scope {
    /// `full_name` => `Identifier`
    identifiers: {String: Identifier} @1

    enclosing: Scope @10
}
