type Term {
    type:  String @1
    value: String @2

    location: PosSpan @10

    terms: [Term] @15
}
