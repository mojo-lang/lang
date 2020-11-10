///
type Attribute {
    /// the package which this attribute belong to
    package: Package @1 | String @2

    name: String @3

    generic_arguments: [GenericArgument] @4
    
    expressions: [Expression] @5
}