///
///
type Statement = ReturnStmt |
                 BreakStmt |
                 ContinueStmt |
                 MatchStmt |
                 IfStmt |
                 ForEachStmt |
                 WhileStmt |
                 RepeatWhileStmt |
                 BlockStmt |
                 Declaration |
                 Expression

/// Stmt - Base type for all statements in mojo.
type Stmt {
    /// Kind - The subclass of Stmt that this is.
    kind : Int

    /// Implicit - Whether this statement is implicit.
    implicit : Bool
}

type BadStmt : Stmt {
}

type EmptyStmt : Stmt {}

/// BlockStmt - A brace enclosed sequence of expressions, stmts, or decls, like
/// { var x = 10; print(x) }.
type BlockStmt : Stmt {
}

type ControlTransferStmt : Stmt {}

/// ReturnStmt - A return statement.  The result is optional; "return" without
/// an expression is semantically equivalent to "return ()".
///    return 42
type ReturnStmt : ControlTransferStmt {
    result: Expression @10 @optional
}

/// BreakStmt - The "break" and "break label" statement.
type BreakStmt : ControlTransferStmt

/// ContinueStmt - The "continue" and "continue label" statement.
type ContinueStmt : ControlTransferStmt

///
type BranchStmt : Stmt {
}

/// IfStmt - if/then/else statement.  If no 'else' is specified, then the
/// ElseLoc location is not specified and the Else statement is null. After
/// type-checking, the condition is of type Builtin.Int1.
type IfStmt : BranchStmt {
    conditions: [Expression]

    then: Statement
    else: Statement
}

type MatchCase : Stmt {
    pattern: Pattern
    body: BlockStmt | Expression
}

/// Match statement.
type MatchStmt : BranchStmt {
    subject: Expression
    cases: MatchCase
}

///
type LoopStmt : Stmt {
    body: BlockStmt
}

/// ForEachStmt - foreach statement that iterates over the elements in a
/// container.
///
/// Example:
/// ```
/// for i in 0..10 {
///   print(i.to<String>)
/// }
/// ```
type ForEachStmt : LoopStmt {
    sequence: Expression
}

/// WhileStmt - while statement. After type-checking, the condition is of
/// type Builtin.Int1.
type WhileStmt : LoopStmt {
    conditions: [Expression]
}

/// RepeatWhileStmt - repeat/while statement. After type-checking, the
/// condition is of type Builtin.Int1.
type RepeatWhileStmt : LoopStmt {
    condition: Expression
}
