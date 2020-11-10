/// ReturnStmt - A return statement.  The result is optional; "return" without
/// an expression is semantically equivalent to "return ()".
///    return 42
type ReturnStmt : ControlTransferStmt {
    result: Expression @10 @optional
}