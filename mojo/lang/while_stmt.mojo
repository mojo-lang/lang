/// WhileStmt - while statement. After type-checking, the condition is of
/// type Builtin.Int1.
type WhileStmt : LoopStmt {
    condition: Expression @20
}