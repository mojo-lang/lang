
///
type Attribute {
    /// the package which this attribute belong to
    package: Package @1 | String @2

    name: String @3
    
    arguments: [Argument] @5

    value: Expression @10
}