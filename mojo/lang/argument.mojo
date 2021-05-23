/// argument like in function call like expressions
type Argument {
    /// position of first character belonging to the Expr
    start_position: Position @1

    /// position of first character immediately after the Expr
    end_position: Position @2

    /// 
    label: String @10
    
    value: Expression @11
}
