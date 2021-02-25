/// `InterfaceDecl` - A declaration of an interface, for example:
///
/// ```mojo
///   interface Drawable {
///      draw()
///   }
/// ```
///
type InterfaceDecl : TypeDecl @skip_fields(enclosing_type) {
    ///
    type: InterfaceType @20

    ///
    type_alias_decls: [TypeAliasDecl] @21

    scope:       Scope @30
}
