///
///
type ValueGroupDecl : Decl {
    /// come from `group_decl_name` attribute, otherwise will be the index from "0"
    name: String @11
    
    /// `inline` means 
    attributes: [Attribute] @13

    ///
    declarations: [ValueDecl @reference('name')]
}
