/// `InterfaceDecl` - A declaration of an interface, for example:
///
/// ```mojo
///   interface Drawable {
///      draw()
///   }
/// ```
///
type InterfaceDecl : TypeDecl {
    ///
    type: InterfaceType @13

    ///
    type_alias_decls: [TypeAliasDecl] @15
}
