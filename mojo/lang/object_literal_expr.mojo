// ObjectLiteralExpr - An expression of the form
// '{red: 1, blue: 0, green: 0, alpha: 1}' with a name and a list
// argument. The components of the list argument are meant to be themselves
// constant.
type ObjectLiteralExpr : LiteralExpr {
    type Field {
        name: Expression @1
        value: Expression @2
    }

    fields: [Field] @15
}
