///
///
type Term {
    type:     String @1
    value:    String @2

    location: PositionSpan @10

    terms:    [Term] @15
}
