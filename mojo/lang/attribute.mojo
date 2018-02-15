///
type Attribute {
    package: Package @1 | String @2

    name: String @3
    expression: Expression @4
}