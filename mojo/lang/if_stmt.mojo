/// IfStmt - if/then/else statement.  If no 'else' is specified, then the
/// ElseLoc location is not specified and the Else statement is null. After
/// type-checking, the condition is of type Builtin.Int1.
type IfStmt : BranchStmt {
    condition: Expression @20

    then_branch: BlockStmt @21
    else_branch: BlockStmt @22 | IfStmt @23
}
