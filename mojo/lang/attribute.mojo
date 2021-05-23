
///
type Attribute {
    /// position of first character belonging to the Expr
    start_position: Position @1

    /// position of first character immediately after the Expr
    end_position: Position @2

    /// the package which this attribute belong to
    package: String @5

    name: String @10
    
    declaration: AttributeDecl @11 @reference

    ///
    generic_arguments: [GenericArgument] @12

    arguments: [Argument] @13

    value: Expression @20
}