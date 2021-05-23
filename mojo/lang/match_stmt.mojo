type MatchCase : BranchStmt {
    pattern: Pattern @20
    body: BlockStmt @21 | Expression @22
}

/// Match statement.
type MatchStmt : BranchStmt {
    subject: Expression @20
    cases: MatchCase @22
}