/// IfStmt - if/then/else statement.  If no 'else' is specified, then the
/// ElseLoc location is not specified and the Else statement is null. After
/// type-checking, the condition is of type Builtin.Int1.
type IfStmt : BranchStmt {
    condition: Expression

    then_branch: BlockStmt
    else_branch: BlockStmt | IfStmt
}
