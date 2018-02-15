///
///
///
type InterfaceType {
    ///
    name: String @1

    ///
    package: Package @5 @link

    ///
    methods: [FuncDecl] @10

    ///
    inherits: [NominalType] @11
}