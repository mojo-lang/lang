///
///
type DeclGroup : Decl {
    /// come from `decl_group_name` attribute, otherwise will be the index from "0"
    name: String @11
    
    ///
    attributes: [Attribute] @12
}
