///
///
///
type EnumType : DataType {
    type Value {
        id: Int
        name: String
        camel_name: String

        document: String
        attributes: Attributes

        default: Any
    }

    enclosing_type : StructType @10
    values: {String, Value} @11
}