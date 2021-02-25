///
///
@discriminator('@type')
@label_format('{}')
type Declaration = PackageDecl     @1
                 | ImportDecl      @2
                 | EnumDecl        @3
                 | StructDecl      @4
                 | TypeAliasDecl   @5
                 | InterfaceDecl   @6
                 | ConstantDecl    @10
                 | VariableDecl    @11
                 | AttributeDecl   @12
                 | FunctionDecl    @13
                 | ConstructorDecl @14
                 | GenericParameter @19
                 | GroupDecl       @20
