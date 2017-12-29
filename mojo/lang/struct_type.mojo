///
///
///
type StructType : DataType {
    // A Field represents a Field declaration list in a struct type,
    // a method list in an interface type, or a parameter/result declaration
    // in a signature.
    //
    type Field {
        id: UInt @1
        name: String @2
        camel_name: String @3

        document: Document @4
        attributes: Attributes @5

        type: (DataType @6 | StructType @7 | EnumType @8) @link
    }

    enclosing_struct_type : StructType @10

    nested_struct_types: [StructType] @11
    nested_enum_types: [EnumType] @12

    fields: [Field] @13
}
