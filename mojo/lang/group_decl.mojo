///
///
type GroupDecl : Decl {
    /// come from `decl_group_name` attribute, otherwise will be the index from "0"
    name: String @11

    /// may be the `struct`, `enum`, `constant`, `variable`, `attribute` `function` `value`
    type: String @12
    
    ///
    attributes: [Attribute] @13

    ///
    declarations: [Declaration @reference]
}
