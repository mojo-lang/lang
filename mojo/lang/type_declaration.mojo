///
///
@discriminator('@type')
@label_format('{}')
type TypeDeclaration = EnumDecl        @3
                     | StructDecl      @4
                     | TypeAliasDecl   @5
                     | InterfaceDecl   @6
                     | GenericParameter @19