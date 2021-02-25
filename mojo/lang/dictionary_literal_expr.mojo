/// \brief A dictionary literal expression {a : x, b : y, c : z}.
type DictionaryLiteralExpr : LiteralExpr {
    type Entry {
        key: Expression @1
        value: Expression @2
    }

    entries: [Entry] @15
}
