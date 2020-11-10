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
    type: InterfaceType @15

    ///
    type_alias_decls: [TypeAliasDecl] @16

    ///
    nominal_decls: [NominalDecl] @17
}
