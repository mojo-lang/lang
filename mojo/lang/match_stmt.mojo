type MatchCase : BranchStmt {
    pattern: Pattern
    body: BlockStmt | Expression
}

/// Match statement.
type MatchStmt : BranchStmt {
    subject: Expression
    cases: MatchCase
}