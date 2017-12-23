
type Attributes = {String, Term @shared}

type DataType {
    type Inherit {
        type: DataType @1 @link
        attributes: [Attribute] @2
    }

    name: String @1
    full_name: String @2
    document: String @3
    attributes: Attributes @4

    package: Package @5 @link

    inherits: [Inherit] @6
}
