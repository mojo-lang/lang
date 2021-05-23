///
///
type ValueGroupDecl : Decl {
    /// come from `group_decl_name` attribute, otherwise will be the index from "0"
    name: String @10
    
    /// `inline` means 
    attributes: [Attribute] @12

    ///
    declarations: [ValueDecl @reference('name')] @15
}
