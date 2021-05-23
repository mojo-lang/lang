///
///
type GroupDecl : Decl {
    /// come from `decl_group_name` attribute, otherwise will be the index from "0"
    name: String @10

    ///
    attributes: [Attribute] @12

    /// may be the `struct`, `enum`, `constant`, `variable`, `attribute` `function`
    type: Identifier.Kind @20

    ///
    declarations: [Declaration @reference] @21
}
