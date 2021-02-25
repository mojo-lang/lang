/// Stmt - Base type for all statements in mojo.
type Stmt {
    // statement start position
    start_position: Position @1 

    // statement end position
    end_position: Position @2

    /// Kind - The subclass of Stmt that this is.
    kind : Int @3

    /// Implicit - Whether this statement is implicit.
    implicit : Bool @4
}