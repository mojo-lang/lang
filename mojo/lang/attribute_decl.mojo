///
///
type AttributeDecl : Decl {
    ///
    name: String @9

    ///
    attributes: [Attribute] @11

    ///
    group: GroupDecl @13

    ///
    type: NominalType @15 | StructType @16

    ///
    default_value: Expression @17
}