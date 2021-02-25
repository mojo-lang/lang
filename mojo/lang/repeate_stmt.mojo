/// RepeatStmt - repeat/while statement. After type-checking, the
/// condition is of type Builtin.Int1.
type RepeatStmt : LoopStmt {
    condition: Expression @20
}