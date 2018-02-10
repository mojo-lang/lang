/// ProtocolDecl - A declaration of a protocol, for example:
///
///   interface Drawable {
///      draw()
///   }
type InterfaceDecl : TypeDecl @disable_fields(generic_params) {
    ///
    type: InterfaceType @13
    
    ///
    attributes: [Attribute] @13
}