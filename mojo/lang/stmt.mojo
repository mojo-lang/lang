/// Stmt - Base type for all statements in mojo.
type Stmt {
    /// Kind - The subclass of Stmt that this is.
    kind : Int @1

    /// Implicit - Whether this statement is implicit.
    implicit : Bool @2
}